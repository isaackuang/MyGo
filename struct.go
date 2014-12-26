package main

import (
    "encoding/json"
    "fmt"
)

type Graph struct {
    Name  string
    GraphID string
}

type Host struct {
    Name   string
    Graphs []*Graph
}


func main() {
    graph1 := &Graph{"CPU_load", "251"}
    graph2 := &Graph{"Network_flow", "252"}
    Server1 := Host{"Timely", []*Graph{graph1, graph2}}

    js, _ := json.Marshal(Server1)
    fmt.Printf("JSON format: %s\n", js)
}
