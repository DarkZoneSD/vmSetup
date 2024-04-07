package args

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/DarkZoneSD/vmSetup/src/network"
)

var configurationBlueprint string = `hostname: REPLACE_HOST_NAME
ipaddress: REPLACE_IP_ADDRESS
gateway: REPLACE_GATEWAY_ADDRESS
dns: REPLACE_DNS_ADDRESS`

var flagToPlaceholder = map[string]string{
	"-i":          "REPLACE_IP_ADDRESS",
	"--ipaddress": "REPLACE_IP_ADDRESS",
	"-g":          "REPLACE_GATEWAY_ADDRESS",
	"--gateway":   "REPLACE_GATEWAY_ADDRESS",
	"-d":          "REPLACE_DNS_ADDRESS",
	"--dns":       "REPLACE_DNS_ADDRESS",
	"-n":          "REPLACE_HOST_NAME",
	"--name":      "REPLACE_HOST_NAME",
}

func HandleArgs(args []string) {
	//Creates temporary file in which the args get saved.
	filepath := "/tmp"
	filename := "vmSetupConf.*.yaml"

	file, err := ioutil.TempFile(filepath, filename)
	if err != nil {
		log.Fatal(err)
	}

	for i := range args {
		if replacement, ok := flagToPlaceholder[args[i]]; ok {
			if i+1 < len(args) {
				configurationBlueprint = strings.Replace(configurationBlueprint, replacement, args[i+1], -1)
			} else {
				fmt.Println("Error: Missing value for", args[i])
				return
			}
		} else if args[i] == "-h" || args[i] == "--help" {
			DisplayHelpText()
		} else if args[i] == "-c" || args[i] == "--console" {
			interactiveConsole()
		}
	}

	fmt.Println("Saved new configuration to file:", file.Name())
	// fmt.Println(configurationBlueprint)
	err = ioutil.WriteFile(file.Name(), []byte(configurationBlueprint), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	c := &conf{}
	ipAddress, err := c.getIPAddress(file.Name())
	if err != nil {
		fmt.Println("Error reading IP address:", err)
		return
	}
	gateway, err := c.getGateway(file.Name())
	if err != nil {
		fmt.Println("Error reading Gateway:", err)
		return
	}

	//Removes the temporary configuration file
	defer os.Remove(file.Name())
	network.IsGatewayInsideTheNetwork(ipAddress, gateway)
}

// Displays the possible arguments when calling this program
//
//  -n NewHostName         New Hostname of the Machine
//  -i IPAddress           New IPAddress of the Machine
//  -g Gateway             Gateway of the new Network
//  -d DNS                 Nameservers of the new Network
//  -c Console 		Starts an interactive console
func DisplayHelpText() {
	fmt.Println(`
	-n NewHostName         New Hostname of the Machine
	-i IPAddress           New IPAddress of the Machine
	-g Gateway             Gateway of the new Network
	-d DNS                 Nameservers of the new Network, use comma to set multiple nameservers
	-c Console			   Starts an interactive console

	For Example:
	./vmSetup -n NewHost -i 192.168.10.12/24 -g 192.168.10.254 -d 192.168.10.254
	./vmSetup --name HostName --ipaddress 10.0.1.10/25 --gateway 10.0.0.1 --dns 10.0.0.2,10.0.0.3
	`)
	os.Exit(0)
}
