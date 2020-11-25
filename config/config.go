package config

import "github.com/olebedev/config"

var Config *config.Config

func LoadConfig() error {

	config, err := load()
	if err != nil {
		return err
	}
	Config = config
	return nil
}

func load() (*config.Config, error) {

	cfg, err := config.ParseYamlFile("config.yml")
	if err != nil {
		return nil, err
	}

	cfg = cfg.Env()

	//reemplazo por configuraciones especificas definidas en parametros
	cfg = cfg.Flag()

	return cfg, nil
}