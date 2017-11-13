// cpu.go
package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
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

func CpuProducer() {
	for {
		tj1, wj1 := countJiffies()
		time.Sleep(time.Second)
		tj2, wj2 := countJiffies()

		top := (tj2 - tj1)
		wop := (wj2 - wj1)
		cpu := (float64(wop) / float64(top)) * 100.0

		Output <- &CloudShellOutput{
			Type:  "cpu",
			Lines: []string{progress("Cpu", int(cpu))},
		}
	}
}
