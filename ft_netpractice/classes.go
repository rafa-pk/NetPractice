package main


type Interface struct {
	name string
	ip uint32
	mask uint32
	connectedTo string
}

type Route struct {
	dest uint32
	mask uint32
	gateway uint32
}

type Device struct {
	name string
	interfaces []Interface
	routes []Route
}