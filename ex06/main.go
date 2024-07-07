package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	var a string
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	fmt.Printf("%#v\n", piscine.Split(a, "HA"))
	fmt.Printf("%#v\n", piscine.Split("", "HA"))
	fmt.Printf("%#v\n", piscine.Split("H", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HA", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HelloHAhowHAareHAyou?HA", "HA"))
	fmt.Printf("%#v\n", piscine.Split("HAHelloHAhowHAareHAyou?HA", "HA"))
}
