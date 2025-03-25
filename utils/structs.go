package utils


type Journey struct {
	TotalFare float64 
	Coaches []string
	RailLegs []RailLeg  
} 


type Coach struct {
	Type 	string `json:"code"` 
	Closed string `json:"status"`
	Status 	string `json:"status_shortform"` 
	UpdatedBefore 	interface{} `json:"time_diff,omitempty"`  // float64
	Quota string `json:"quota"`
	Fare interface{} `json:"fare" ` // float64
}

type RailLeg struct {
	ID 		string `json:"trainNumber"`
	Name    string `json:"trainName"`
	Source  string `json:"source"`
	SourceName  string `json:"source_name"`
	Destination string  `json:"destination"`
	DestinationName  string `json:"destination_name"`
	Arrival string `json:"arrival" `  
	Departure string `json:"departure" ` 

	Coaches []Coach `json:"availability" `
}


type ApiResponse struct {
	Status struct {
		Result string `json:"result"`
	} `json:"status"` 
	Body struct {
		RailLegs []RailLeg `json:"trains"`
	} `json:"body"`
}


type Station struct {
	Code string `json:"stationCode"` 
	Name string `json:"stationName"`
	DayCount string `json:"dayCount"`
	Arrival string `json:"arrivalTime"`
	Distance string `json:"distance"` 
}

type StationsApiResponse struct {
	Status struct {
		Result string `json:"result"`
	} `json:"status"` 
	Body struct {
		Stations []Station `json:"stationList"`
	} `json:"body"`
}


