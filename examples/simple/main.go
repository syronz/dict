package main

import (
	"fmt"

	"github.com/syronz/dict"
)

func main() {

	// true means translate in backend, if it is false params not inserted
	dict.Init("./terms.toml", true)

	translated := dict.T("route not found", dict.Ku)
	fmt.Println("route not found: ", translated)

	// translate with params
	phrase := "%v shouldn't be more than %v"
	translated = dict.T(phrase, dict.Ku, dict.R("age"), 18)
	fmt.Println(phrase+": ", translated)

	// if the word not exist around the output add '!!!'
	translated = dict.T("term not exist", dict.Ku)
	fmt.Println("term not exist: ", translated)

	// for prevent adding '!!!' around not existed terms we can use SafeTranslate
	var ok bool
	translated, ok = dict.SafeTranslate("term not exist", dict.Ku)
	fmt.Println("term not exist: ", translated, ok)

}
