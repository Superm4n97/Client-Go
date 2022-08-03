package config

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func getKubeConfigFilePath() *string {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "exact path of the config file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "exact path of the config file")
	}
	return kubeconfig
}

func CreateNewClientSet() kubernetes.Interface {

	kubeconfig := getKubeConfigFilePath()

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("can not build the config file")
		panic(err)
	}

	var clientset kubernetes.Interface
	clientset, err = kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("failed creating clientset")
		panic(err)
	}

	return clientset
}
