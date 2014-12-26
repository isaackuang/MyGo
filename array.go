package main
import "fmt"
func main() {  
    x := []float64{
	23,
	45,
	33,
	21,
	55,
	88,
    }  

    var total float64 = 0  
    for _, value := range x {  
        total += value  
    }  
    fmt.Println(total / float64(len(x)))  
}  
