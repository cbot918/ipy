package main

import (
	// "github.com/cbot918/liby/cmdy"

	"fmt"
	"os"

	"github.com/cbot918/liby/cmdy"
	u "github.com/cbot918/liby/util"
)


var index int32
var content string

func isAlpha(target byte) bool {
	if target >= 'a' && target <= 'z' || target >= 'A' && target <='Z' {
		return true
	}
	return false
}

func isNumber(target byte) bool {
	if target >= '0' && target <= '9' {
		return true
	}
	return false
}

func next() []byte {
	var result []byte
	for {
		if isAlpha(content[index]) || isNumber(content[index]){
			for isAlpha(content[index]) || isNumber(content[index]) || content[index] == byte('.'){
				result = append(result, content[index])
				index++
			}
			return result
		} else {
			index++
			continue 
		}
	}
	return result
}

func skipTo(target string) {
	for string(next()) != target{}
}

func parse(c []byte,containerName string) []byte {
	var tok string
	for content[index] !='\n' {
		tok = string(next())
		if tok == containerName {
			skipTo("IPv4Address")
			return next()
		}
		// fmt.Println( tok )
	}
	return []byte("")
}

func main(){
	index = 0
	var err error

	args := os.Args
	containerName := args[1]

	c := cmdy.New()
	content,err = c.RunAndGet([]string{"docker network inspect bridge"}); u.Checke(err, "cmdy runandget failed")

	fmt.Println(string(parse([]byte(content),containerName)))

}

