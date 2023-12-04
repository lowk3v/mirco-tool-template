package config

import (
	"github.com/lowk3v/micro-tool-template/pkg/log"
	"gopkg.in/yaml.v3"
	"os"
)

var Config Yaml
var Log *log.Logger

type Yaml struct {
	// Your config load from config.yml here
}

func NewConfig(cfgPath string) error {
	Log = log.New("debug")
	Config = Yaml{}

	// Open config file
	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&Config); err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
