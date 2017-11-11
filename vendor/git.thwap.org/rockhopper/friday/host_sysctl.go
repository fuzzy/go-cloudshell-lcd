// +build freebsd openbsd netbsd

package friday

import "runtime"

func (f *Friday) _getUnameInfo() {
	data := map[string]string{
		"Sysname":  "kern.ostype",
		"Release":  "kern.osrelease",
		"Machine":  "hw.machine",
		"Nodename": "kern.hostname",
		"Version":  "kern.version",
	}
	for k, s := range data {
		f.Add(k, f._sysctl(s))
	}
	f.Add("Sysname_lc", runtime.GOOS)
}
