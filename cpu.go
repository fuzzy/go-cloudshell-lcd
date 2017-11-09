// cpu.go
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

func jiffies() []string {
	fp, er := os.Open("/proc/stat")
	if er != nil {
		panic(er)
	}

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	fp.Close()
	return strings.Split(scanner.Text(), " ")
}

func countJiffies() (int64, int64) {
	var totj int64
	var wrkj int64

	for k, v := range jiffies() {
		if k >= 2 {
			tv, er := strconv.ParseInt(v, 10, 64)
			if er != nil {
				panic(er)
			}
			totj += tv
		}
		if k >= 2 && k <= 4 {
			tv, er := strconv.ParseInt(v, 10, 64)
			if er != nil {
				panic(er)
			}
			wrkj += tv
		}
	}
	return totj, wrkj
}

func workJiffies()

func CpuUsage(c chan string) {
	for {
		tj1, wj1 := countJiffies()
		time.Sleep(50 * time.Millisecond)
		tj2, wj2 := countJiffies()

		top := (tj2 - tj1)
		wop := (wj2 - wj1)
		cpu := (float64(wop) / float64(top)) * 100.0

		c <-fmt.Sprintf("%s: %s", gout.Bold(gout.White("CPU")), progress(int(cpu)))
	}
}
