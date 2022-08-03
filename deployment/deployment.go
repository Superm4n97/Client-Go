package deployment

import (
	"context"
	"fmt"
	"github.com/Superm4n97/Client-Go/config"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDeploymentList() {
	clientset := config.CreateNewClientSet()
	dep_lst, err := clientset.AppsV1().Deployments(v1.NamespaceDefault).List(context.TODO(), v12.ListOptions{})

	if err != nil {
		fmt.Println("failed to create deployment list")
		panic(err)
	}

	fmt.Println("Deployments:")
	for _, item := range dep_lst.Items {
		fmt.Println(item.Name)
	}
}
