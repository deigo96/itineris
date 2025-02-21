package main

import "github.com/deigo96/bpkp/app/config"

func main() {
	configuration := config.GetConfig()
	_ = config.DBConnection(configuration)
}
