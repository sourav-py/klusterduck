package core 

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FetchDeployments(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
    deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    deploymentNames := make([]string, 0)
    for _, deployment := range deployments.Items {
        deploymentNames = append(deploymentNames, deployment.Name)
    }
    fmt.Println("Fetched Deployments:", deploymentNames)
    return deploymentNames, nil
}

