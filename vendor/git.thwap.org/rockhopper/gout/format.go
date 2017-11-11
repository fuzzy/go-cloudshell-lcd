package gout

import (
	"fmt"
	"math"
	"strings"
)

func Strappend(p, a string) string {
	return strings.Join([]string{p, a}, "")
}

func HumanTimeParse(s int64) map[string]int64 {
	smap := map[string]int64{
		"sec": 1,
		"min": 60,
		"hr":  (60 * 60),
	}
	keys := []string{"hr", "min", "sec"}
	retv := map[string]int64{}
	for _, value := range keys {
		if s >= smap[value] {
			val := (s / smap[value])
			tgt := (val * smap[value])
			rmn := (s - tgt)
			retv[value] = val
			s = rmn
		}
	}
	return retv
}

func HumanTimeColon(s int64) string {
	data := HumanTimeParse(s)
	return fmt.Sprintf("%02d:%02d:%02d", data["hr"], data["min"], data["sec"])
}

func HumanTimeConcise(s int64) string {
	retv := ""
	data := HumanTimeParse(s)
	keys := []string{"hr", "min", "sec"}
	for _, val := range keys {
		if data[val] > 0 {
			retv = Strappend(retv, fmt.Sprintf("%02d%s", data[val], string(val[0])))
		}
	}
	return retv
}

func HumanSize(s int64) string {
	smap := map[string]int64{
		"Byte": int64(1),
		"Kilo": int64(1024),
		"Mega": int64(math.Pow(float64(1024), float64(2))),
		"Giga": int64(math.Pow(float64(1024), float64(3))),
		"Tera": int64(math.Pow(float64(1024), float64(4))),
		"Peta": int64(math.Pow(float64(1024), float64(5))),
		"Exxa": int64(math.Pow(float64(1024), float64(6))),
	}
	keys := []string{"Exxa", "Peta", "Tera", "Giga", "Mega", "Kilo", "Byte"}
	for index, value := range keys {
		if s >= smap[value] {
			bigB := "B"
			if index == len(keys)-1 {
				bigB = ""
				return fmt.Sprintf("%d%s%s",
					s,
					string(value[0]),
					bigB)
			}
			return fmt.Sprintf("%.02f%s%s",
				float64(s)/float64(smap[value]),
				string(value[0]),
				bigB)
		}
	}
	return ""
}
