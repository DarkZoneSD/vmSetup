package network

import (
	"fmt"
	"strconv"
	"strings"
)

//	Determines if an IP address is within the same network as a specified CIDR IP.
//		Returns true if within the same network, false otherwise.
func IsIpInsideNetwork(newCidr string, newIp string) bool {
	subnetBits, err := strconv.Atoi(strings.Split(newCidr, "/")[1])
	if err != nil {
		panic(err)
	}
	ipBitArray := returnBits(strings.Split(newCidr, "/")[0])
	gatewayBitArray := returnBits(newIp)

	for i := 0; i < subnetBits; i++ {
		if ipBitArray[i] != gatewayBitArray[i] {
			return false
		}
	}

	return true
}

func returnBits(ip string) [32]int {
	var result [32]int
	binaryStr := ""
	for _, octet := range strings.Split(ip, ".") {
		sTemp, err := strconv.Atoi(octet)
		if err != nil {
			fmt.Println("Error:", err)
			return result
		}
		binaryStr += fmt.Sprintf("%08b", sTemp)
	}

	if len(binaryStr) != 32 {
		fmt.Println("Error: IP address does not represent a 32-bit binary number")
		return result
	}

	for i, bit := range binaryStr {
		if bit == '1' {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}
