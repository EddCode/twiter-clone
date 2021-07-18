package main

import (
	"log"

	"github.com/EddCode/twitter-clone/api"
	"github.com/EddCode/twitter-clone/cmd/config"
)

const defaultPort string = "8080"

func main() {
   log.Println("Starting API cmd")
   setting, _ := config.GetConfig("./config/config.yml")

   port := setting.Server.Port

   if setting.Server.Port == "" {
       port = defaultPort
   }

   api.Start(port)

}
