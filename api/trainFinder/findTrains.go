package trainFinder

import (
	"math/rand"
	"time"

	"github.com/kryshhzz/krait/utils"
)

var trainsServiceProviders = []func(date, dest, src string) (*[]utils.RailLeg, error){
	FindTrains_CTKT,
	FindTrains_PAYTM,
}

func FindTrains(date,dest,src string) (*[]utils.RailLeg, error) {
	rand.Seed(time.Now().UnixNano()) // Seed to ensure different results each run
	randomNumber := rand.Intn(len(trainsServiceProviders))     // Generates a number between 0 and 5
	return trainsServiceProviders[randomNumber](date,dest,src)
}
