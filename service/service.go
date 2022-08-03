package service

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetService(svcName string, clientset kubernetes.Interface) v1.Service {
	svc, err := clientset.CoreV1().Services(v1.NamespaceDefault).Get(context.TODO(), svcName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("failed to get service")
		panic(err)
	}

	return *svc
}
