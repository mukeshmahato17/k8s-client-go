package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", clientcmd.RecommendedHomeFile, "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Pods from default namespace")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n ", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployments from default namespace")
	for _, deploy := range deployments.Items {
		fmt.Printf("%s\n", deploy.Name)
	}
}
