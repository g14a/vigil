package ui

import (
	"github.com/rivo/tview"
)

func Flex() *tview.Flex {

	list := tview.NewList()

	list.ShowSecondaryText(false).
		AddItem("Home", "", 'h', func() {}).
		AddItem("Nodes", "", 'n', func() {}).
		AddItem("Indices", "", 'i', func() {}).
		AddItem("Rest", "", 'r', func() {})

	NodeInfoBox := tview.NewTable().SetBorder(true).SetTitle("Node Information")

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().AddItem(list, 20, 1, true), 20, 1, true).
		AddItem(NodeInfoBox, 0, 1, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Cluster Information"), 0, 1, false)

	return flex
}
