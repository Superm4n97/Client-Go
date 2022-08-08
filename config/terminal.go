package config

import (
	"fmt"
	"github.com/Superm4n97/Client-Go/deployment"
	"github.com/Superm4n97/Client-Go/service"
)

func Terminal() {
	clientst := CreateNewClientSet()

	for {
		var cmd, cmd1, cmd2 string
		fmt.Print(">")

		fmt.Scanln(&cmd1, &cmd2)
		cmd = cmd1 + " " + cmd2

		if cmd1 != "" {
			var name string
			var rep int32

			if cmd == "exit " {
				break
			} else if cmd == "create deployment" {
				fmt.Scanln(&name)
				deployment.CreateDeployment(name, clientst)

			} else if cmd == "update deployment" {
				fmt.Scanln(&name, &rep)
				deployment.UpdateDeployment(name, rep, clientst)

			} else if cmd == "get deployment" {
				fmt.Scanln(&name)
				deployment.GetDeployment(name, clientst)
			} else if cmd == "get deployments" {
				deployment.GetDeploymentList(clientst)
			} else if cmd == "delete deployment" {
				fmt.Scanln(&name)
				deployment.DeleteDeployment(name, clientst)
			} else if cmd == "create service" {
				fmt.Scanln(&name)
				service.CreateNewService(name, clientst)
			} else if cmd == "get services" {
				service.GetAllServices(clientst)
			} else if cmd == "delete service" {
				fmt.Scanln(&name)
				service.DeleteService(name, clientst)
			} else {
				fmt.Println("No such command found")
			}
		}
	}
}
