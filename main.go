package main

import (
	"fmt"
	"vigil/client"
	"vigil/service"
	"vigil/utils"
)

func main() {
	client.InitESClient()
	// service.GetIndiceInfo()
	//
	// service.GetDocumentsFromIndex("heroku_2r69v17f.candidates")
	n, err := service.GetNodesStats()
	utils.CheckError(err)
	
	fmt.Println(n)
}
