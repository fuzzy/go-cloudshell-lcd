package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type SwapTotals struct {
	Total int64
	Used  int64
	Free  int64
}

func getSwapTotals() *SwapTotals {
	retv := &SwapTotals{}
	fp, er := os.Open("/proc/swaps")
	pcheck(er)
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		tkn := 0
		for _, v := range strings.Split(scanner.Text(), " ") {
			if len(v) >= 1 {
				if tkn == 2 {
					tsw, er := strconv.ParseInt(v, 10, 64)
					pcheck(er)
					retv.Total += tsw
				} else if tkn == 3 {
					tsu, er := strconv.ParseInt(v, 10, 64)
					pcheck(er)
					retv.Used += tsu
				}
				tkn += 1
			}
		}
	}
	retv.Free = (retv.Total - retv.Used)
	return retv
}

func SwapProducer() {
	for {
		swap := getSwapTotals()
		Output <- &CloudShellOutput{
			Type: "swap",
			Lines: []string{
				progress(
					"Swap",
					int((float64(swap.Used)/float64(swap.Total))*100.0),
				),
			},
		}
		time.Sleep(time.Second)
	}
}
