package main

import (
	"fmt"
	"log"
	"os"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getPodList(namespace string, clientSet *kubernetes.Clientset) (podList *v1.PodList) {

	opts := metav1.ListOptions{}
	podList, err := clientSet.CoreV1().Pods(namespace).List(opts)

	if err != nil {
		log.Fatalln("[ERROR] Failed to list pods. Stacktrace below:")
		log.Fatal(err)
	}

	fmt.Printf("[NAMESPACE]: %s, [POD COUNT]: %v\n", namespace, len(podList.Items))
	fmt.Println("================================================")

	return podList
}

// Note to self: Comment later please Abdul - use cobra instead
func checkParams() string {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Printf("Usage: %v namespace\n", os.Args[0])
		os.Exit(1)
	} else {
		fmt.Println("Performing pods for namespace:", os.Args[1])
	}

	namespace := os.Args[1]
	return namespace
}
