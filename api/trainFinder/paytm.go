package trainFinder

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/kryshhzz/krait/utils"
	"github.com/kryshhzz/krait/api"
)


func unmarshal_PAYTM(body []byte, raillegs *[]utils.RailLeg) error { 
	var apiresp utils.ApiResponse
	err := json.Unmarshal(body, &apiresp)
	if err != nil {
		return err
	}
	*raillegs = apiresp.Body.RailLegs
	return nil 
}

// PAYTM
func FindTrains_PAYTM(date, dest, src string) (*[]utils.RailLeg, error) { 

	fmt.Println("Finding Trains using PAYTM")

	link := fmt.Sprintf("https://travel.paytm.com/api/trains/v5/search?departureDate=%v&destination=%v&dimension114=direct-home&isH5=true&is_new_user=null&quota=GN&show_empty=true&source=%v&client=web&deviceIdentifier=Mozillafirefox-131.0.0.0", date, dest, src)
	resp, err := api.SendReq(link, "PAYTM")
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
	if err != nil {
		return nil, err
	}

	var raillegs []utils.RailLeg

	err = unmarshal_PAYTM(body, &raillegs)
	if err != nil {
		panic(err)
		return nil, err
	}

	defer resp.Body.Close()
	return &raillegs, nil
}
