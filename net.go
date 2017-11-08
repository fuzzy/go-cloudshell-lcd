// net.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"git.thwap.org/rockhopper/gout"
)

type NetIf struct {
	Name       string
	Speed      int64
	Rx_b, Tx_b int64
}

func getTransfer(i string) *NetIf {
	fp, er := os.Open(fmt.Sprintf("/sys/class/net/%s/statistics/rx_bytes", i))
	if er != nil {
		panic(er)
	}
	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	rx, er := strconv.ParseInt(scanner.Text(), 10, 64)
	if er != nil {
		panic(er)
	}

	fp, er = os.Open(fmt.Sprintf("/sys/class/net/%s/statistics/tx_bytes", i))
	if er != nil {
		panic(er)
	}
	scanner = bufio.NewScanner(fp)
	scanner.Scan()
	tx, er := strconv.ParseInt(scanner.Text(), 10, 64)
	if er != nil {
		panic(er)
	}

	retv := &NetIf{Name: i, Rx_b: rx, Tx_b: tx}
	return retv
}

func interfaces() map[string]*NetIf {
	retv := make(map[string]*NetIf)
	f, e := ioutil.ReadDir("/sys/class/net")
	if e != nil {
		panic(e)
	}
	for _, v := range f {
		if string(v.Name()[0]) == "e" {
			retv[string(v.Name())] = getTransfer(v.Name())
			// now get the interface speed
			tf, te := os.Open(fmt.Sprintf("/sys/class/net/%s/speed", v.Name()))
			scanner := bufio.NewScanner(tf)
			scanner.Scan()
			ts, te := strconv.ParseInt(scanner.Text(), 10, 64)
			if te != nil {
				panic(te)
			}
			retv[string(v.Name())].Speed = ts
		}
	}

	return retv
}

func NetUsage() {
	data := interfaces()
	for k, v := range data {
		fmt.Printf(
			"%s: %s %s - %s %s\n",
			gout.Bold(gout.White(k)),
			gout.Yellow("rx"),
			gout.Bold(gout.Green(humanSize(v.Rx_b))),
			gout.Yellow("tx"),
			gout.Bold(gout.Green(humanSize(v.Tx_b))),
		)
	}
}
