package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"git.thwap.org/rockhopper/gout"
)

type Disk struct {
	Major           int
	Minor           int
	Name            string
	ReadsCompleted  int
	ReadsMerged     int
	SectorsRead     int
	ReadTime        int64
	WritesCompleted int
	WritesMerged    int
	SectorsWritten  int
	WriteTime       int64
	CurrentIops     int
	IoTime          int64
	WeightedIoTime  int64
	SizeTotal       uint64
	SizeFree        uint64
	SizeUsed        uint64
}

func d_isIn(a int, b []int) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func parseDiskStats() []*Disk {
	retv := []*Disk{}
	stat := &syscall.Statfs_t{}

	fp, er := os.Open("/proc/diskstats")
	pcheck(er)
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		elem := &Disk{}
		idx := 0
		for _, v := range strings.Split(scanner.Text(), " ") {
			if len(v) > 0 {
				value := int64(9999)
				var er error
				if idx != 2 {
					value, er = strconv.ParseInt(v, 10, 64)
					pcheck(er)
				}
				if value == 9999 {
					if idx == 2 {
						elem.Name = v
					}
				} else {
					if idx == 0 {
						elem.Major = int(value)
					} else if idx == 1 {
						elem.Minor = int(value)
					} else if idx == 3 {
						elem.ReadsCompleted = int(value)
					} else if idx == 4 {
						elem.ReadsMerged = int(value)
					} else if idx == 5 {
						elem.SectorsRead = int(value)
					} else if idx == 6 {
						elem.ReadTime = value
					} else if idx == 7 {
						elem.WritesCompleted = int(value)
					} else if idx == 8 {
						elem.WritesMerged = int(value)
					} else if idx == 9 {
						elem.SectorsWritten = int(value)
					} else if idx == 10 {
						elem.WriteTime = value
					} else if idx == 11 {
						elem.CurrentIops = int(value)
					} else if idx == 12 {
						elem.IoTime = value
					} else if idx == 13 {
						elem.WeightedIoTime = value
					}
				}
				idx += 1
				value = 9999
			}
		}
		// get our filesystem stats
		for _, tv := range Config.Outputs.Disk {
			if tv.Name == elem.Name {
				syscall.Statfs(fmt.Sprintf(tv.Mount), stat)
				elem.SizeTotal = (stat.Blocks * uint64(stat.Bsize))
				elem.SizeFree = (stat.Bfree * uint64(stat.Bsize))
				elem.SizeUsed = (elem.SizeTotal - elem.SizeFree)
			}
		}

		// and append our element
		retv = append(retv, elem)
	}
	return retv
}

func DiskProducer() {
	// get regex pattern ready
	diskRegex, er := regexp.Compile(`^sd[a-z]$`)
	pcheck(er)
	// let's assume a maximum average throughput of 50MB/sec
	max := 50 * math.Pow(1024, 2)

	// no do the things
	for {
		retv := &CloudShellOutput{Lines: []string{}}
		snap1 := parseDiskStats()
		time.Sleep(time.Second)
		snap2 := parseDiskStats()

		// iterate through first list for sd devices
		for _, v := range snap1 {
			if diskRegex.MatchString(v.Name) {
				time.Sleep(2 * time.Second)
				// see if this entry has a matching one in snap2
				for _, iv := range snap2 {
					if v.Name == iv.Name {
						rdSt := (iv.SectorsRead - v.SectorsRead) * 512
						rdPc := (float64(rdSt) / float64(max)) * 100.0
						wrSt := (iv.SectorsWritten - v.SectorsWritten) * 512
						wrPc := (float64(wrSt) / float64(max)) * 100.0
						retv.Lines = append(
							retv.Lines,
							fmt.Sprintf(
								"%s:  %s",
								gout.Bold(gout.White(v.Name)),
								doubleProgress(int(rdPc), int(wrPc), "rd", "wr"),
							),
						)
					}
				}
				retv.Type = v.Name
				// now get disk used percentage
				dup := (float64(v.SizeUsed) / float64(v.SizeTotal)) * 100.0
				retv.Lines = append(
					retv.Lines,
					fmt.Sprintf(
						"%s:  %s",
						gout.Bold(gout.White(v.Name)),
						progress(int(dup)),
					),
				)
			}
		}
		Output <- retv
	}
}
