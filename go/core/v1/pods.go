package core

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FetchPods(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
    pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    podNames := make([]string, 0)
    for _, pod := range pods.Items {
        podNames = append(podNames, pod.Name)
    }
    fmt.Println("Fetched Pods:", podNames)
    return podNames, nil
}

