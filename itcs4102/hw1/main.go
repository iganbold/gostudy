package main

import (
        "log"
        "os"
        
        "github.com/iganbold/gostudy/itcs4102/hw1/search"
    )

func init() {
    log.SetOutput(os.Stdout)
}

func main() {
    search.Run("uncc")
}
