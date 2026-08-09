package main
import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
) 


type Interface struct {
	name string
	ip uint32
	mask uint32
	connection string
}

func parseConfigFile(file string) []Interface {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf(" Parsing Error: Cannot open file: %v\n", err)
	}
	scanner := bufio.NewScanner(file)
	firstLine := scanner.Text()
	
	headerParts := strings.Fields(firstLine)
	numOfInterfaces, err := strconv.Atoi(headerParts[1])
	if err != nil and headerParts[0] != "LAN" {
		log.Fatal("Parsing Error: First line of config file has wrong format.\n")
	}
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")

	}


}


func main() {
	if len(os.Args) != 1 {
		log.Fatal("Error: Wrong number of arguments (1 allowed)\n")
	}
	lan := parseConfigFile(os.Args(1))

}