package friday

import (
	"fmt"

	"git.thwap.org/rockhopper/gout"
)

type Friday struct {
	Facts map[string]string
}

func (f *Friday) _haveFact(k string) bool {
	for mk, _ := range f.Facts {
		if k == mk {
			return true
		}
	}
	return false
}

func (f *Friday) Add(k, v string) {
	f.Facts[k] = v
}

func (f *Friday) Get(k string) string {
	if f._haveFact(k) {
		return f.Facts[k]
	}
	return ""
}

func (f *Friday) Del(k string) {
	if f._haveFact(k) {
		delete(f.Facts, k)
	}
}

func (f *Friday) Print() {
	gout.Setup(false, false, true, "")
	gout.Output.Prompts["info"] = ""
	for k, v := range f.Facts {
		gout.Info(fmt.Sprintf("%s %s %s", k, gout.Bold(gout.Cyan("=>")), v))
	}
}

func (f *Friday) CollectFacts() {
	f._getUnameInfo()
	f._getSysInfo()
	f._getCpuInfo()
}

func New() *Friday {
	retv := Friday{}
	retv.CollectFacts()
	return &retv
}
