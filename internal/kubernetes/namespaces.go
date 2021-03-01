package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getNamespaces(context *string) (*v1.NamespaceList, error) {
	client, err := GetClient()
	if err != nil {
		return nil, err
	}

	return client.CoreV1().Namespaces().List(metav1.ListOptions{})
}

func GetNamespaceNames(context *string) ([]string, error) {
	print("namespace names " + *context)
	nsList, err := getNamespaces(context)
	if err != nil {
		return []string{}, err
	}
	var namespacesNames []string

	for _, namespace := range nsList.Items {
		namespacesNames = append(namespacesNames, namespace.Name)
	}

	return namespacesNames, nil
}
