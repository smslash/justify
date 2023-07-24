package logic

import (
	"strings"
)

func Process(str string, data []byte) string {
	if len(str) == 0 {
		return ""
	}

	s := ""
	for i := 0; i < len(str); i++ {
		if str[i] == 10 {
			s += "\\" + "n"
		} else {
			s += string(str[i])
		}
	}
	s += " "

	replaced := strings.ReplaceAll(string(data), "\n\n", "\n")
	splited := strings.Split(replaced, "\n")

	ascii := make(map[byte]int)
	var q byte
	for q = 32; q <= 126; q++ {
		ascii[q] = (int(q)-32)*8 + 1
	}

	var newLine []int
	counter := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 92 && s[i+1] == 'n' {
			newLine = append(newLine, i)
			counter++
			i++
		}
	}
	var result []string

	if counter == 0 {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(s)-1; j++ {
				result = append(result, splited[ascii[s[j]]+i])
			}
			result = append(result, "\n")
		}
	} else {
		for i := 0; i < 8; i++ {
			for j := 0; j < len(s)-1; j++ {
				if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
					break
				}
				result = append(result, splited[ascii[s[j]]+i])
			}
			result = append(result, "\n")
		}
		for k := 0; k < len(newLine); k++ {
			for i := 0; i < 8; i++ {
				for j := newLine[k] + 2; j < len(s)-1; j++ {
					if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
						break
					}
					result = append(result, splited[ascii[s[j]]+i])
				}
				result = append(result, "\n")
			}
		}
	}
	res := ""
	for i := 0; i < len(result); i++ {
		res += result[i]
	}

	text := strings.ReplaceAll(res, "\n\n\n\n\n\n\n\n", "\n")

	new_line_str := ""
	if counter == 0 {
		return text
	} else if counter == (len(s)-1)/2 {
		for i := 0; i < counter; i++ {
			new_line_str += "\n"
		}
		return new_line_str
	}
	return text
}
