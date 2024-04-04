package args

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var configurationBlueprint string = `conf:
   hostname: REPLACE_HOST_NAME
   ipaddress: REPLACE_IP_ADDRESS
   gateway: REPLACE_GATEWAY_ADDRESS
   dns: REPLACE_DNS_ADDRESS`

func HandleArgs(args []string) {
	//Creates temporary file in which the args get saved.
	filepath := "/tmp"
	filename := "vmSetupConf.*.yaml"

	file, err := ioutil.TempFile(filepath, filename)
	if err != nil {
		log.Fatal(err)
	}

	for i := range args {
		switch args[i] {
		case "-i":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_IP_ADDRESS", args[i+1], -1)
			fmt.Println(args[i+1])
		case "--ipaddress":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_IP_ADDRESS", args[i+1], -1)
		case "-g":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_GATEWAY_ADDRESS", args[i+1], -1)
		case "--gateway":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_GATEWAY_ADDRESS", args[i+1], -1)
		case "-d":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_DNS_ADDRESS", args[i+1], -1)
		case "--dns":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_DNS_ADDRESS", args[i+1], -1)
		case "-n":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_HOST_NAME", args[i+1], -1)
		case "--name":
			configurationBlueprint = strings.Replace(configurationBlueprint, "REPLACE_HOST_NAME", args[i+1], -1)
		case "-h":
			DisplayHelpText()
		case "--help":
			DisplayHelpText()
		case "-c":
			interactiveConsole()
		case "--console":
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

	var c conf
	ipAddress, err := c.getIPAddress(file.Name())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("IP Address:", ipAddress)

	//Removes the temporary configuration file
	defer os.Remove(file.Name())
}

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
