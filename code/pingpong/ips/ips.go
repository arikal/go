package ips

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type IpAddress string
type Ips []IpAddress

func LoadFromFile(filename string) Ips {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(bs), "\n")
	var ips Ips

	for _, line := range lines {
		ips = append(ips, IpAddress(line))
	}

	return ips
}

func (i Ips) GetRandomIp() IpAddress {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(i))

	return i[n]
}
