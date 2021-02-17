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

	configYaml, _ := os.Open(*kubeconfig)
	byteValue, _ := ioutil.ReadAll(configYaml)

	path, err := yaml.PathString("$.contexts[*].name")
	if err != nil {
		return nil
	}
	s := string(byteValue)

	var names []string
	if err := path.Read(strings.NewReader(s), &names); err != nil {
		return nil
	}
	defer configYaml.Close()

	return names

}
