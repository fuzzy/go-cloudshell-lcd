package gout

import (
	"fmt"
	"io"
	"time"
)

func repeat(c string, n int) string {
	retv := ""
	for i := 0; i <= n; i++ {
		retv += "#"
	}
	return retv
}

func padding(spaces int) string {
	pad := " "
	retv := ""
	for i := 0; i < spaces; i++ {
		retv = fmt.Sprintf("%s%s", retv, pad)
	}
	return retv
}

func Copy(dst io.Writer, src io.Reader, data *Progress) (int64, error) {
	var err error
	buf := make([]byte, 64*1024)

	for {
		// we move the ConsoleInfo() call inside the loop, that way terminal size
		// changes get handled properly.
		cons := ConsInfo()
		nr, er := src.Read(buf)
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
		data.CurrentIn += int64(nr)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
			data.CurrentOut += int64(nw)
		}
		if (time.Now().Unix() - data.TimeStarted) >= 1 {
			Status("%s", ProgressStatus(data, int(cons.Col-5)))
		}
	}
	return data.CurrentOut, err
}
