package main

import (
	"github.com/sparrc/go-ping"
)

// Ping - ping host and return stat
func Ping(host string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return nil, err
	}

	pinger.Count = 5
	pinger.Run()
	stats := pinger.Statistics()

	return stats, nil
}
