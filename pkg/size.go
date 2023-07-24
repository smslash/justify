package pkg

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"git/ssengerb/ascii-art-justif1y/logic"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func Size(option string, text string, banner string) {
	if text == "" {
		return
	}
	width1, err := getTerminalWidth()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	width := int(width1)
	textTerminal := logic.Default(text, banner)
	textSize := len(textTerminal)/8 - 1

	if width < textSize {
		fmt.Println("Error: Too much characters")
		return
	}

	size := width - textSize
	words := strings.Fields(text)

	if len(words) == 1 && (option == "left" || option == "justify") {
		fmt.Print(textTerminal)
	} else if option == "left" {
		fmt.Print(textTerminal)
	} else if option == "right" {
		space := 0
		for i := 0; i < len(textTerminal); i++ {
			if i == 0 {
				fmt.Print(" ")
				for j := 0; j < size; j++ {
					fmt.Print(" ")
				}
			} else if textTerminal[i] == '\n' && space < 7 {
				fmt.Println()
				space++
				for j := 0; j < size; j++ {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(string(textTerminal[i]))
			}
		}
	} else if option == "center" {
		if size%2 == 1 {
			size++
		}
		space := 0
		for i := 0; i < len(textTerminal); i++ {
			if i == 0 {
				fmt.Print(" ")
				for j := 0; j < size/2; j++ {
					fmt.Print(" ")
				}
			} else if textTerminal[i] == '\n' && space < 7 {
				space++
				fmt.Println()
				for j := 0; j < size/2; j++ {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(string(textTerminal[i]))
			}
		}
	} else {
		arr := []string{}
		for _, v := range words {
			arr = append(arr, logic.Default(v, banner))
		}
		size /= (len(arr) - 1)
		size += (2 * len(arr))

		for k := 0; k < 8; k++ {
			for i := 0; i < len(arr)-1; i++ {
				for j := (len(arr[i]) / 8) * k; j < len(arr[i]); j++ {
					if arr[i][j] == '\n' {
						break
					} else {
						fmt.Print(string(arr[i][j]))
					}
				}
				for j := 0; j < size; j++ {
					fmt.Print(" ")
				}
			}
			for i := len(arr) - 1; i < len(arr); i++ {
				for j := (len(arr[i]) / 8) * k; j < len(arr[i]); j++ {
					if arr[i][j] == '\n' {
						break
					} else {
						fmt.Print(string(arr[i][j]))
					}
				}
			}
			fmt.Println()
		}
	}
	fmt.Println()
}

func getTerminalWidth() (uint16, error) {
	ws := &winsize{}

	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		return 0, errno
	}

	return ws.Col, nil
}
