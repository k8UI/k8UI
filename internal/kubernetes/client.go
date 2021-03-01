package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var client *kubernetes.Clientset

func getConfig(configPath *string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", *configPath)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func getClient(config *rest.Config) *kubernetes.Clientset {
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return client
}

func GetClient() *kubernetes.Clientset {
	return getClient(getConfig(GetConfigPath("kubeconfig")))
}
