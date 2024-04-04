package args

import (
	"github.com/3th1nk/cidr"
)

func checkIpAddress(newIP string) string {
	ip, err := cidr.Parse(newIP)
	if err != nil {
		return err.Error()
	}
	//var netadddress = ip.Network().String()
	return ip.Network().String()
}
