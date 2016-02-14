package main

import "fmt"

func main() {
    m := make(map[string]int)
    
    m["k1"] = 5
    m["k2"] = 10
    
    fmt.Println("map: ", m)
    
    v1 := m["k1"]
    fmt.Println("v1: ",v1)
    
    fmt.Println("len: ",len(m))
    
    delete(m,"k2")
    fmt.Println("map: ",m)
    
    _, exists := m["k2"]
    fmt.Println("exists: ",exists)
    
    n := map[string]int{"foo":1,"bar":2}
    fmt.Println("map: ",n)
    
}
