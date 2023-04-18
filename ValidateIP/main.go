package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dataset := []string{
		"172.16.254.1",                      // IPv4
		"2001:0db8:85a3:0:0:8A2E:0370:7334", // IPv6
		"256.256.256.256",                   // neither
	}

	for _, data := range dataset {
		fmt.Println(validIPAddress(data), ":", data)
	}
}

func validIPAddress(queryIP string) string {
	// check if it is a valid IPv4 address
	parts := strings.Split(queryIP, ".")
	if len(parts) == 4 {
		for _, part := range parts {
			// check if the part is a valid number between 0 and 255
			num, err := strconv.Atoi(part)
			if err != nil || num < 0 || num > 255 {
				return "Neither"
			}
			// check if the part has leading zeroes
			if len(part) > 1 && part[0] == '0' {
				return "Neither"
			}
		}
		return "IPv4"
	}

	// check if it is a valid IPv6 address
	parts = strings.Split(queryIP, ":")
	if len(parts) == 8 {
		for _, part := range parts {
			// check if the part is a valid hexadecimal string
			if len(part) < 1 || len(part) > 4 {
				return "Neither"
			}
			for _, c := range part {
				if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}

	// if it is not a valid IPv4 or IPv6 address
	return "Neither"
}
