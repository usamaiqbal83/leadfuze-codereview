package users

import (
	"strings"

	"github.com/usamaiqbal83/leadfuzz-test/domain"
)

const (
	FetchUser = "Fetch User"
)

const defaultBasePath = "/api/user"

func (resource *Resource) generateRoutes(basePath string) *domain.Routes {

	if basePath == "" {
		basePath = defaultBasePath
	}

	var baseRoutes = domain.Routes{
		domain.Route{
			Name:         FetchUser,
			Method:       "POST",
			Pattern:      defaultBasePath,
			RouteHandler: resource.HandleFetchUser,
		},
	}

	routes := domain.Routes{}

	for _, route := range baseRoutes {
		r := domain.Route{
			Name:         route.Name,
			Method:       route.Method,
			Pattern:      strings.Replace(route.Pattern, defaultBasePath, basePath, -1),
			RouteHandler: route.RouteHandler,
		}

		routes = routes.Append(&domain.Routes{r})
	}

	resource.routes = &routes
	return resource.routes
}
