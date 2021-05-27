package main

import "github.com/yoshikouki/semaphore-server/semapi"

func main() {
	err := semapi.Launch(semapi.Config{})
	if err != nil {
		panic(err)
	}
}
