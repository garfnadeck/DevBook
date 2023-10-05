package route

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []RouteStructure{
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Func:         controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Func:         controllers.GetAllUsers,
		RequiredAuth: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodGet,
		Func:         controllers.GetSingleUser,
		RequiredAuth: false,
	},
	{
		URI:          "/user/{userID}",
		Method:       http.MethodPut,
		Func:         controllers.UpdateUser,
		RequiredAuth: false,
	}, {
		URI:          "/user/{userID}",
		Method:       http.MethodDelete,
		Func:         controllers.DeleteUser,
		RequiredAuth: false,
	},
}
