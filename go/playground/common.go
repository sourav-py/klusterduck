package main

import (
	"flag"
	"os"
	"path/filepath"

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



