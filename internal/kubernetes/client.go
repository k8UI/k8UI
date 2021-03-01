package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var client *kubernetes.Clientset

func getConfig(configPath *string) (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *configPath)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getClient(config *rest.Config) (*kubernetes.Clientset, error) {
	var err error

	if client == nil {
		client, err = kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

func GetClient() (*kubernetes.Clientset, error) {
	config, err := getConfig(GetConfigPath("kubeconfig"))

	if err != nil {
		return nil, err
	}

	return getClient(config)
}
