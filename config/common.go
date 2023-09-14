package config

type Server struct {
	HostName string `mapstructure:"hostname"`
	IP       string `mapstructure:"ip"`
	KeyAuth  bool   `mapstructure:"keyAuth"`
}

type Group struct {
	Name    string   `mapstructure:"name"`
	User    string   `mapstructure:"user"`
	Servers []Server `mapstructure:"servers"`
}

type Config struct {
	Groups []Group `mapstructure:"groups"`
}
