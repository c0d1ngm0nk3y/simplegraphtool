package main

import (
	"fmt"
	"io/ioutil"
	"model"
	"strings"
)

func main() {
	graph := model.NewGraphData()

	data, err := ioutil.ReadFile("edges.csv")
	if err == nil {
		dataAsString := string(data)
		lines := strings.Split(dataAsString, "\n")
		for _, line := range lines {
			success, edge := model.ParseEdge(line)
			if success {
				graph.Add(edge)
			}
		}

	}
	asString := graph.String()
	fmt.Println(asString)
}
