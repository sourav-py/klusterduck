package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Use kubeconfig file from ~/.kube/config by default, or the one specified via `-kubeconfig` flag
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	configPath := flag.String("kubeconfig", kubeconfig, "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	// Build the configuration from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *configPath)
	if err != nil {
		panic(err.Error())
	}

	// Create a new clientset to interact with the Kubernetes API
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}


  ListDeploymentsInNamespace(clientset,"local");


}

// homeDir returns the home directory of the user running this program
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}

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

