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

	//ListPodsInNamespace(clientset, "kube-system")
	ListServices(clientset)

}

// homeDir returns the home directory of the user running this program
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}



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
