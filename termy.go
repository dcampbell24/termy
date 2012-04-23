package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const ESC = 0x1B

// Attributes
const (
	RESET = 0
	BOLD = 1
	UNDERLINE = 4
	BLINK = 5
	REVERSE = 7
	HIDDEN  = 8
)

// Colors
const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

func TextColor(attr, fg, bg int) string {
	if fg == -1 && bg == -1 {
		return fmt.Sprintf("%c[%dm", ESC, attr)
	} else if fg == -1 {
		return fmt.Sprintf("%c[%d;%dm", ESC, attr, bg + 40)
	} else if bg == -1 {
		return fmt.Sprintf("%c[%d;%dm", ESC, attr, fg + 30)
	}
	return fmt.Sprintf("%c[%d;%d;%dm", ESC, attr, fg + 30, bg + 40)
}

// GetSize returns the dimensions of the window.
func GetSize() (width, height int, err error) {
	var dimensions [4]uint16

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL,
			uintptr(syscall.Stdout),
			uintptr(syscall.TIOCGWINSZ),
			uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0);
			err != 0 {
		return -1, -1, err
	}
	return int(dimensions[1]), int(dimensions[0]), nil
}

var (
	attr = []string{"RESET", "BOLD", "UNDERLINE", "BLINK", "REVERSE", "HIDDEN"}
	clr = []string{"BLACK", "RED", "GREEN", "YELLOW", "BLUE", "MAGENTA", "CYAN", "WHITE"}
)

func main() {
	for i, a := range []int{0, 1, 4, 5, 7, 8} {
		fmt.Printf("%s%s%s\n", TextColor(a, -1, -1), attr[i], TextColor(RESET, -1, -1))
	}
	for i := range clr {
		fmt.Printf("%s%-7s ",  TextColor(RESET, i, -1), clr[i])
		for j := range clr {
			fmt.Printf("%s%s ", TextColor(RESET, j, i), clr[j])
		}
		fmt.Println(TextColor(RESET, -1, -1))
	}
	width, height, err := GetSize()
	if err != nil {
		fmt.Println("termy: failed to get window dimensions:", err)
	}
	fmt.Printf("Width: %d\nHeight: %d\n", width, height)
}
