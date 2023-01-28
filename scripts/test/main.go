// GOLANG PROGRAM TO CONVERT DATA TO HEXADECIMAL
package main

// fmt package allows us to print anything on the screen
// fmt.Sprint function is defined under the fmt package
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/takuoki/clmconv"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// start the function main ()
// this function is the entry point of the executable program
func main() {
	fmt.Println("Golang Program to convert data to hexadecimal")

	// initialize the integer variable
	int_value := 1

	// calculate the hex value by calling the function fmt.Sprintf()
	// %x prints the hexadecimal characters in lowercase
	hex_value := fmt.Sprintf("%x", int_value)
	fmt.Printf("Hex value of %d is = %s\n", int_value, hex_value)

	// %X prints the hexadecimal characters in uppercase
	hex_value = fmt.Sprintf("%X", int_value)
	fmt.Printf("Hex value of %d is = %s\n", int_value, hex_value)

	// initialize the integer variable
	int_value = 27

	// calculate the hex value by calling the function fmt.Sprintf()
	// %x prints the hexadecimal characters in lowercase
	hex_value = fmt.Sprintf("%x", int_value)
	fmt.Printf("Hex value of %d is = %s\n", int_value, hex_value)

	// %X prints the hexadecimal characters in uppercase
	hex_value = fmt.Sprintf("%X", int_value)
	fmt.Printf("Hex value of %d is = %s\n", int_value, hex_value)

	converter := clmconv.New(clmconv.WithStartFromOne(), clmconv.WithLowercase())
	// a := converter.Itoa(1) // a = "a"
	// b := converter.Itoa(26) // a = "a"
	// c := converter.Itoa(77) // a = "a"

	log.Println("clmconv:", converter.Itoa(1))
	log.Println("clmconv:", converter.Itoa(26))
	log.Println("clmconv:", converter.Itoa(27))
	log.Println("clmconv:", converter.Itoa(28))
	log.Println("clmconv:", converter.Itoa(77))

	log.Println("GetMD5Hash:", GetMD5Hash("123456"))
	// print the result
}
