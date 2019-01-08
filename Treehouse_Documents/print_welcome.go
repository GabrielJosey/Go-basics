package main

import (
  "fmt"
  "github.com/treehouse-projects/go-intro/welcome"
)

func main() {
	fmt.Println(welcome.English)
	fmt.Println(welcome.Japanese)
  // If you uncomment this line, you'll get the error "cannot refer to unexported name welcome.klingon"
	// fmt.Println(welcome.klingon)
}
