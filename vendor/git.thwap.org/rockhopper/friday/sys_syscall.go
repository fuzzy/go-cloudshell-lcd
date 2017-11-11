// +build linux solaris

package friday

import (
	"fmt"
	"syscall"
)

func (f *Friday) _getSysInfo() {
	var sinfo syscall.Sysinfo_t
	if err := syscall.Sysinfo(&sinfo); err != nil {
		panic(err)
	}
	f.Add("Totalram", fmt.Sprint(sinfo.Totalram))
	f.Add("Freeram", fmt.Sprint(sinfo.Freeram))
	f.Add("Totalswap", fmt.Sprint(sinfo.Totalswap))
	f.Add("Freeswap", fmt.Sprint(sinfo.Freeswap))
	f.Add("Uptime", fmt.Sprint(sinfo.Uptime))
	for i, v := range []string{"5minLoad", "10minLoad", "15minLoad"} {
		f.Add(v, fmt.Sprintf("%.02f", float64(sinfo.Loads[i])/65536.0))
	}
}
