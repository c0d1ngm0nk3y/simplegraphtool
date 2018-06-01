package main

import (
	"fmt"
	"model"
)

func main() {
	graph := model.NewGraphData()
	asString := graph.String()
	fmt.Println(asString)
}
