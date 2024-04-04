package args

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	HostName  string `yaml:"hostname"`
	IPAddress string `yaml:"ipaddress"`
	Gateway   string `yaml:"gateway"`
	Dns       string `yaml:"dns"`
}

func (c *conf) getIPAddress(filepath string) (string, error) {
	fmt.Println("Reading from file: ", filepath)
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("yamlFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return "", fmt.Errorf("Unmarshal: %v", err)
	}
	return c.IPAddress, nil
}
