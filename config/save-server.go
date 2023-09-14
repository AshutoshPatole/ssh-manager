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
	for idx, grp := range config.Groups {
		if group == grp.Name {
			existingGroup = idx
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
		isDup := checkDuplicateServer(hostname, config.Groups[existingGroup].Servers)
		// save info
		if !isDup {
			config.Groups[existingGroup].Servers = append(config.Groups[existingGroup].Servers, server)
		} else {
			fmt.Println(color.InYellow("Server already exists in group"))
		}
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
		// if addr is nil then its a domain
		ips, err := net.LookupIP(host)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if len(ips) != 0 {
			return ips[0].String(), nil
		} else {
			return "", nil
		}

	} else {
		// since it is not a domain it has to be a ip
		return host, nil
	}
}

func checkDuplicateServer(host string, servers []Server) bool {
	isDuplicate := false
	ip, _ := IP(host)
	for _, server := range servers {
		if server.HostName == host && ip == server.IP {
			isDuplicate = true
		}
	}
	return isDuplicate
}
