package configs

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	CentralLatitude   float64 `yaml:"central_latitude"`
	CentralLongitude  float64 `yaml:"central_longitude"`
	PathToNamesData   string  `yaml:"path_to_names_data"`
	Host              string  `yaml:"host"`
	Radius            int     `yaml:"radius"`
	Port              int     `yaml:"port"`
	MinSecSleepDriver int     `yaml:"min_sec_sleep_driver"`
	MaxSecSleepDriver int     `yaml:"max_sec_sleep_driver"`
	MinSecSleepOrder  int     `yaml:"min_sec_sleep_order"`
	MaxSecSleepOrder  int     `yaml:"max_sec_sleep_order"`
}

func New(filename string) (cfg *Config, err error) {
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return
	}

	return
}
