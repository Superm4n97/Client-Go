package config

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func DynamicClient() {
	kubeConfig := getKubeConfigFilePath()
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		fmt.Println("failed to get config from path")
		panic(err)
	}

	dyClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("failed to create dynamic client from config file")
		panic(err)
	}

	pod, err := dyClient.Resource(schema.GroupVersionResource{
		Group:    "apis/apps",
		Version:  "v1",
		Resource: "deployment",
	}).Get(context.TODO(), "rasel", metav1.GetOptions{})

	if err != nil {
		fmt.Println("failed to get the resource from dynamic client")
		panic(err)
	}

	fmt.Println("Resources found")
	fmt.Println(pod.GetName())
}
