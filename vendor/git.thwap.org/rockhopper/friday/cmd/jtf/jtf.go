package main

import "git.thwap.org/rockhopper/friday"

func main() {
	f := friday.Friday{Facts: make(map[string]string)}
	f.CollectFacts()
	f.Print()
}
