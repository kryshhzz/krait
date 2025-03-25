package core

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/kryshhzz/krait/api"
	"github.com/kryshhzz/krait/utils"
)

func FindPriorSourceTrains(rrl *utils.RailLeg) { 

	stations := new([]utils.Station)
	var err error 
	rsrc := rand.NewSource(time.Now().UnixNano())
	if rsrc.Int63() % 2 == 0 { 
		yellow.Printf("Finding Stations of %v\n", rrl.Name)
		stations,err = api.FindStops2(utils.DATE,rrl.ID,utils.SRC) 
		if err != nil {
			fmt.Println("Could not find stops : ",err)
			return
		}
	}else{ 
		blue.Printf("Finding Stations of %v\n", rrl.Name)
		stations,err = api.FindStops(utils.DATE,rrl.ID,utils.SRC) 
		if err != nil {
			fmt.Println("Could not find stops : ",err)
			return
		}
	}
	
	//log.Printf("Stations data fetched succesfully of %v \n",rrl.Name) 

	var srcStationID = 0
	var destStationID  = len(*stations)
	for i,station := range *stations{ 
		if station.Code == utils.SRC {
			srcStationID = i
		}
		if station.Code == utils.DEST {
			destStationID = i
		}
	}
  
	left := srcStationID 

	wg2 := sync.WaitGroup{}
	for {
		if left < 0 {
			break
		}
		for right := destStationID; right < len(*stations); right++ { 
			if ! utils.PRIOR_TRAINS_FOUND[(*stations)[left].Code+"->"+(*stations)[right].Code] {
				utils.MU.Lock()
				utils.PRIOR_TRAINS_FOUND[(*stations)[left].Code+"->"+(*stations)[right].Code]  = true
				utils.MU.Unlock()
				log.Printf("Prior Finding trains : %v -> %v \n",(*stations)[left].Name,(*stations)[right].Name)
				wg2.Add(1)
				go func(rcode,lcode string, wg2 *sync.WaitGroup){
					FindDirectTrains(utils.DATE,rcode,lcode)
					defer wg2.Done()
				}((*stations)[right].Code,(*stations)[left].Code,&wg2)
			}
		}
		left -= 1
	}
	wg2.Wait()
}
