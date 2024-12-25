package lem

var (
	Ways       = make(map[string][]string)
	Emptyroom  = make(map[string]bool)
	checklinks =make(map[string]bool)
	Start, End string
	Ants       int
	seen       = make(map[string]bool)
	// max        = 0
	// badrooms   = []string{}
	visited   = make(map[string]bool)
	solutions [][]string
)
