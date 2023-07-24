package logic

import (
	"fmt"
	"io/ioutil"
)

func Default(text string, banner string) string {
	data, err := ioutil.ReadFile("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	fileHash := MD5(string(data))
	if !CheckHash(fileHash, banner) {
		fmt.Println("Error: hash code of " + banner + ".txt has been changed")
		return ""
	}

	return Process(text, data)
}
