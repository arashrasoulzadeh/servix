package main

import (
	"fmt"
	config "servix/config"
	db "servix/database"
	"servix/server"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now().String()
	version := 0.1
	fmt.Println("loading config files...")
	fmt.Println("loading database file...")
	fmt.Println("server", version, "started.")
	var c *config.ConfigFile
	c = config.Parse()
	db.Set("started_at", now)
	for _, item := range c.Config {

		wg.Add(1)
		go PrepareToServe(item, &wg)
	}
	wg.Wait()

}

func PrepareToServe(config config.Config, wg *sync.WaitGroup) {
	server.FileServe(":"+strconv.Itoa(config.Port), config.Root, config.Server)
}
