package main

import (
	"context"
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func NewClientset() kubernetes.Interface {
	var kubeconfig *string

	//Generating the path of the config file
	if home := homedir.HomeDir(); home != "" {
		//for outside the cluster
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "absolute path to the config file")
	} else {
		//for inside the cluster
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the config file")
	}
	//fmt.Println(*kubeconfig)

	// flag.Parse parses the command-line flags from os.Args[1:]. Must be called
	// after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	//creating clientset.
	//clientset is a kubernetes.interface type
	//var clientset kubernetes.Interface
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

func main() {
	clientset := NewClientset()

	fmt.Println("clientset created")

	podlst, err := clientset.CoreV1().Pods(corev1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, item := range podlst.Items {
		fmt.Print(item.Name)
		fmt.Println("\t", item.Namespace)
	}

	//fmt.Println(podlst.Items[0])
}
