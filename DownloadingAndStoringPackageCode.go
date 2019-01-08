package main

import (
	"fmt"
	//if the link below was not on this machine and needed to be downloaded then you could run "go get [url that needs to be downloaded]"
	//"github.com/treehouse-projects/go-intro/welcome"
)

func main() {
	fmt.Println(welcome.English)
	fmt.Println(welcome.Japanese)
	// If you un comment this line, you'll get the error "cannot refer to unexported name welcome.klingon"
	// fmt.Println(welcome.klingon)
}
