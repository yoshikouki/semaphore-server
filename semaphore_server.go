package main

import "github.com/yoshikouki/semaphore-server/server"

func main()  {
	if err := server.Launch(); err != nil {
		panic(err)
	}
}
