package main

import (
	"example.com/m/v2/d2w"
	"example.com/m/v2/rename"
	"example.com/m/v2/tag"
)

var years = []string{"2019"}

func main() {
	for _, y := range years {
		rename.Rename(y)
		d2w.D2W(y)
		tag.Tag(y)
	}
}
