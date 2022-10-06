package main

import "fmt"
import "salutgo/util"
import . "salutgo/practise"
import _ "salutgo/oauth2.0"

func main() {
	fmt.Println("hello")
	a := util.MinInt(12, 11)
	fmt.Printf("a: %v\n", a)
	jake := Person{Name: "name"}
	fmt.Printf("jake: %v\n", jake)
}