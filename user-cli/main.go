package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	 pb "user-service/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Flags(
				cli.StringFlag{
					Name: "name",
					Usage: "Your Name",
				},
				cli.StringFlag{
					Name: "email",
					Usage: "Your Email",
				},
				cli.StringFlag{
					Name:"password",
					Usage:"Your Password",
				},

			),
	)


}
