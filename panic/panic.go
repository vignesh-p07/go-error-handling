package main

import "fmt"

func main() {
	fmt.Println("Start of main")

	safeFunction()

	fmt.Println("End of main (program did not crash)")
}

// This function demonstrates panic and recover
func safeFunction() {
	defer func() {
		// This defer will run when panic happens
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	defer func() {
		// This defer will run when panic happens
		fmt.Println("this is defer 2")
	}()
	fmt.Println("About to panic...")

	panic("something went very wrong!")

	// This line will not run because panic stops the flow
	fmt.Println("This won't be printed")
}
