package main

import (
	"go_tutorials/channels"
	"go_tutorials/collections"
	"go_tutorials/conditional"
	"go_tutorials/generics"
	"go_tutorials/memory"
	"go_tutorials/pointers"
	"go_tutorials/routines"
	"go_tutorials/stringstut"
	"go_tutorials/structs"
)

func main() {
	conditional.ConditionalBlocks()
	collections.Collections()
	memory.MemoryAllocationSpeedTest()
	stringstut.StringsInGo()
	structs.Structs()
	pointers.Pointers()
	routines.Routines()
	channels.Channels()
	generics.Generics()
}
