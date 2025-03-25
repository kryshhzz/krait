package utils

import "sync"

var DATE = ""
var SRC = ""
var DEST = ""

var MU sync.Mutex

var BEST_JOURNEY = Journey{
	-1,
	[]string{},
	[]RailLeg{},
}

var TOTAL_JOURNEYS = []Journey{} 

var PREFFERED_TRAINS = map[string]bool{}

var PRIOR_TRAINS_FOUND = map[string]bool{}

var PREFFERED_COACHES = map[string]bool{
	"SL": true,
	"3E": true,
	"3A": true,
	"2A": true,
	"1A": true,
	"CC": true,
	"EC": true,
	"2S": true,
}

var PREFFERED_TICKET_STATUSES = map[string]bool{
	"AVL": true,
	"RAC": true,
}

var LINK = ""
