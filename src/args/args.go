package args

import (
	"fmt"
	"os"
)

func HandleArgs(args []string) {
	for i := range args {
		switch args[i] {
		case "-s":
			fmt.Println("S - arg found")
		case "--subnet":
			fmt.Println("Subnet - arg found")
		case "-i":
			localAddresses()
		case "--ipaddress":
			localAddresses()
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
}

func DisplayHelpText() {
	fmt.Println(`
	-n NewHostName         New Hostname of the Machine
	-i IPAddress           New IPAddress of the Machine
	-s SubnetMask          Subnetmask of the new Network
	-g Gateway             Gateway of the new Network
	-d DNS                 Nameservers of the new Network, use comma to set multiple nameservers

	-c
	--console			   Starts an interactive console
	For Example:
	./vmSetup -n NewHost -i 192.168.10.12 -s 255.255.255.0 -g 192.168.10.254 -d 192.168.10.254
	./vmSetup --name HostName --ipaddress 10.0.1.10 --ssubnet 255.255.128.0 --gateway 10.0.0.1 --dns 10.0.0.2,10.0.0.3
	`)
	os.Exit(0)
}
