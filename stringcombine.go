package main

import (
    "bufio"
    "bytes"
    //"fmt"
    //"io/ioutil"
    "os"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var buffer bytes.Buffer
    var gid string = "451"

        buffer.WriteString("|")
        buffer.WriteString("![graph](http://cacti.hiiir.com/cacti/graph_image.php?action=view&local_graph_id=")
        buffer.WriteString(gid)
        buffer.WriteString(")")
    f, err := os.Create("/tmp/test.md")
    check(err)
    defer f.Close()
    w := bufio.NewWriter(f)
    w.WriteString(buffer.String())


    w.Flush()

    //fmt.Println(buffer.String())
}
