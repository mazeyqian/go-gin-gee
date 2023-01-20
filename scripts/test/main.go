// Golang program to show the usage of
// Setenv(), Getenv and Unsetenv method

package main

import (
	"fmt"
	"os"
)

// Main function
func main() {

	// set environment variable GEEKS
	os.Setenv("GEEKS", "geeks")

	// returns value of GEEKS
	fmt.Println("GEEKS:", os.Getenv("GEEKS"))

	// Unset environment variable GEEKS
	os.Unsetenv("GEEKS")

	// returns empty string and false,
	// because we removed the GEEKS variable
	value, ok := os.LookupEnv("GEEKS")

	fmt.Println("GEEKS:", value, " Is present:", ok)

}
