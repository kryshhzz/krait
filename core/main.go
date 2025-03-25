package core

import (
	"sync"

	"github.com/fatih/color"

	"github.com/kryshhzz/krait/utils"
)

// find the trains running from src -> dest
// print the available trains

var red *color.Color
var green *color.Color
var yellow *color.Color
var blue *color.Color



func init() {
	red = color.New(color.FgRed)
	green = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	blue = color.New(color.FgHiBlue)
}

func Krait() { 
	
	if len(utils.PREFFERED_TRAINS) > 0 { 
		// if there are any preffered trains
		wg := sync.WaitGroup{}
		for pref_train_ID,_ := range utils.PREFFERED_TRAINS {
			wg.Add(1)
			rrl := utils.RailLeg {
				ID : pref_train_ID,
			}
			go func(rrl utils.RailLeg, wg *sync.WaitGroup) {
				FindPriorSourceTrains(&rrl)
				wg.Done()
			}(rrl,&wg)
		}
		wg.Wait()
	}else{
		RedRailLegs := FindDirectTrains(utils.DATE,utils.DEST,utils.SRC)
		if RedRailLegs != nil && len(*RedRailLegs) > 0 {
			wg := sync.WaitGroup{}
			for _, rrl := range *RedRailLegs {
				wg.Add(1)
				go func(rrl utils.RailLeg, wg *sync.WaitGroup) {
					FindPriorSourceTrains(&rrl)
					wg.Done()
				}(rrl,&wg)
			}
			wg.Wait()
		}
	}


	if len(utils.TOTAL_JOURNEYS) <= 0 {
		red.Println("NO TRAINS FOUND")
		return 
	}
	for _,j := range utils.TOTAL_JOURNEYS {
		green.Printf("=> : %v -> %v on %v is to go : %v -> %v  on %v for only %v rupees in [%v]; \n", utils.SRC, utils.DEST, utils.DATE, j.RailLegs[0].Source, j.RailLegs[0].Destination, j.RailLegs[0].Name, j.TotalFare, j.Coaches[0])
	}

	green.Printf("The best & cheapest way to go : %v -> %v on %v is to go : %v -> %v on %v for only %v rupees in [%v]; \n", utils.SRC, utils.DEST, utils.DATE, utils.BEST_JOURNEY.RailLegs[0].Source, utils.BEST_JOURNEY.RailLegs[0].Destination , utils.BEST_JOURNEY.RailLegs[0].Name, utils.BEST_JOURNEY.TotalFare,utils.BEST_JOURNEY.Coaches[0])

}
