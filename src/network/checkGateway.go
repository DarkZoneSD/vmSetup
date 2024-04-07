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
func IsIpInsideNetwork(newIP string, newGateway string) bool {
	subnetBits, err := strconv.Atoi(strings.Split(newIP, "/")[1])
	if err != nil {
		panic(err)
	}
	ipBitArray := returnOctettBits(strings.Split(newIP, "/")[0])
	gatewayBitArray := returnOctettBits(newGateway)

	for i := 0; i < subnetBits; i++ {
		if ipBitArray[i] != gatewayBitArray[i] {
			return false
		}
	}

	return true
}

func returnOctettBits(ip string) [32]int {
	var result [32]int
	binaryStr := ""
	for _, octet := range strings.Split(ip, ".") {
		sTemp, err := strconv.Atoi(octet)
		if err != nil {
			fmt.Println("Error:", err)
			return result // Return the result array as is if there's an error
		}
		// Convert the octet to binary and append to the binary string
		binaryStr += fmt.Sprintf("%08b", sTemp)
	}

	// Ensure the binary string is exactly 32 characters long
	if len(binaryStr) != 32 {
		fmt.Println("Error: IP address does not represent a 32-bit binary number")
		return result // Return the result array as is if the length is incorrect
	}

	// Convert each character in the binary string to an integer and store in the result array
	for i, bit := range binaryStr {
		if bit == '1' {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}
