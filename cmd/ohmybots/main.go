package main

import (
	"fmt"

	"github.com/Brom95/ohmybot/pkg/config"
)

func main() {
	fmt.Println(config.ReadConfig("config/config.yaml"))
}
