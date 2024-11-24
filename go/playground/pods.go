package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)



func ListPods(clientset *kubernetes.Clientset, showLabels bool, showOwnerRef bool) {
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
		
		if showLabels == true {
			//Print labels for the pod
			for label, value := range pod.Labels {
				fmt.Printf("\t%s:%s\n", label, value)
			}

		}

    if showOwnerRef == true {
      ownerReferences := pod.OwnerReferences;
      fmt.Printf("\t%v\n",ownerReferences) 
      /*
      for _, ownerRef := range ownerReferences.Items {
        fmt.Printf("\t%s \n", ownerRef.Name)
      }
      */
    }
		itr += 1
	}

}

// List all the pods in a namespace
func ListPodsInNamespace(clientset *kubernetes.Clientset, namespace string, showLabels bool, showOwnerRef bool) {
	fmt.Printf("Pods in the '%s' namespace\n", namespace)

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	itr := 1
	for _, pod := range pods.Items {
		fmt.Printf("%d Pod name: %s, Namespace: %s\n", itr, pod.Name, pod.Namespace)

    if showOwnerRef == true {
      ownerReferences := pod.OwnerReferences;
      fmt.Printf("\t%v\n",ownerReferences) 
      /*
      for _, ownerRef := range ownerReferences.Items {
        fmt.Printf("\t%s \n", ownerRef.Name)
      }
      */
    }

		itr += 1
	}
}

