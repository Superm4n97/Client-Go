package main

import (
	"github.com/Superm4n97/Client-Go/config"
	"github.com/Superm4n97/Client-Go/deployment"
)

func main() {
	//deployment.GetDeploymentList()
	//deployment.GetDeployment("book-server")
	deployment.CreateDeployment(config.CreateNewClientSet())
}
