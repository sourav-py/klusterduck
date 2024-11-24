package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



//List all the pods in the cluster
func ListDeployments(clientset *kubernetes.Clientset) {

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	itr := 1
	for _, deployment := range deployments.Items {
		//Print Name and Namespace of the pod
    fmt.Printf("%d Deployment name: %s, Namesapce: %s\n", itr, deployment.Name, deployment.Namespace)
		itr += 1
	}

}

// List all the pods in a namespace
func ListDeploymentsInNamespace(clientset *kubernetes.Clientset, namespace string) {

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	itr := 1
	for _, deployment := range deployments.Items {
		fmt.Printf("%d Deployment name: %s, Replicas: %v\n", itr, deployment.Name, *deployment.Spec.Replicas)
		itr += 1
	}
}



