package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"git.thwap.org/rockhopper/gout"
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

func SwapUsage(c chan string) {
	for {
		swap := getSwapTotals()
		c <- fmt.Sprintf(
			"%s: %s",
			gout.Bold(gout.White("SWAP")),
			progress(int((float64(swap.Used)/float64(swap.Total))*100.0)),
		)
		time.Sleep(time.Second)
	}
}
