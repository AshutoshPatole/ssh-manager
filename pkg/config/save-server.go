package config

import (
	"fmt"
	"log"
	"net"

	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
)

func SaveServer(hostname, user, group string, keyAuth bool) {
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	existingGroup := -1
	for _, grp := range config.Groups {
		if group == grp.Name {
			existingGroup = 1
		}
	}
	ip, err := IP(hostname)
	if err != nil {
		fmt.Println(color.InYellow("Could not resolve IP address"))
	}

	server := Server{
		HostName: hostname,
		IP:       ip,
		KeyAuth:  keyAuth,
	}

	if existingGroup == -1 {
		// create a group and save info
		newGroup := Group{
			Name:    group,
			User:    user,
			Servers: []Server{server},
		}
		config.Groups = append(config.Groups, newGroup)
	} else {
		// save info
		config.Groups[existingGroup].Servers = append(config.Groups[existingGroup].Servers, server)
	}

	// save the information in config file
	viper.Set("groups", config.Groups)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalln(err)
	}

}

func IP(host string) (string, error) {
	addr := net.ParseIP(host)
	if addr == nil {
		// fmt.Println("Given String is a Domain Name")
		ips, err := net.LookupIP(host)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Print the IP addresses
		for _, ip := range ips {
			fmt.Println(ip)
		}

		return ips[0].String(), nil

	} else {
		// fmt.Println("Given String is a Ip Address")
		return host, nil
	}

}
