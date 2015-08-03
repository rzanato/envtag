package main

import (
	"fmt"
	"io/ioutil"

	"github.com/rzanato/envtag"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Foo  string `env:"FOO"`
	Bar  string `yaml:"bar" env:"BAR"`
	Misc string `yaml:"misc"`
}

func main() {
	data, err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		panic(err)
	}

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	err = envtag.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)
}
