package main

import (
	KrunkerAPI "krunker-api"
	"log"
)

func main() {
	api, _ := KrunkerAPI.NewKrunkerAPI()
	profile, _ := api.GetProfile("a6a6")

	log.Println(*profile)
}
