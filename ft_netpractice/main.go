package main
import (
	"os"
	"log"
) 


func main() {
	if len(os.Args) != 1 {
		log.Fatal("Error: Wrong number of arguments (1 allowed)\n")
	}
	lan := parseConfigFile(os.Args(1))

}