/*
   Copyright (c) 2014, Mike 'Fuzzy' Partin <fuzzy@fumanchu.org>
   All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:

  1. Redistributions of source code must retain the above copyright notice, this
     list of conditions and the following disclaimer.

  2. Redistributions in binary form must reproduce the above copyright notice,
     this list of conditions and the following disclaimer in the documentation
     and/or other materials provided with the distribution.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package gout

import (
	"log"
	"os"
)

var (
	Output  output
	Winsize winsize
	Logfile *os.File
)

// Output setup function
func Setup(d bool, q bool, v bool, f string) {
	var toFile bool
	if len(f) > 0 {
		Logfile, e := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if e != nil {
			panic(e)
		}
		log.SetOutput(Logfile)
		toFile = true
	} else {
		toFile = false
	}
	Output = output{
		Prompts:   make(map[string]string),
		Debug:     d,
		Quiet:     q,
		Verbose:   v,
		ToFile:    toFile,
		Throbber:  []string{},
		lastThrob: 0,
	}
	Output.Prompts["info"] = "INFO"
	Output.Prompts["warn"] = "WARN"
	Output.Prompts["debug"] = "DEBUG"
	Output.Prompts["error"] = "ERROR"
	Output.Prompts["status"] = ""
	Output.Throbber = []string{"-", "\\", "|", "/"}
}

// Setup example
/*
func init() {
	Winsize = consInfo()
	Setup(true, false, true, false)
	Output.Prompts["info"] = fmt.Sprintf("%s%s%s",
		String(".").Cyan(),
		String(".").Bold().Cyan(),
		String(".").Bold().White())
	Output.Prompts["warn"] = fmt.Sprintf("%s%s%s",
		String(".").Yellow(),
		String(".").Bold().Yellow(),
		String(".").Bold().White())
	Output.Prompts["debug"] = fmt.Sprintf("%s%s%s",
		String(".").Purple(),
		String(".").Bold().Purple(),
		String(".").Bold().White())
	Output.Prompts["error"] = fmt.Sprintf("%s%s%s",
		String(".").Red(),
		String(".").Bold().Red(),
		String(".").Bold().White())
	Output.Prompts["status"] = fmt.Sprintf("%s%s%s",
		String("-").Cyan(),
		String("-").Bold().Cyan(),
		String("-").Bold().White())
}
*/
