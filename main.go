package main

import (
	"flag"
	"fmt"
	"github.com/player-server/conf"
	"github.com/player-server/router"
)

func main() {
	env:= flag.String("e", "dev", "Your environment")
	flag.Parse()
	fmt.Printf("Using environment variable %s.yml\n", env)
	conf.Init(*env)
	router.Init()
}
