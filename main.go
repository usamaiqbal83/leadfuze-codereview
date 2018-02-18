package main

import (
	"math/rand"
	"time"

	"github.com/usamaiqbal83/leadfuze-codereview/resources/users"
	"github.com/usamaiqbal83/leadfuze-codereview/server"
	"github.com/usamaiqbal83/leadfuze-codereview/services"
	"github.com/usamaiqbal83/leadfuze-codereview/client"
)

func main(){
	rand.Seed(time.Now().UnixNano())

	s := server.NewServer(&server.Config{
		// some configurations will be added in future
	})

	// web client
	webClient := client.NewWebClient()

	// email verification service
	service := services.NewService(&services.Options{WebClient: webClient, Key:"tmpkey321"})

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
