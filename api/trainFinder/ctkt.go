
package trainFinder

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/kryshhzz/krait/utils"
	"github.com/kryshhzz/krait/api"
)
 

func unmarshal_CTKT(body []byte, raillegs *[]utils.RailLeg) error { 

	type tRailLeg struct {
		ID          string `json:"trainNumber"`
		Name        string `json:"trainName"`
		Source      string `json:"fromStnCode"`
		Destination string `json:"toStnCode"`
		Arrival     string `json:"arrivalTime" `
		Departure   string `json:"departureTime" `

		Coaches interface{} `json:"avaiblitycache" `
	}

	type tResp struct {
		Body []tRailLeg `json:"trainBtwnStnsList"`
	}

	var apiresp tResp
	//  fmt.Println(string(body))
	err := json.Unmarshal(body, &apiresp)
	if err != nil {
		return err
	}

	actrls := []utils.RailLeg{}

	for _,v := range apiresp.Body { 
		tcs := []utils.Coach {}
		for _,ttc := range v.Coaches.(map[string]interface{}) {   
			ttcm,ok := ttc.(map[string]interface{}) 
			if !ok {
				fmt.Println("error converting interfafce to map")
			}else{
				if ttcm["Availability"] != nil {
					tc := utils.Coach {
						Type : ttcm["TravelClass"].(string),
						Closed : ttcm["Availability"].(string),
						Status : ttcm["AvailabilityDisplayName"].(string),
						Quota : ttcm["Quota"].(string),
						Fare: ttcm["Fare"],
					}
					tcs = append(tcs,tc)
				}
			}
		}
		
		t := utils.RailLeg{
			ID : v.ID,
			Name : v.Name,
			Source: v.Source,
			Destination: v.Destination,
			Arrival: v.Arrival,
			Departure: v.Departure, 

			Coaches: tcs,
		}
		actrls = append(actrls, t)
	}

	*raillegs = actrls
	return nil
}


// CTKT
func FindTrains_CTKT(date, dest, src string) (*[]utils.RailLeg, error) {

	fmt.Println("Finding Trains using CTKT")

	newdate := date[6:] + "-" + date[4:6] + "-" + date[:4]
	link := fmt.Sprintf("https://securedapi.confirmtkt.com/api/platform/trainbooking/tatwnstns?fromStnCode=%v&destStnCode=%v&doj=%v&token=&quota=GN&appVersion=290&androidid=mwebd_android", src, dest, newdate)
	resp, err := api.SendReq(link, "CTKT")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		fmt.Print()
	case 451:
		fmt.Printf("No trains : %v -> %v on %v \n", src, dest, date)
		return new([]utils.RailLeg), nil
	default:
		fmt.Println(link)
		return nil, fmt.Errorf("status Code not Suitable : %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	// fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	var raillegs []utils.RailLeg

	err = unmarshal_CTKT(body, &raillegs)
	if err != nil {
		panic(err)
		return nil, err
	}

	defer resp.Body.Close()
	return &raillegs, nil
}
