package main

import (
	"fmt"
	"testing"
)
func TestFormat(t *testing.T) {
	got:=Format_ans("hello, hell,  oklarigh")
	fmt.Printf("got %s", got)
}