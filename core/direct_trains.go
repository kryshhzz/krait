package core

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	// "time"
	// "math/rand"

	_"github.com/kryshhzz/krait/api"
	"github.com/kryshhzz/krait/api/trainFinder"
	"github.com/kryshhzz/krait/utils"
)


func FindDirectTrains(date,dest,src string) *[]utils.RailLeg { 

	var RedRailLegs []utils.RailLeg
	var err error

	log.Printf("Finding Direct Trains : %v -> %v \n", src,dest) 
	
	var railLegs *[]utils.RailLeg
	// rsrc := rand.NewSource(time.Now().UnixNano())
	// if rsrc.Int63() % 2 == 0 { 
	// 	yellow.Printf("Finding Trains using CTKT %v -> %v \n", src, dest)
	// 	railLegs,err = trainFinder.FindTrains2(date,dest,src) 
	// }else{
	// 	blue.Printf("Finding Trains using PAYTM %v -> %v \n", src, dest)
	// 	railLegs,err = trainFinder.FindTrains(date,dest,src) 
	// } 
	railLegs,err = trainFinder.FindTrains(date,dest,src) 

	
	if err != nil {
		fmt.Println(err)
		return nil
	}


	//log.Println("Data fetched successfully")
	
	for _,rl := range *railLegs { 
		// fmt.Printf("%v : %v ",rl.ID,rl.Name)  
		if len(utils.PREFFERED_TRAINS) > 0 && !utils.PREFFERED_TRAINS[rl.ID] {
			continue
		}
		fmt.Printf("%v : %v [ %v -> %v ] ",rl.ID, rl.Name, rl.Source, rl.Destination) 

		ZeroAvl := true 
		minPrice := float64(-1)
		cheapCoach := ""
		for _,c := range rl.Coaches { 
			if c.Quota != "TQ" && utils.PREFFERED_COACHES[c.Type] && c.Closed != "CLOSED" {  
				AVL := false
				if utils.PREFFERED_TICKET_STATUSES[strings.Split(c.Status," ")[0]] {
					AVL = true
				} 

				if !AVL { 
					//red.Printf("| %v %v |", c.Type, c.Status)  
				}else{
					green.Printf("| %v %v %v |", c.Type,c.Status, c.Fare)   
					
					ffare ,ok := c.Fare.(float64)
					if !ok {
						ffare,err = strconv.ParseFloat(c.Fare.(string),64)
						if err != nil {
							fmt.Println("fare convesion to float64 error")
							return nil
						}
					}
					
					if ( minPrice != float64(-1) && minPrice > ffare ) || minPrice == float64(-1){
						minPrice = ffare
						cheapCoach = c.Type
					}  
					ZeroAvl = false
				}

			} 
		}
		if ZeroAvl {
			RedRailLegs = append(RedRailLegs, rl)
		}else{
			j := utils.Journey{
					TotalFare : minPrice,
					Coaches : []string{cheapCoach},
					RailLegs : []utils.RailLeg{rl},
				}
			utils.TOTAL_JOURNEYS = append(utils.TOTAL_JOURNEYS,j)
			if ( utils.BEST_JOURNEY.TotalFare != float64(-1) && utils.BEST_JOURNEY.TotalFare > minPrice ) || utils.BEST_JOURNEY.TotalFare == float64(-1) {
				utils.BEST_JOURNEY = j
			}	
		}
		fmt.Println()
	}
	return &RedRailLegs
}