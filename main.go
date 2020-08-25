package main

import (
	"fmt"
	"vigil/client"
	"vigil/service"
	"vigil/utils"
)

func main() {
	client.InitESClient()

	service.GetDocumentsFromIndex("twitter", "*", 10, 20)
	
	n, err := service.GetNodesStats()
	utils.CheckError(err)

	fmt.Println(n)
}
