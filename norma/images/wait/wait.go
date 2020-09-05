package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"time"
)

func main() {
	ns := os.Getenv("NAMESPACE")
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ok := false
	for !ok {
		job, err := clientset.BatchV1().Jobs(ns).Get("install-db", metav1.GetOptions{})
		if err == nil {
			if job.Status.Succeeded > 0 {
				ok = true
				break
			} else {
				println(fmt.Sprintf("install-db job in %s state, waiting", job.Status.String()))
			}
		} else {
			println(fmt.Sprintf("Error getting install-db state %s", err.Error()))
		}
		time.Sleep(5 * time.Second)
	}
}
