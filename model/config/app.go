package config

type App struct {
	Port    int16  `mapstructure:"port"`
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Host    string `mapstuctures:"host"`
	Version string `mapstructure:"version"`
}
