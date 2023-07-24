package main

import (
	"fmt"
	"os"

	"git/ssengerb/ascii-art-justif1y/logic"
	"git/ssengerb/ascii-art-justif1y/pkg"
)

func main() {
	length := len(os.Args)
	if length == 2 {
		if !logic.Characters(os.Args[1]) {
			logic.Error(2)
			return
		}
		fmt.Println(logic.Default(os.Args[1], "standard"))
	} else if length == 3 {
		if len(os.Args[1]) < 8 || os.Args[1][:8] != "--align=" {
			logic.Error(1)
			return
		}

		option := os.Args[1][8:]
		if option != "center" && option != "left" && option != "right" && option != "justify" {
			logic.Error(4)
			return
		}

		if !logic.Characters(os.Args[2]) {
			logic.Error(2)
			return
		}

		s := os.Args[2]
		arr := []string{}
		ans := ""

		for i := 0; i < len(s); i++ {
			if s[i] == 10 {
				arr = append(arr, ans)
				ans = ""
			} else if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
				i++
				arr = append(arr, ans)
				ans = ""
			} else {
				ans += string(s[i])
			}
		}
		if ans != "" {
			arr = append(arr, ans)
		}
		for i := 0; i < len(arr); i++ {
			pkg.Size(option, arr[i], "standard")
		}
	} else if length == 4 {
		if len(os.Args[1]) < 8 || os.Args[1][:8] != "--align=" {
			logic.Error(1)
			return
		}

		option := os.Args[1][8:]
		if option != "center" && option != "left" && option != "right" && option != "justify" {
			logic.Error(3)
			return
		}

		if !logic.Characters(os.Args[2]) {
			logic.Error(2)
			return
		}

		if os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy" {
			logic.Error(3)
			return
		}

		s := os.Args[2]
		arr := []string{}
		ans := ""

		for i := 0; i < len(s); i++ {
			if s[i] == 10 {
				arr = append(arr, ans)
				ans = ""
			} else if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
				i++
				arr = append(arr, ans)
				ans = ""
			} else {
				ans += string(s[i])
			}
		}
		if ans != "" {
			arr = append(arr, ans)
		}
		for i := 0; i < len(arr); i++ {
			pkg.Size(option, arr[i], os.Args[3])
		}

	} else {
		logic.Error(1)
		return
	}
}
