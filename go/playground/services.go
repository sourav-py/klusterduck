package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



// List all services in the cluster
func ListServices(clientset *kubernetes.Clientset) {
	fmt.Println("Services in the cluster:")

	services, err := clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, svc := range services.Items {
		fmt.Printf("Service Name: %s, Namespace: %s, Type: %s\n", svc.Name, svc.Namespace, svc.Spec.Type)
	}
}

func ListServicesInNamespace(clientset *kubernetes.Clientset, namespace string) {
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, svc := range services.Items {
		fmt.Printf("Service Name: %s, Namespace: %s, Type: %s\n", svc.Name, svc.Namespace, svc.Spec.Type)
	}
}
