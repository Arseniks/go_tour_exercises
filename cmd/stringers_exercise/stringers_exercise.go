package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	result := ""
	for i, v := range string(ip[:]) {
		if i != 0 {
			result += "."
		}
		result += fmt.Sprint(v)
	}
	return result
}

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
