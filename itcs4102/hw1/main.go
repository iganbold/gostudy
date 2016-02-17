package main

import (
        "log"
        "os"
        
        "github.com/iganbold/gostudy/itcs4102/hw1/load"   //when you run the code, please make sure change the load folder path based on the GOPATH or the GOROOT 
    )

func init() {
    log.SetOutput(os.Stdout)
}

func main() {
    load.Run()
}
