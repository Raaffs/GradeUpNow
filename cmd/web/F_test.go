package main

import (
	"fmt"
	"testing"
)
func TestFormat(t *testing.T) {
	got:=Format_ans("hello, ,+-*/  ok alrigh")
	fmt.Printf("got %s", got)
}
