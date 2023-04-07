package config

type Option struct {
	Config         string `long:"config" short:"c" description:"path to configuration file" required:"true"`
	ManagerAddress string `long:"manager-address" short:"m" description:"log manager address"`
	Version        bool   `long:"version" short:"v" description:"show version"`
}
