package main

import (
	"gorm-hex/repository"
	"gorm-hex/service"
	"log"
)

func main() {
	_, err := repository.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	log.Println("Connection Success")

	searchSer := service.NewSearchService()
	landmarks, err := searchSer.FindByName("พระ")
	if err != nil {
		panic(err)
	}
	for _, landmark := range landmarks {
		log.Printf("%v %v", landmark.Name, landmark.Country)
	}

	StartServer()
} 