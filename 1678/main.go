package main

import "strings"

var str1 = "()"
var str3 = "(al)"

func main() {

}
func interpret(command string) string {

	strings.Replace(command, "()", "o")
	return ""
}
