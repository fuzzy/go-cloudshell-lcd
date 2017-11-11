package gout

import "fmt"

func Black(i string) string {
	return fmt.Sprintf("\033[30m%s\033[0m", i)
}

func Red(i string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", i)
}

func Green(i string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", i)
}

func Yellow(i string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", i)
}

func Blue(i string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", i)
}

func Purple(i string) string {
	return fmt.Sprintf("\033[35m%s\033[0m", i)
}

func Cyan(i string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", i)
}

func White(i string) string {
	return fmt.Sprintf("\033[37m%s\033[0m", i)
}

func Bold(i string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", i)
}

func Underline(i string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", i)
}

func Blink(i string) string {
	return fmt.Sprintf("\033[5m%s\033[0m", i)
}

func Reverse(i string) string {
	return fmt.Sprintf("\033[7m%s\033[0m", i)
}

func Conceal(i string) string {
	return fmt.Sprintf("\033[8m%s\033[0m", i)
}
