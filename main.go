package main

import (
	"fmt"
	src "github.com/anupamdas24/text-to-markdown/src"
	"io/ioutil"
	"os"
)

func main(){
	contentInBytes, _ := ioutil.ReadAll(os.Stdin)
	str := string(contentInBytes)
	str = src.ParseTextToMarkdown(str)
	fmt.Printf("%s\n",str)
}
