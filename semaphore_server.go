package main

import "github.com/yoshikouki/semapi/semapi"

func main() {
	err := semapi.Launch(semapi.Config{})
	if err != nil {
		panic(err)
	}
}
