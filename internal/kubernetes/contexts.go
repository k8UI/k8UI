package kubernetes

import (
	"github.com/goccy/go-yaml"
	"io/ioutil"

	"os"
	"strings"
)

var kubeconfig *string

func GetContexts() ([]string, error) {
	if kubeconfig == nil {
		kubeconfig = GetConfigPath("contextreader")
	}

	configYaml, err := os.Open(*kubeconfig)
	if err != nil {
		return nil, err //TODO some interface for sending errors to FE?
	}

	byteValue, err := ioutil.ReadAll(configYaml)
	if err != nil {
		return nil, err
	}

	path, err := yaml.PathString("$.contexts[*].name")
	if err != nil {
		return nil, err
	}

	var names []string
	if err := path.Read(strings.NewReader(string(byteValue)), &names); err != nil {
		return nil, err
	}
	defer configYaml.Close()

	return names, nil

}
