package plugin

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/ymatsukawa/backbase-go/internal/entity"
	"github.com/ymatsukawa/backbase-go/internal/handler"
	"github.com/ymatsukawa/backbase-go/internal/repository"
	"github.com/ymatsukawa/backbase-go/internal/service"
)

type Handlers struct {
	UserHandler handler.UserHandler
}

type Layers interface {
	Build()
	GetHandlers() *Handlers
}

type layers struct {
	DBClient *pgx.Conn
	Logger   *slog.Logger
	Handlers *Handlers
}

func NewLayers(dbClient *pgx.Conn, logger *slog.Logger) Layers {
	return &layers{
		DBClient: dbClient,
		Logger:   logger,
		Handlers: &Handlers{},
	}
}

func (l *layers) Build() {
	queries := entity.New(l.DBClient)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)

	l.Handlers.UserHandler = handler.NewUserHandler(userService, l.Logger)
}

func (l *layers) GetHandlers() *Handlers {
	return l.Handlers
}
