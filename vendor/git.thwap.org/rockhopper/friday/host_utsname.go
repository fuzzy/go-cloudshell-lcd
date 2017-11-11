// +build linux solaris

package friday

import (
	"runtime"
	"strings"

	"github.com/fatih/structs"

	"golang.org/x/sys/unix"
)

func _toByte(i int8) byte {
	if i < 0 {
		return byte(256 + int(i))
	} else {
		return byte(int(i))
	}
}

func _utsnameToMap(u unix.Utsname) map[string]string {
	retv := make(map[string]string)
	for k, v := range structs.Map(u) {
		switch v := v.(type) {
		case [65]uint8:
			b := []byte{}
			for _, d := range v {
				b = append(b, _toByte(int8(d)))
			}
			retv[k] = strings.TrimRight(string(b), "\x00")
		case [257]int8:
			b := []byte{}
			for _, d := range v {
				b = append(b, _toByte(d))
			}
			retv[k] = strings.TrimRight(string(b), "\x00")
		}
	}
	return retv
}

func (f *Friday) _getUnameInfo() {
	var uname unix.Utsname
	if err := unix.Uname(&uname); err != nil {
		panic(err)
	}
	for k, v := range _utsnameToMap(uname) {
		f.Add(k, v)
	}
	f.Add("Sysname_lc", runtime.GOOS)
}
