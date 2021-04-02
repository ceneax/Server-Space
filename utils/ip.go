package utils

import (
	"net"
)

func IptoInt(ip string) int64 {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0
	}
	return int64(b[3]) | int64(b[2])<<8 | int64(b[1])<<16 | int64(b[0])<<24
}
