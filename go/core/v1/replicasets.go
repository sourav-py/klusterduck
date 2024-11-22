package core 

import (
    "context"
    "fmt"
    "k8s.io/client-go/kubernetes"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FetchReplicaSets(clientset *kubernetes.Clientset, namespace string) ([]string, error) {
    replicaSets, err := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    rsNames := make([]string, 0)
    for _, rs := range replicaSets.Items {
        rsNames = append(rsNames, rs.Name)
    }
    fmt.Println("Fetched ReplicaSets:", rsNames)
    return rsNames, nil
}

