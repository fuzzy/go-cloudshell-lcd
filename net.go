// net.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"

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
	fp.Close()

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
	fp.Close()

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
			retv[string(v.Name())] = &NetIf{Name: v.Name()}
			point_a := getTransfer(v.Name())
			time.Sleep(time.Second)
			point_b := getTransfer(v.Name())

			// now get the interface speed
			tf, te := os.Open(fmt.Sprintf("/sys/class/net/%s/speed", v.Name()))
			scanner := bufio.NewScanner(tf)
			scanner.Scan()
			ts, te := strconv.ParseInt(scanner.Text(), 10, 64)
			if te != nil {
				panic(te)
			}
			retv[string(v.Name())].Speed = ts

			// now calc the difference between points a and b (speed)
			retv[string(v.Name())].Rx_b = (point_b.Rx_b - point_a.Rx_b)
			retv[string(v.Name())].Tx_b = (point_b.Tx_b - point_a.Tx_b)

			// cleanup the filehandle
			tf.Close()
		}
	}

	return retv
}

func NetUsage(c chan []string) {
	for {
		retv := []string{}
		data := interfaces()
		var rwp, twp float64
		for k, v := range data {
			if v.Rx_b > 0 {
				rwp = 100.0 * (100.0 * ((float64(v.Rx_b) / float64(v.Speed*int64(math.Pow(1024, 3)))) * 100.0))
			} else {
				rwp = 0
			}
			if v.Tx_b > 0 {
				twp = 100.0 * (100.0 * ((float64(v.Tx_b) / float64(v.Speed*int64(math.Pow(1024, 3)))) * 100.0))
			} else {
				twp = 0
			}
			retv = append(retv, fmt.Sprintf(
				"%s: %s",
				gout.Bold(gout.White(k)),
				doubleProgress(int(rwp), int(twp), "rx", "tx"),
			))
		}
		c <- retv
		time.Sleep(time.Second)
	}
}
