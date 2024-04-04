package args

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func interactiveConsole() {
	printConsoleHelpText()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("interfaces", text) == 0 {
			localAddresses()
		} else if strings.Compare("quit", text) == 0 || strings.Compare("q", text) == 0 {
			return
		} else if strings.Compare("help", text) == 0 || strings.Compare("h", text) == 0 {
			printConsoleHelpText()
		} else if strings.Compare("clear", text) == 0 || strings.Compare("c", text) == 0 {
			CallClear()
		} else {
			fmt.Println("No command found for the input.")
		}

	}
}
func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v", err.Error()))
		}

		for _, a := range addrs {
			fmt.Printf("%v - %v\n", i.Name, a)
		}
	}
}

func printConsoleHelpText() {
	fmt.Println(`
All commands can be shortened to the first letter
-------------------------------------------------
interfaces         New Hostname of the Machine
clear 		    Clears the console window
help 		   Displays this help text
quit 		   Quits the interactive console`)
}
