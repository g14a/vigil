package main

import (
	"fmt"
	"vigil/client"
)

func main() {

	client.InitESClient()

	//	service.GetDocumentsFromIndex("twitter", "*", 10, 20)
	nodeInfo := client.NodeInfo(client.ESClient)

	fmt.Println(nodeInfo)

	// n, err := service.GetNodesStats()
	// utils.CheckError(err)
	//
	// fmt.Println(n)

	// pages := tview.NewPages()
	// flex := ui.Flex()
	// pages.AddPage("flex", flex, true, true)
	//
	// if err := tview.NewApplication().SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
	// 	panic(err)
	// }
}
