package main

import (
	"vigil/client"
	"vigil/service"
)

func main() {
	client.InitESClient()
	service.GetDistinctIndexes()
}
