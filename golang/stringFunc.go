package main

import (
	"fmt"
	"strings"
)

// Contains
// 判断字串是否存在于父串中
func call_Contains() {
	strings.Contains("hahahaha", "ha")
	// strings.Contains("hahahaha", 'h')
	strings.Contains("hahahaha", "")
	strings.Contains("", "")
}

func main() {
	call_Contains()
}