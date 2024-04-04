package args

import (
	"fmt"

	"github.com/3th1nk/cidr"
)

func checkGateway(newIP string, newGateway string) {
	ip, err := cidr.Parse(newGateway)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ip.Network().String())
}
