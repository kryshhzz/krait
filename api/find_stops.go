package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/kryshhzz/krait/utils"
)

func FindStops2(date, ID, src string) (*[]utils.Station, error) {
	link := fmt.Sprintf("https://www.cleartrip.com/trains/%v/", ID)
	resp, err := SendReq(link,"")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		fmt.Print()
	case 403:
		fmt.Printf("No stations for : %v on %v \n",ID , date)
		return nil, fmt.Errorf("no stations found")
	default:
		fmt.Println(link)
		return nil, fmt.Errorf("status Code not Suitable : %v", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stations []utils.Station

	doc.Find("table tr").Each(func(index int, row *goquery.Selection) {
		columns := row.Find("td")
		if columns.Length() >= 7 { // Ensure there are enough columns

			splittedStatName := strings.Split(strings.TrimSpace(columns.Eq(1).Text()), "(")
			statCode := splittedStatName[1]

			stations = append(stations, utils.Station{
				Code:     statCode[:len(statCode)-1],
				Name:     splittedStatName[0],
				DayCount: strings.TrimSpace(columns.Eq(6).Text()),
				Arrival:  strings.TrimSpace(columns.Eq(2).Text()),
				Distance: strings.TrimSpace(columns.Eq(5).Text()),
			})

		}
	})

	fmt.Println(stations)

	defer resp.Body.Close()
	return &stations, nil
}

func FindStops(date, ID, src string) (*[]utils.Station, error) {
	link := fmt.Sprintf("https://travel.paytm.com/api/trains/v1/schedule?departureDate=%v&isH5=true&source=%v&trainNumber=%v&client=web&deviceIdentifier=Mozillafirefox-131.0.0.0", date, src, ID)
	resp, err := SendReq(link,"PAYTM")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	switch resp.StatusCode {
	case 200:
		fmt.Print()
	case 451:
		fmt.Printf("No stations for : %v on %v \n",ID , date)
		return nil, fmt.Errorf("no stations found")
	default:
		fmt.Println(link)
		return nil, fmt.Errorf("status Code not Suitable : %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var apiresp utils.StationsApiResponse

	err = json.Unmarshal(body, &apiresp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Println(apiresp)

	if apiresp.Status.Result != "success" {
		log.Println("Error Fetching Data")
		return nil, fmt.Errorf("Error fetching stations of : %v ", ID)
	}

	defer resp.Body.Close()
	return &apiresp.Body.Stations, nil
}
