package main

import "fmt"

func main()  {
	var lado uint64
  fmt.Scan(&lado)
  area := lado * lado
  fmt.Println(area)
}