package ingress

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func emptyIngressTemplate(ingsName string) v12.Ingress {
	return v12.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: ingsName,
		},
		Spec: v12.IngressSpec{
			Rules: []v12.IngressRule{
				{
					Host:             "book-api.localhost.test",
					IngressRuleValue: v12.IngressRuleValue{
						//HTTP: field.Path{
						//},
					},
				},
			},
		},
	}
}

func CreateIngress(clientset kubernetes.Interface) {
	emptyIngressTemplate("book-api-server")
}

func GetIngress(ingressName string, clientset kubernetes.Interface) {
	res, err := clientset.NetworkingV1().Ingresses(v1.NamespaceDefault).Get(context.TODO(), ingressName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("failed to get ingress ", ingressName)
		panic(err)
	}

	fmt.Println(res.Name)
}

func GetAllIngress(clientst kubernetes.Interface) {
	ingList, err := clientst.NetworkingV1().Ingresses(v1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("failed to get the ingress list")
		panic(err)
	}

	for _, item := range ingList.Items {
		fmt.Println(item.Name)
	}
}

func DeleteIngress(ingressName string, clientset kubernetes.Interface) {
	err := clientset.NetworkingV1().Ingresses(v1.NamespaceDefault).Delete(context.TODO(), ingressName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("failed to delete ingress ", ingressName)
		panic(err)
	}

	fmt.Println("ingress ", ingressName, " deleted")
}
