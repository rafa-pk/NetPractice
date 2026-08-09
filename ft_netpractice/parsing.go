package main
import (
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)


func PreliminaryChecks(file string) ([]string, int){
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Parsing Error: Cannot open file: %v.\n", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) == 0 {
		log.Fatalf("Parsing Error: Empty file.\n")
	}

	header := strings.Fields(lines[0])
	if len(header) != 2 || header[0] != "LAN" {
		log.Fatalf("Parsing Error: First line must be LAN <count>.\n")
	}
	deviceCount, err := strconv.Atoi(header[1])
	if err != nil {
		log.Fatalf("Parsing Error: LAN count must be an integer.\n")
	}
	return lines, deviceCount
}

func parseCIDR(ipv4 string) (uint32, error) {
	
	bytes, err := strconv.Atoi(ipv4[1:])
	if err != nil || bytes < 0 || bytes > 32 {
		fmt.Errorf("Invalid IPV4 address")
	}

	var mask uint32
	mask = ^uint32(0) << (32 - bytes)
	return mask, nil
}

func parseIPV4(ipv4 string) (uint32, error) {

	ipv4 = strings.TrimSpace(ipv4)
	if strings.HasPrefix(ipv4, "/") {
		return parseCIDR(ipv4)
	}

	octets := strings.Split(ipv4, ".")
	if len(octets) != 4 {
		log.Printf("%v\n", ipv4)
		return 0, fmt.Errorf("Invalid IPV4 address")
	}

	var address uint32
	for _, octet := range octets {
		byte, err := strconv.Atoi(octet)
		if err != nil {
			return 0, fmt.Errorf("Octet couldn't be converted.\n")
		}
		address = address << 8 | uint32(byte)
	}	
	return address, nil
}


func parseConfigFile(file string) []Device {	
	
	lines, deviceCount := PreliminaryChecks(file)
	var LAN []Device
	var i = -1

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
		case strings.HasPrefix(line, "interface"):
			_, name, _ := strings.Cut(line, ":")
			LAN[i].interfaces = append(LAN[i].interfaces, Interface{name: strings.TrimSpace(name)})
		case strings.HasPrefix(line, "ip"):
			ix := len(LAN[i].interfaces) - 1
			_, value, _	:= strings.Cut(line, ":")
			ip, err := parseIPV4(value)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
			LAN[i].interfaces[ix].ip = ip
		case strings.HasPrefix(line, "mask"):
			ix := len(LAN[i].interfaces) - 1
			_, value, _ := strings.Cut(line, ":")
			mask, err := parseIPV4(value)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
			LAN[i].interfaces[ix].mask = mask
		case strings.HasPrefix(line, "connection"):
			ix := len(LAN[i].interfaces) - 1
			_, connection, _ := strings.Cut(line, ":")
			LAN[i].interfaces[ix].connectedTo = connection
		case line == "route:":
			LAN[i].route = Route{}
		case strings.HasPrefix(line, "destination"):
			_, destRaw, _ := strings.Cut(line, ":")
			dest, err := parseIPV4(destRaw)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
			LAN[i].route.dest = dest
		case strings.HasPrefix(line, "r_mask"):
			_, maskRaw, _ := strings.Cut(line, ":")
			mask, err := parseIPV4(maskRaw)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n", err)
			}
			LAN[i].route.mask = mask
		case strings.HasPrefix(line, "gateway"):
			_, gw, _ := strings.Cut(line, ":")
			gateway, err := parseIPV4(gw)
			if err != nil {
				log.Fatalf("Parsing Error: %v.\n")
			}
			LAN[i].route.gateway = gateway
		default:
			log.Printf("Line: %v\n", line)
			log.Fatalf("Parsing Error: Line not recognised.\n")

		}
	}
	if len(LAN) != deviceCount {
		log.Fatalf("Parsing Error: LAN count does not match devices on network.\n")
	}
	return LAN
}