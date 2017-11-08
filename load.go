// load.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"git.thwap.org/rockhopper/gout"
)

func LoadAverages() {
	t := time.Now()
	t.Format(time.RFC3339)

	fp, er := os.Open("/proc/loadavg")
	if er != nil {
		panic(er)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	tdata := strings.Split(scanner.Text(), " ")

	fmt.Printf(
		"%s: %s - %s: %s, %s\n",
		gout.Bold(gout.White("TIME")),
		t.Format("03:04PM"),
		gout.Bold(gout.White("LOAD")),
		tdata[0],
		tdata[1],
	)

}
