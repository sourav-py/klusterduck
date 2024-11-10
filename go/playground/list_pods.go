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
	ListPods(clientset,false)

}

// homeDir returns the home directory of the user running this program
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}

func ListPods(clientset *kubernetes.Clientset, labels bool) {
	//List all the pods in the cluster
	fmt.Println("Pods in the cluster: ")

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	itr := 1
	for _, pod := range pods.Items {
		//Print Name and Namespace of the pod
		fmt.Printf("%d Pod name: %s, Namesapce: %s\n", itr, pod.Name, pod.Namespace)
		
		if labels == true {
			//Print labels for the pod
			for label, value := range pod.Labels {
				fmt.Printf("\t%s:%s\n", label, value)
			}

		}
		itr += 1
	}

}

// List all the pods in a namespace
func ListPodsInNamespace(clientset *kubernetes.Clientset, namespace string) {
	fmt.Printf("Pods in the '%s' namespace\n", namespace)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	itr := 1
	for _, pod := range pods.Items {
		fmt.Printf("%d Pod name: %s, Namespace: %s\n", itr, pod.Name, pod.Namespace)
		itr += 1
	}
}

