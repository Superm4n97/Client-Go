package service

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

//creates a dummy template for services for given data, it returns a service type
func emptyServiceTemplate(svcName, lb1, lb2 string, exposePort, podPort int32) v1.Service {
	return v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: svcName,
			Labels: map[string]string{
				lb1: lb2,
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Port: exposePort,
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: podPort,
					},
				},
			},
			Selector: map[string]string{
				lb1: lb2,
			},
		},
	}
}

func CreateNewService(svcName string, clientst kubernetes.Interface) {
	svc := emptyServiceTemplate(svcName, "app", "book-api-server", 9090, 8080)

	res, err := clientst.CoreV1().Services(v1.NamespaceDefault).Create(context.TODO(), &svc, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("failed to create service")
		panic(err)
	}

	fmt.Println(res.Name, " service created")
}

func GetService(svcName string, clientset kubernetes.Interface) v1.Service {
	svc, err := clientset.CoreV1().Services(v1.NamespaceDefault).Get(context.TODO(), svcName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("failed to get service")
		panic(err)
	}

	return *svc
}

func GetAllServices(clientset kubernetes.Interface) {

	svcList, err := clientset.CoreV1().Services(v1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("failed to list services")
		panic(err)
	}

	for _, item := range svcList.Items {
		fmt.Println(item.Name, " ", item.Spec.Ports)
	}
}

func DeleteService(svcName string, clientset kubernetes.Interface) {
	err := clientset.CoreV1().Services(v1.NamespaceDefault).Delete(context.TODO(), svcName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("failed to delete the service ", svcName)
		panic(err)
	}

	fmt.Println("service ", svcName, " deleted")
}
