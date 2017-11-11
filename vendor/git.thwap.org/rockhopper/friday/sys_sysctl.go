// +build freebsd openbsd netbsd

package friday

import (
	"os/exec"
	"regexp"
	"strings"
)

func (f *Friday) _sysctl(oid string) string {
	if out, err := exec.Command("/sbin/sysctl", oid).Output(); err == nil {
		fbsd, _ := regexp.Compile("^([a-zA-Z0-9\\._%\\-]*)(:\\ )(.*)$")
		obsd, _ := regexp.Compile("^([a-zA-Z0-9\\._\\-]*)(=)(.*)$")
		nbsd, _ := regexp.Compile("^([a-zA-Z0-9\\._]+)(\\ =\\ )(.*)")
		for _, v := range []*regexp.Regexp{fbsd, obsd, nbsd} {
			if v.MatchString(strings.TrimSpace(string(out))) {
				data := v.FindAllStringSubmatch(strings.TrimSpace(string(out)), -1)
				return data[0][len(data[0])-1]
			}
		}
		return ""
	}
	return ""
}

func (f *Friday) _getSysInfo() {
	var data map[string]string
	var load []string

	data = map[string]string{
		"Totalmem": "hw.physmem",
		"Usermem":  "hw.usermem",
		"Realmem":  "hw.realmem",
	}

	switch m := f.Get("Sysname"); m {
	case "FreeBSD":
		load = strings.Split(f._sysctl("vm.loadavg"), " ")[1:4]
	case "NetBSD":
		load = strings.Split(f._sysctl("vm.loadavg"), " ")
	case "OpenBSD":
		load = strings.Split(f._sysctl("vm.loadavg"), " ")
	}

	for k, s := range data {
		f.Add(k, f._sysctl(s))
	}

	if len(load) == 3 {
		f.Add("5minLoad", load[0])
		f.Add("10minLoad", load[1])
		f.Add("15minLoad", load[2])
	}

}
