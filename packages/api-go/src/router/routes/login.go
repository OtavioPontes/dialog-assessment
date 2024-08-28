package routes

import (
	"net/http"

	"github.com/otaviopontes/api-go/src/controllers"
)

var loginRoute = Route{
	Uri:                   "/api/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequireAuthentication: false,
}
