package main

import (
	"fmt"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
)

func main() {
	APIKEY := os.Getenv("APIKEY")
	client := golio.NewClient(APIKEY,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))

	summoner, _ := client.Riot.Summoner.GetByName("DreiAugenFlappe")
	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)
}
