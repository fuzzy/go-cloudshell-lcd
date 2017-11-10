package main

import (
	"fmt"
	"time"
)

func PrintLine(l int, c chan string) {
	fmt.Printf("\033[%d;1H%s", l, <-c)
}

func main() {
	load_c := make(chan string)
	cpu_c := make(chan string)
	ram_c := make(chan string)
	swap_c := make(chan string)
	net_c := make(chan string)

	go LoadAverages(load_c)
	go CpuUsage(cpu_c)
	go RamUsage(ram_c)
	go SwapUsage(swap_c)
	go NetUsage(net_c)

	fmt.Println("\033[2J")

	for {
		go PrintLine(1, load_c)
		go PrintLine(2, cpu_c)
		go PrintLine(3, ram_c)
		go PrintLine(4, swap_c)
		go PrintLine(5, net_c)
		time.Sleep(2 * time.Second)
	}

}
