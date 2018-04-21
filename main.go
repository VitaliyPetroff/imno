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

	//check (ping) hosts one by one
	for _, v := range config.Target.Host {
		stats, err := Ping(v)
		if err != nil {
			fmt.Println(err)
		} else {
			if stats.PacketsSent-stats.PacketsRecv == 0 {
				fmt.Println(stats.Addr, ": good (", stats.AvgRtt, ")")
			} else {
				fmt.Println(stats.Addr, ": BAD (", stats.PacketLoss, " packets lost)")
			}
		}
	}
}
