package main

import (
	"github.com/Superm4n97/Client-Go/config"
	"github.com/Superm4n97/Client-Go/deployment"
)

func main() {
	//deployment.GetDeploymentList()
	//deployment.GetDeployment("book-server")
	//deployment.CreateDeployment(config.CreateNewClientSet())
	deployment.UpdateDeployment("book-api-server", 3, config.CreateNewClientSet())

	//fmt.Println(*deployment.GetDeployment("book-server", config.CreateNewClientSet()).Spec.Replicas)

}
