package main

import (
	"fmt"
	"time"

	"git.thwap.org/rockhopper/friday"
	"git.thwap.org/rockhopper/gout"
)

func HostProducer() {
	f := friday.Friday{Facts: make(map[string]string)}
	f.CollectFacts()

	for {
		Output <- &CloudShellOutput{
			Type: "host",
			Lines: []string{
				fmt.Sprintf(
					"%s - %s - %s",
					gout.Bold(gout.White(f.Get("Nodename"))),
					gout.Bold(gout.White(f.Get("Sysname"))),
					gout.Bold(gout.Cyan(f.Get("Release"))),
				),
			},
		}
		time.Sleep(300 * time.Second)
	}
}
