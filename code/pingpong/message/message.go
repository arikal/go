package message

import "arikal.com/go/pingpong/ips"

type message struct {
	from ips.IpAddress
	to   ips.IpAddress
	body string
}
