package domain

import (
	"github.com/labstack/echo"
)

// Route type
// Note that DefaultVersion must exists in RouteHandlers map
// See routes.go for examples
type Route struct {
	Name         string
	Method       string
	Pattern      string
	RouteHandler echo.HandlerFunc
}

// Routes type
type Routes []Route

// Append Returns a new slice of Routes
func (r *Routes) Append(routes ...*Routes) Routes {
	res := Routes{}
	// copy current route
	for _, route := range *r {
		res = append(res, route)
	}
	for _, rts := range routes {
		for _, route := range *rts {
			res = append(res, route)
		}
	}
	return res
}
