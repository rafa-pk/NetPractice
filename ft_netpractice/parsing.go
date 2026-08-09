package main
import (
	"os"
	"strings"
	"strconv"
)


func PreliminaryChecks() ([]string, int){
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Parsing Error: Cannot open file: %v.\n", err)
	}

	lines := strings.Split(content, "\n")
	if len(lines) == 0 {
		log.Fatalf("Parsing Error: Empty file.\n")
	}

	header := strings.Fields(lines[0])
	if len(header) != 2 || header[0] != "LAN" {
		log.Fatalf("Parsing Error: First line must be LAN <count>.\n")
	}
	deviceCount, err := strconv.Atoi(lines[1])
	if err != nil {
		log.Fatalf("Parsing Error: LAN count must be an integer.\n")
	}
	return (lines, deviceCount)
}


func parseConfigFile(file string) []Device {	
	
	lines, deviceCount := PreliminaryChecks()
	var LAN []Device
	var i := -1

	for _, rawLine := range lines[1:] {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		switch {
		case strings.HasPrefix(line, "device"):
			_, name, _ := strings.Cut(line, ":")
			LAN = append(LAN, Device{name: strings.TrimSpace(name)})
			i = len(LAN) - 1
		case line == "interface":
			_, name, _ := strings.Cut(line, ":")
			LAN[i].interfaces = append(LAN[i].interfaces, Interface{name: strings.TrimSpace(interfaceName)})
		case strings.HasPrefix(line, "ip"):
			ix := len(LAN.interfaces) - 1
			_, value, _	= strings.Cut(line, ":")
			ip, err = parseIPV4(value)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
			LAN[i].interfaces[ix].ip = ip
		case strings.HasPrefix(line, "mask"):
			ix := len(LAN.interfaces[ix]) - 1
			_, value, _ := strings.Cut(line, ":")
			mask, err := parseIPV4(value)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
		case strings.HasPrefix(line, "connection"):
			ix := len(LAN.interfaces[ix]) - 1
			_, connection, _ := strings.Cut(line, ":")
			LAN[i].interfaces[ix].connectedTo = connection
		case strings.HasPrefix(line, "destination"):

		default:
			log.Fatalf("Parsing Error: Line not recognised.\n")

		}
	}
	if len(LAN) != deviceCount {
		log.Fatalf("Parsing Error: LAN count does not match devices on network.\n")
	}
	return LAN
}