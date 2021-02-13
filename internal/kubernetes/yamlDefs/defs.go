package yamlDefs

type YamlConfig struct {
	apiVersion string
	//clusters []Cluster `yaml:",flow"`
}

type Cluster struct {
	certificateAuthorityData string `yaml:"certificate-authority-data"`
	server                   string `yaml:"server"`
	name                     string
}
