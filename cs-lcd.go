package main

import (
	"fmt"
	"strings"
	"time"
)

type CloudShellOutput struct {
	Type  string
	Lines []string
}

var Config *CloudShellConfig
var Output chan *CloudShellOutput

func Outputter() {
	// clear the screen
	fmt.Printf("\033[2J")

	for {
		// define our default starting point
		start := Config.Padding.Top

		// Get a line of output
		tmp := <-Output

		// stuff that should output before net/disks
		for k, v := range Config.outputs {
			if tmp.Type == v {
				start += (k + 1)
				fmt.Printf("\033[%d;1H%s", start, strings.Repeat(" ", Config.Padding.Left))
				for _, v := range tmp.Lines {
					fmt.Println(v)
				}
			}
		}

		// net output
		for _, pt := range []string{"et", "en", "wl"} {
			if tmp.Type[:2] == pt {
				for i, n := range Config.Outputs.Net {
					if n.Name == tmp.Type && n.Enabled {
						fmt.Printf(
							"\033[%d;1H%s%s",
							start+len(Config.outputs)+i+2,
							strings.Repeat(" ", Config.Padding.Left),
							tmp.Lines[0],
						)
					}
				}
			}
		}

		// disk output
		for _, pt := range []string{"sd"} {
			if tmp.Type[:2] == pt {
				cnt := 0
				for _, tv := range Config.Outputs.Net {
					if tv.Enabled {
						cnt += 1
					}
				}
				for i, n := range Config.Outputs.Disk {
					if n.Name == tmp.Type && n.Enabled {
						fmt.Printf(
							"\033[%d;1H%s%s",
							start+len(Config.outputs)+cnt+i+3,
							strings.Repeat(" ", Config.Padding.Left),
							tmp.Lines[0],
						)
					}
				}
			}
		}

		fmt.Printf("\033[0;1H")
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	Config = ReadConfig("/etc/go-cloudshell.yml")
	Output = make(chan *CloudShellOutput, 7)

	// Producer startup
	if Config.Outputs.Host {
		go HostProducer()
	}
	if Config.Outputs.Load {
		go LoadProducer()
	}
	if Config.Outputs.Cpu {
		go CpuProducer()
	}
	if Config.Outputs.Ram {
		go RamProducer()
	}
	if Config.Outputs.Swap {
		go SwapProducer()
	}
	go NetProducer()
	go DiskProducer()

	fmt.Printf("%+v\n", Config)
	Outputter()
}
