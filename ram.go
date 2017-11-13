// ram.go
package main

// #include <unistd.h>
// #include <sys/sysctl.h>
// #include <sys/types.h>
import "C"

import (
	"time"
)

func RamProducer() {
	for {
		maxRam := int64(C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
		freeRam := int64(C.sysconf(C._SC_AVPHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
		usedRam := (maxRam - freeRam)
		ramPercUsed := (float64(usedRam) / float64(maxRam)) * 100.0
		Output <- &CloudShellOutput{
			Type:  "ram",
			Lines: []string{progress("Ram", int(ramPercUsed))},
		}
		time.Sleep(time.Second)
	}
}
