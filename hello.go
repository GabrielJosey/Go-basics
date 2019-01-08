// every Go file is part of a package. The package we are using in this file is the "main" package. 
// a Go program is split into several libraries called packages.
// Each file begins with a package declaration. 
// Packages are often used to set up libraries for other programs
// The special package that provides a starting point for Go programs to run, the "main" package


package main

// the "import" statement is to state what other packages it needs to import
// once the package is imported you can then use such things as variables, constants and such in your file

import "fmt" 

// after the import statement comes the actual package contents. This concists of 
// variables, functions and various other things you want your package to include


// when a Go program starts it looks for a function named "main" and invokes that. 
func main() {
  // we will now call a function from the format (fmt) package called Println
  // to use contents in a package you must call them from the package. 
  // for example to call "Println" we have to first call it from "fmt"
   fmt.Println("Hello, world!")

   // because otherFunction() is stated already in this package it does not need
   // to be qualified like fmt.Println() did. Println exists in the fmt package
   otherFunction()
}

func otherFunction() {
 // this is a change
}