package main

import (
  "os"
  "fmt"
  "github.com/xuri/excelize/v2"
)

func main() {

  args := os.Args[1:]

  f, err := excelize.OpenFile(args[0])
  if err != nil {
    fmt.Println(err)
    return
  }

  err = f.SetCellFormula(args[1], args[2], args[3])
  if err != nil {
    fmt.Println(err)
    return
  }

  err = f.Save()
  if err != nil {
    fmt.Println(err)
  }

}