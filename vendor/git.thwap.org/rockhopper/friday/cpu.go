package friday

import (
	"fmt"
	"runtime"
)

func (f *Friday) _getCpuInfo() {
	f.Add("ProcessorCount", fmt.Sprintf("%d", runtime.NumCPU()))
	f.Add("Architecture", runtime.GOARCH)
}
