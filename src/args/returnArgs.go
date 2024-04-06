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

func (c *conf) getGateway(filepath string) (string, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("yamlFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return "", fmt.Errorf("Unmarshal: %v", err)
	}
	return c.Gateway, nil
}

func (c *conf) getDns(filepath string) (string, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("yamlFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return "", fmt.Errorf("Unmarshal: %v", err)
	}
	return c.Dns, nil
}
func (c *conf) getHostname(filepath string) (string, error) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("yamlFile.Get err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return "", fmt.Errorf("Unmarshal: %v", err)
	}
	return c.HostName, nil
}
