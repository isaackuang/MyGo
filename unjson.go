
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

var hosts Host
myServer := []byte(`{ {"Name":"Timely","Graphs":[{"Name":"CPU_load","GraphID":"251"},{"Name":"Network_flow","GraphID":"252"}]}, {"Name":"Timely","Graphs":[{"Name":"CPU_load","GraphID":"251"},{"Name":"Network_flow","GraphID":"252"}]}}`)
json.Unmarshal(myServer, &hosts)
fmt.Println(len(hosts))
//  for i := range hosts {
//    fmt.Println(hosts[i][0])
//  }

}
