package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//TODO: get client in order

func GetNamespacesForContext(context string) *corev1.NamespaceList {
	list, err := GetClient().CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	return list
}
