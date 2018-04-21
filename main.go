package main

import (
	"fmt"
)

// Panic - catch critical error
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// load config data
	configPath := "config/imno.conf"

	config := NewConfig()
	err := config.Load(configPath)
	Panic(err)

	//print loaded config to stderr
	fmt.Println(config)
}
