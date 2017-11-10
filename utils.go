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

func pcheck(e error) {
	if e != nil {
		panic(e)
	}
}

func doubleProgress(a, b int, l1, l2 string) string {
	if a > 100 {
		a = 100
	}
	if b > 100 {
		b = 100
	}
	max := ((32 - (len(l1) + len(l2) + 7)) / 2)
	b1 := float64(max) * (float64(a) / 100.0)
	b2 := float64(max) * (float64(b) / 100.0)
	s1 := (max - int(b1))
	s2 := (max - int(b2))

	return fmt.Sprintf(
		"%s:%s%s%s%s %s:%s%s%s%s",
		gout.Bold(gout.Cyan(l1)),
		gout.Yellow("["),
		gout.Bold(gout.Green(strings.Repeat("#", int(b1)))),
		strings.Repeat(" ", int(s1)),
		gout.Yellow("]"),
		gout.Bold(gout.Cyan(l2)),
		gout.Yellow("["),
		gout.Bold(gout.Green(strings.Repeat("#", int(b2)))),
		strings.Repeat(" ", int(s2)),
		gout.Yellow("]"),
	)
}

func progress(i int) string {
	if i > 100 {
		i = 100
	}
	max := 25
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
