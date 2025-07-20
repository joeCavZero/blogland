package logger

import (
	"fmt"
	"os"
)

const (
	red    = "\033[1;31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"

	//prefix_color = "\033[1;33m"
)

func Infof(format string, args ...any) {
	fmt.Printf(
		"%s[INFO] %s%s\n",
		green,
		reset,
		fmt.Sprintf(format, args...),
	)
}

func ExitErrorf(format string, args ...any) {
	fmt.Printf(
		"%s[ERROR] %s%s\n",
		red,
		reset,
		fmt.Sprintf(format, args...),
	)
	os.Exit(1)
}

func WebInfof(format string, args ...any) {
	fmt.Printf(
		"%s[WEB] %s[INFO] %s%s\n",
		yellow,
		green,
		reset,
		fmt.Sprintf(format, args...),
	)
}

func WebErrorf(format string, args ...any) {
	fmt.Printf(
		"%s[WEB] %s[ERROR] %s%s\n",
		yellow,
		red,
		reset,
		fmt.Sprintf(format, args...),
	)
}

func WebExitErrorf(format string, args ...any) {
	fmt.Printf(
		"%s[WEB]%s[ERROR] %s%s\n",
		yellow,
		red,
		reset,
		fmt.Sprintf(format, args...),
	)
	os.Exit(1)
}

func APIInfof(format string, args ...any) {
	fmt.Printf(
		"%s[API] %s[INFO] %s%s\n",
		cyan,
		green,
		reset,
		fmt.Sprintf(format, args...),
	)
}

func APIErrorf(format string, args ...any) {
	fmt.Printf(
		"%s[API] %s[ERROR] %s%s\n",
		cyan,
		red,
		reset,
		fmt.Sprintf(format, args...),
	)
}

func APIExitErrorf(format string, args ...any) {
	fmt.Printf(
		"%s[API] %s[ERROR] %s%s\n",
		cyan,
		red,
		reset,
		fmt.Sprintf(format, args...),
	)
	os.Exit(1)
}
