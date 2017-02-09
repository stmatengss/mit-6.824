package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) 
	f := func(c rune) bool {
		if c == ' ' {
			return true
		} else {
			return false
		}
	}
	for _, i := range strings.FieldsFunc(s, f)  {
		m[i] ++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
