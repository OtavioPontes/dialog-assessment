package routes

import (
	"net/http"

	"github.com/otaviopontes/api-go/src/controllers"
)

var userRoutes = []Route{
	{
		Uri:                   "/api/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},

	{
		Uri:                   "/api/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/users/{id}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
