package utils

type Server struct {
	HostName string `mapstructure:"hostname"`
	IP       string `mapstructure:"ip"`
	KeyAuth  bool   `mapstructure:"keyAuth"`
}

type Env struct {
	Name    string   `mapstructure:"name"`
	Servers []Server `mapstructure:"servers"`
}

type Group struct {
	Name        string `mapstructure:"name"`
	User        string `mapstructure:"user"`
	Environment []Env  `mapstructure:"env"`
}

type Config struct {
	Groups []Group `mapstructure:"groups"`
}
