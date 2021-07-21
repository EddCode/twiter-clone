package main

import (
	"log"

	"github.com/EddCode/twitter-clone/api"
	"github.com/EddCode/twitter-clone/cmd/config"
)

const defaultPort string = "8080"

func main() {
   log.Println("Starting API cmd")
   setting, err := config.GetConfig()

   if err != nil {
       log.Fatal(err.Error())
   }

   port := setting.Server.Port

   if setting.Server.Port == "" {
       port = defaultPort
   }

   api.Start(port)

}
