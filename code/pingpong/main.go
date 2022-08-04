package main

import (
	"fmt"

	"arikal.com/go/pingpong/ips"
)

func main() {
	ips := ips.LoadFromFile("ips.txt")

	fmt.Println(ips)
	fmt.Println(ips.GetRandomIp())
}
