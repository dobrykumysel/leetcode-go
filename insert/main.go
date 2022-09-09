package main

import (
	"fmt"
	"strings"
)

func main() {

	res := interpret("G()()()()(al)")
	fmt.Println(res)
}
func interpret(command string) string {
	var result string
	s := strings.Split(command, "()")
	for _, val := range s {
		if val == "" {
			result += "o"
		}
		result += val
	}
	var res string
	s = strings.Split(result, "(al)")
	for _, val := range s {
		if val == "" {
			res += "al"
		}
		res += val
	}
	//fmt.Println(s)
	//G -> G
	//() -> o
	//(al) -> al
	return res
}
