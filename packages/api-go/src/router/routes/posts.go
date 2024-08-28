package routes

import (
	"net/http"

	"github.com/otaviopontes/api-go/src/controllers"
)

var routesPosts = []Route{
	{
		Uri:                   "/api/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/posts",
		Method:                http.MethodGet,
		Function:              controllers.GetPosts,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/posts/{id}",
		Method:                http.MethodGet,
		Function:              controllers.GetPost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/posts/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/posts/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequireAuthentication: true,
	},

	{
		Uri:                   "/api/posts/{id}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/api/posts/{id}/dislike",
		Method:                http.MethodPost,
		Function:              controllers.DislikePost,
		RequireAuthentication: true,
	},
}
