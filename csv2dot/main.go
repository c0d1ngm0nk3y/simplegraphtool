package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/c0d1ngm0nk3y/simplegraphtool/model"
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
