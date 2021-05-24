package main

import "github.com/yoshikouki/semaphore-server/server"

func main() {
	err := server.Launch(server.Config{})
	if err != nil {
		panic(err)
	}
}
