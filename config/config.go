package config

import "github.com/olebedev/config"

var Config *config.Config
var TestMode bool

func LoadConfig() error {

	config, err := load()
	if err != nil {
		return err
	}
	Config = config
	return nil
}

func load() (*config.Config, error) {

	file := "config.yml"
	if TestMode {
		file = "config_test.yml"
	}

	cfg, err := config.ParseYamlFile(file)
	if err != nil {
		return nil, err
	}

	cfg = cfg.Env()

	//reemplazo por configuraciones especificas definidas en parametros
	cfg = cfg.Flag()

	return cfg, nil
}