package network

import (
	"fmt"

	"github.com/3th1nk/cidr"
)

//TODO
// func (netaddress, gateway)
//return true if gateway is in the network
// false if it is outside of the network
func CheckGateway(newIP string, newGateway string) {
	ip, err := cidr.Parse(newIP)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ip.Network().String())
}
func isGatewayInsideTheNetwork(newIP string, newGateway string) bool {
	return true
}