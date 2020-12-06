package main

import (
	"fmt"
	config "servix/config"
	"servix/server"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	version := 0.1
	fmt.Println("loading config files...")
	fmt.Println("server", version, "started.")
	var c *config.ConfigFile
	c = config.Parse()
	for _, item := range c.Config {
		wg.Add(1)
		go PrepareToServe(item, &wg)
	}
	wg.Wait()

}

func PrepareToServe(config config.Config, wg *sync.WaitGroup) {
	server.FileServe(":"+strconv.Itoa(config.Port), config.Root)
}
