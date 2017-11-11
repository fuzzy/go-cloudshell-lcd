// load.go
package main

import (
	"fmt"
	"time"

	"git.thwap.org/rockhopper/friday"
	"git.thwap.org/rockhopper/gout"
)

func LoadProducer() {
	f := friday.Friday{Facts: make(map[string]string)}

	for {
		t := time.Now()
		t.Format(time.RFC3339)
		f.CollectFacts()

		Output <- &CloudShellOutput{
			Type: "load",
			Lines: []string{
				fmt.Sprintf(
					"%s: %s - %s: %s, %s, %s",
					gout.Bold(gout.White("TIME")),
					t.Format("03:04PM"),
					gout.Bold(gout.White("LOAD")),
					f.Get("5minLoad"),
					f.Get("10minLoad"),
					f.Get("15minLoad"),
				),
			},
		}
		time.Sleep(5 * time.Second)
	}
}
