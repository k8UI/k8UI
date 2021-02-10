package config

type Config struct {
	WindowWidth   int `envconfig:"default=640"`
	WindowsHeight int `envconfig:"default=460"`
}
