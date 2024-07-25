package main

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/rconway/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Printf("%s - %s\n", config, err)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// Print pod details
	fmt.Println("----")
	for _, pod := range pods.Items {
		fmt.Printf("Pod: %s\n", pod.Name)
		fmt.Printf("Status: %s\n", pod.Status.Phase)
		fmt.Printf("Ready Containers: %d/%d\n", len(pod.Status.ContainerStatuses), len(pod.Spec.Containers))
		fmt.Printf("Restarts: %d\n", pod.Status.ContainerStatuses[0].RestartCount)
		fmt.Printf("Age: %s\n", time.Since(pod.CreationTimestamp.Time).Round(time.Second))
		fmt.Println("----")
	}
}
