package kubernetes

import (
	"flag"
	"os"
	"path/filepath"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func GetConfigPath(flagName string) *string {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String(flagName, filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String(flagName, "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return kubeconfig
}
