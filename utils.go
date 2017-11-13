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
		return fmt.Sprintf("%.02fK", (float64(i) / 1024.0))
	} else if i >= int64(math.Pow(1024, 2)) && i <= int64(math.Pow(1024, 3)) {
		return fmt.Sprintf("%.02fM", (float64(i) / math.Pow(1024, 2)))
	} else if i >= int64(math.Pow(1024, 3)) && i <= int64(math.Pow(1024, 4)) {
		return fmt.Sprintf("%.02fG", (float64(i) / math.Pow(1024, 3)))
	} else if i >= int64(math.Pow(1024, 4)) && i <= int64(math.Pow(1024, 5)) {
		return fmt.Sprintf("%.02fT", (float64(i) / math.Pow(1024, 4)))
	} else {
		return fmt.Sprintf("%dB", i)
	}
}

func pcheck(e error) {
	if e != nil {
		panic(e)
	}
}

func doubleProgress(tag string, a, b int, l1, l2 string) string {
	var p1, p2 string
	// max length
	mlen := (40 - (Config.Padding.Left + Config.Padding.Right))
	// ensure we are not over 100%
	if a > 100 {
		a = 100
	}
	if b > 100 {
		b = 100
	}
	// ensure our labels are only of len() 2
	if len(l1) > 2 {
		t := l1[:2]
		l1 = t
	}
	if len(l2) > 2 {
		t := l2[:2]
		l2 = t
	}
	// and trunc or pad our tag
	if len(tag) > 7 {
		t := tag[:7]
		tag = t
	} else if len(tag) < 7 {
		t := fmt.Sprintf("%s%s", tag, strings.Repeat(" ", (7-len(tag))))
		tag = t
	}
	// bar length
	blen := ((mlen - 19) / 2)
	// bar 1
	bar1 := int(float64(blen) * (float64(a) / 100.0))
	spc1 := (blen - bar1)
	if a >= 90 {
		p1 = gout.Bold(gout.Red(strings.Repeat("#", int(bar1))))
	} else if a >= 80 {
		p1 = gout.Bold(gout.Yellow(strings.Repeat("#", int(bar1))))
	} else {
		p1 = gout.Bold(gout.Green(strings.Repeat("#", int(bar1))))
	}
	// bar 2
	bar2 := int(float64(blen) * (float64(b) / 100.0))
	spc2 := (blen - bar2)
	if b >= 90 {
		p2 = gout.Bold(gout.Red(strings.Repeat("#", int(bar2))))
	} else if b >= 80 {
		p2 = gout.Bold(gout.Yellow(strings.Repeat("#", int(bar2))))
	} else {
		p2 = gout.Bold(gout.Green(strings.Repeat("#", int(bar2))))
	}
	// and hand it off
	return fmt.Sprintf(
		"%s: %s:%s%s%s%s %s:%s%s%s%s",
		gout.Bold(gout.White(tag)),
		gout.Bold(gout.Cyan(l1)),
		gout.Yellow("["),
		p1,
		gout.Cyan(strings.Repeat("-", int(spc1))),
		gout.Yellow("]"),
		gout.Bold(gout.Cyan(l2)),
		gout.Yellow("["),
		p2,
		gout.Cyan(strings.Repeat("-", int(spc2))),
		gout.Yellow("]"),
	)
}

func progress(l string, i int) string {
	var bars, space string
	mlen := (40 - (Config.Padding.Left + Config.Padding.Right))
	// make sure we don't go over 100%
	if i > 100 {
		i = 100
	}
	// if the label is more than 15chars, trunc it, sorry, space is a premium
	if len(l) > 7 {
		t := l[:7]
		l = t
	} else if len(l) < 7 {
		t := fmt.Sprintf("%s%s", l, strings.Repeat(" ", (7-len(l))))
		l = t
	}
	// record our label
	label := fmt.Sprintf("%s: ", gout.Bold(gout.White(l)))
	// calculate the bar length
	blen := int((mlen - (len(l) + 9)))
	// calculate the number of hashes and dashes
	hash := int(float64(blen) * (float64(i) / 100.0))
	dash := int(blen - hash)
	// record our progress bar
	le := gout.Yellow("[")
	re := gout.Yellow("]")
	if i >= 90 {
		bars = gout.Bold(gout.Red(strings.Repeat("#", hash)))
	} else if i >= 80 {
		bars = gout.Bold(gout.Yellow(strings.Repeat("#", hash)))
	} else {
		bars = gout.Bold(gout.Green(strings.Repeat("#", hash)))
	}
	space = gout.Cyan(strings.Repeat("-", dash))
	// and hand off our result
	return fmt.Sprintf(
		"%s%s%s%s%s %s%%",
		label,
		le,
		bars,
		space,
		re,
		gout.Bold(gout.Green(fmt.Sprintf("%3d", i))),
	)
}
