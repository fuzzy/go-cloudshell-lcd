package main

import (
	"fmt"
	"time"
)

func PrintLine(l int, c chan string) {
	fmt.Printf("\033[%d;1H%s", l, <-c)
}

func PrintLines(l int, c chan []string) int {
	for i, v := range <-c {
		fmt.Printf("\033[%d;1H%s", l, v)
		l += i
	}
	return l
}

func main() {
	load_c := make(chan string)
	cpu_c := make(chan string)
	ram_c := make(chan string)
	swap_c := make(chan string)
	net_c := make(chan []string)
	disk_c := make(chan []string)

	go LoadAverages(load_c)
	go CpuUsage(cpu_c)
	go RamUsage(ram_c)
	go SwapUsage(swap_c)
	go NetUsage(net_c)
	go DiskUsage(disk_c)

	fmt.Println("\033[2J")

	for {
		go PrintLine(2, load_c)
		go PrintLine(3, cpu_c)
		go PrintLine(4, ram_c)
		go PrintLine(5, swap_c)
		go PrintLines(6, net_c)
		go PrintLines(8, disk_c)
		time.Sleep(2 * time.Second)
	}

}
