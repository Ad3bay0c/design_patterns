package main

import (
	"fmt"
	"strings"
)
const (
	indentSize = 2
)
type HTMLElement struct {
	name string
	text string
	elements []HTMLElement
}

func geeralExample() {
	sb := strings.Builder{}
	sb.WriteString("Hello")
	st := []strings.Builder{}
	st = append(st, sb)
	fmt.Println(st, sb.Len(), sb.Cap())
}

func main() {
	geeralExample()
}
