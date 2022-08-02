package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {

	//path to the config file in my local device
	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)

	//BuildConfigFromFlags builds configs from filepath (outside cluster) or master url
	//of the cluster (in cluster)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	//this config file contains the information about the cluster. It has the path,
	//host address, password and many more.

	//we have created configs, now we have to create a client from that config.
	//type Clientset from kubernetes package is an object that helps us to create
	//serializer client. So first we have to build the Clientset object from config.

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	//now we can use this clientset to build our api
	//clientset.CoreV1()

	fmt.Println(clientset.CoreV1())
	fmt.Println(config.Host)
	fmt.Println(config.APIPath)
}
