// utils.go
package main

import (
	"fmt"
	"math"
	"strings"

	"git.thwap.org/rockhopper/gout"
)

func humanSize(i int64) string {
	if i >= 1024 && i <= int64(math.Pow(1024, 2)) {
		return fmt.Sprintf("%.02fKB", (float64(i) / 1024.0))
	} else if i >= int64(math.Pow(1024, 2)) && i <= int64(math.Pow(1024, 3)) {
		return fmt.Sprintf("%.02fMB", (float64(i) / math.Pow(1024, 2)))
	} else if i >= int64(math.Pow(1024, 3)) && i <= int64(math.Pow(1024, 4)) {
		return fmt.Sprintf("%.02fGB", (float64(i) / math.Pow(1024, 3)))
	} else if i >= int64(math.Pow(1024, 4)) && i <= int64(math.Pow(1024, 5)) {
		return fmt.Sprintf("%.02fTB", (float64(i) / math.Pow(1024, 4)))
	} else {
		return fmt.Sprintf("%dB", i)
	}
}

func progress(i int) string {
	max := 20
	bars := float64(max) * (float64(i) / 100.0)
	spcs := (max - int(bars))
	if bars >= 18 {
		return fmt.Sprintf(
			"%s%s%s%s %s%%",
			gout.Yellow("["),
			gout.Bold(gout.Red(strings.Repeat("#", int(bars)))),
			strings.Repeat(" ", int(spcs)),
			gout.Yellow("]"),
			gout.Bold(gout.Red(fmt.Sprintf("%3d", i))),
		)
	} else if bars >= 16 {
		return fmt.Sprintf(
			"%s%s%s%s %s%%",
			gout.Yellow("["),
			gout.Bold(gout.Yellow(strings.Repeat("#", int(bars)))),
			strings.Repeat(" ", int(spcs)),
			gout.Yellow("]"),
			gout.Bold(gout.Yellow(fmt.Sprintf("%3d", i))),
		)
	} else {
		return fmt.Sprintf(
			"%s%s%s%s %s%%",
			gout.Yellow("["),
			gout.Bold(gout.Green(strings.Repeat("#", int(bars)))),
			strings.Repeat(" ", int(spcs)),
			gout.Yellow("]"),
			gout.Bold(gout.Green(fmt.Sprintf("%3d", i))),
		)
	}
}
