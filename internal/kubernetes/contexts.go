package kubernetes

import (
	"flag"

	"github.com/goccy/go-yaml"
	"io/ioutil"

	"os"
	"path/filepath"

	"strings"
)

var kubeconfig *string

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func GetContexts() []string {
	if kubeconfig == nil {
		if home := homeDir(); home != "" {
			kubeconfig = flag.String("contextreader", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("contextreader", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()
	}

	configYaml, err := os.Open(*kubeconfig)
	if err != nil {
		return nil //TODO some interface for sending errors to FE?
	}

	byteValue, err := ioutil.ReadAll(configYaml)
	if err != nil {
		return nil
	}

	path, err := yaml.PathString("$.contexts[*].name")
	if err != nil {
		return nil
	}

	var names []string
	if err := path.Read(strings.NewReader(string(byteValue)), &names); err != nil {
		return nil
	}
	defer configYaml.Close()

	return names

}
