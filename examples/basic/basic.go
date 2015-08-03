package main

import (
	"fmt"

	"github.com/rzanato/envtag"
)

type Config struct {
	Location string `env:"LOCATION"`
	Home     string `env:"HOME"`
}

func main() {
	config := Config{}

	err := envtag.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	fmt.Println("HOME:", config.Home)
	fmt.Println("LOCATION:", config.Location)
}
