package network

import (
	"fmt"
	"strconv"
	"strings"

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
func IsGatewayInsideTheNetwork(newIP string, newGateway string) bool {
	subnetBits, err := strconv.Atoi(strings.Split(newIP, "/")[1])
	if err != nil {
		panic(err)
	}
	ipBitArray := returnOctettBits(strings.Split(newIP, "/")[0])
	gatewayBitArray := returnOctettBits(newGateway)

	for i := 0; i < subnetBits; i++ {

	}

	return true
}

func returnOctettBits(ip string) [4]string {
	var result [4]string
	for i := 0; i < len(result); i++ {
		sTemp, err := strconv.Atoi(strings.Split(ip, ".")[i])
		if err != nil {
			fmt.Println("Error:", err)
		}
		binaryStr := fmt.Sprintf("%08b", sTemp)
		result[i] = binaryStr
	}
	return result
}

func convCidrToBitArray(cidr int) []string {
	var sTempArray []string
	var tempString string

	for i := 0; i < cidr/8; i++ {
		sTempArray = append(sTempArray, "11111111")
	}
	for i := 0; i < cidr%8; i++ {
		tempString += "1"
	}

	sTempArray = append(sTempArray, tempString)

	return sTempArray
}
