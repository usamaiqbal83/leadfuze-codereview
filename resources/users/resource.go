package users

import (
	"github.com/usamaiqbal83/leadfuzz-test/domain"
)

type Options struct {
	EmailService domain.IEmail
}

func NewResource(options *Options) *Resource {

	if options.EmailService == nil {
		panic("users.options.service is not present")
	}

	// create resource
	resource := &Resource{EmailService: options.EmailService}
	// generate resource routes
	resource.generateRoutes("")

	return resource
}

// UsersResource implements IResource
type Resource struct {
	EmailService domain.IEmail
	routes *domain.Routes
}

func (resource *Resource) Routes() *domain.Routes {
	return resource.routes
}

