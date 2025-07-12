package response

type SimpleErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	SimpleError400 = SimpleErrorResponse{
		Status:  400,
		Message: "Bad Request",
	}
	SimpleError401 = SimpleErrorResponse{
		Status:  401,
		Message: "Unauthorized",
	}
	SimpleError500 = SimpleErrorResponse{
		Status:  500,
		Message: "Internal Server Error",
	}
)
