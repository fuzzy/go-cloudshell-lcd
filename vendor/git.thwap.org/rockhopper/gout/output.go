package gout

import (
	"fmt"
	"log"
	"time"

	eval "github.com/Knetic/govaluate"
)

// Output functions
func consoleOutput(t string, e string, f string, args ...interface{}) {
	Winsize = ConsInfo()
	if Output.ToFile {
		log.Printf("%s %s\n", t, fmt.Sprintf(f, args...))
	}
	fmt.Printf("%s %s %s%s",
		Output.Prompts[t],
		fmt.Sprintf(f, args...),
		padding(int(Winsize.Col)-(7+len(fmt.Sprintf(f, args...)))),
		e)
}

func Info(f string, args ...interface{}) {
	if !Output.Quiet {
		consoleOutput("info", "\n", f, args...)
	}
}

func Debug(f string, args ...interface{}) {
	if Output.Debug {
		consoleOutput("debug", "\n", f, args...)
	}
}

func Warn(f string, args ...interface{}) {
	if Output.Verbose {
		consoleOutput("warn", "\n", f, args...)
	}
}

func Error(f string, args ...interface{}) {
	consoleOutput("error", "\n", f, args...)
}

func Status(f string, args ...interface{}) {
	consoleOutput("status", "\r", f, args...)
}

func ProgressStatus(d *Progress, l int) string {
	// compute runtime
	runtime := (time.Now().Unix() - d.TimeStarted)
	// compute speed
	if runtime <= 0 {
		runtime = 1
	}
	speed := (d.CurrentOut / runtime)
	// compute estimated time remaining
	if speed <= 0 {
		speed = 0
	}
	etr := ((d.Total - d.CurrentOut) / speed)
	// compute progress percentage
	progress := ((float64(d.CurrentOut) / float64(d.Total)) * 100.0)

	// now build our return string
	current_s := HumanSize(d.CurrentOut)
	speed_s := HumanSize(speed)
	etr_s := HumanTimeColon(etr)
	var progress_s string

	// get the current string length
	first := fmt.Sprintf("%9s @ %9s %6.02f%s", current_s, speed_s, progress, "%%")

	if (len(first) + len(etr_s)) <= (l - 16) {
		progress_s = ProgressMeter((l - (len(first) + len(etr_s) + 2)), int(progress))
	} else {
		progress_s = ""
	}

	return fmt.Sprintf("%s %s %s", first, progress_s, etr_s)
}

func ProgressMeter(l int, p int) string {
	rl := l - 2
	var rslt interface{}
	var sp string
	if p < 100 {
		expr, _ := eval.NewEvaluableExpression(fmt.Sprintf("%d * 0.%02d", rl, p))
		rslt, _ = expr.Evaluate(nil)
		sp = repeat("#", int(rslt.(float64)))
	} else {
		rslt = rl
		sp = repeat("#", rslt.(int))
	}
	pd := padding(rl - len(sp))
	return fmt.Sprintf("[%s%s]", sp, pd)
}

func Throbber() string {
	if Output.lastThrob == len(Output.Throbber)+1 {
		Output.lastThrob = 1
	} else {
		Output.lastThrob++
	}
	return Output.Throbber[Output.lastThrob-1]
}
