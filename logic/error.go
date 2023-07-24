package logic

import "fmt"

func Error(a int) {
	if a == 1 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
		fmt.Println("Example: go run . --align=right something standard")
		return
	} else if a == 2 {
		fmt.Println("Error: Input only ASCII Table characters between 32-126 numbers")
		return
	} else if a == 3 {
		fmt.Println("Error: Only standard, shadow, and thinkertoy banners are available")
		return
	} else if a == 4 {
		fmt.Println("Error: Only left, right, center, and justify flags are available")
		return
	}
}
