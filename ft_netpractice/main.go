package main
import (
	"os"
	"log"
	"fmt"
) 


func main() {
	if len(os.Args) != 2 {
		log.Fatal("Error: Wrong number of arguments (1 allowed)\n")
	}
	lan := parseConfigFile(os.Args[1])
	fmt.Printf("%+v\n", lan)
}