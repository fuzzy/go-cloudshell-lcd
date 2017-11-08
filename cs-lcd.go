package main

// #include <unistd.h>
// #include <sys/sysctl.h>
// #include <sys/types.h>
import "C"

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
	return fmt.Sprintf(
		"%s%s%s%s",
		gout.Yellow("["),
		gout.Bold(gout.Cyan(strings.Repeat("#", int(bars)))),
		strings.Repeat(" ", int(spcs)),
		gout.Yellow("]"),
	)
}

func main() {	
	maxRam := int64(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE))
	freeRam := int64(C.sysconf(C._SC_AVPHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE))
	usedRam := (maxRam - freeRam)
	ramPercUsed := (float64(usedRam) / float64(maxRam)) * 100.0
	fmt.Printf("%s: %s %.02f%%\n", gout.Bold(gout.White("RAM")), progress(int(ramPercUsed)), ramPercUsed)

}
