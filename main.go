package main

import "github.com/Superm4n97/Client-Go/config"

func main() {
	//clientset := config.CreateNewClientSet()
	//deployment.GetDeploymentList()
	//deployment.GetDeployment("book-server")
	//deployment.CreateDeployment(config.CreateNewClientSet())
	//deployment.UpdateDeployment("book-api-server", 3, config.CreateNewClientSet())
	//fmt.Println(*deployment.GetDeployment("book-server", config.CreateNewClientSet()).Spec.Replicas)

	//fmt.Println(service.GetService("kubernetes", config.CreateNewClientSet()).Namespace)
	//service.GetAllServices(config.CreateNewClientSet())
	//service.CreateNewService(config.CreateNewClientSet())
	//service.DeleteService("book-api-server", config.CreateNewClientSet())

	//ingress.GetIngress("book-api-server2", clientset)
	//ingress.GetAllIngress(clientset)
	//ingress.DeleteIngress("book-api-server2", clientset)

	config.Terminal()
}
