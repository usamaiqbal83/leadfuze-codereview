package main

import (
	"math/rand"
	"time"

	"github.com/usamaiqbal83/leadfuzz-test/resources/users"
	"github.com/usamaiqbal83/leadfuzz-test/server"
	"github.com/usamaiqbal83/leadfuzz-test/services"
)

func main(){
	rand.Seed(time.Now().UnixNano())

	s := server.NewServer(&server.Config{
		// some configurations will be added in future
	})

	// email verification service
	service := services.NewService("tmpkey321")

	// create user resource
	userResource := users.NewResource(&users.Options{EmailService: service})

	// configure CORS
	s.ConfigureCORS(&server.CORSOptions{
		AllowOrigin:  []string{"*"},
		AllowMethods: []string{"POST"},
	})

	// add resources to server
	s.AddResources(userResource)

	// run server
	s.Run(":8085",&server.Options{
		Timeout: 0,
	})
}
