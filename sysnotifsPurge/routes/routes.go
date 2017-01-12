package routes

import (
	"sysnotifsPurge/handlers"
	"sysnotifsPurge/structs"
)

//ListRoutes ...
var ListRoutes = structs.Routes{
	// GET
	// get user information
	structs.Route{
		Name:        "CheckRatio",
		Method:      "GET",
		Pattern:     "/CheckRatio",
		HandlerFunc: handlers.CheckRatio,
	},
}
