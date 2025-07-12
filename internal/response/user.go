package response

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ymatsukawa/backbase-go/internal/entity"
)

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

func ToUsersResponse(users []entity.User) UsersResponse {
	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = toUserResponse(user)
	}
	return UsersResponse{
		Users: userResponses,
	}
}

func toUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: pgTypeToTime(user.CreatedAt),
		UpdatedAt: pgTypeToTime(user.UpdatedAt),
	}
}

func pgTypeToTime(pgTime pgtype.Timestamp) time.Time {
	if pgTime.Valid {
		return pgTime.Time
	}
	return time.Time{}
}
