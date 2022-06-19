package main

import (
	"fmt"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	APIKEY := os.Getenv("APIKEY")
	client := golio.NewClient(APIKEY,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))

	summoner1, summoner2 := getSummoners(*client)
	matchlist1 := getMatchhistory(*client, summoner1)
	matchlist2 := getMatchhistory(*client, summoner2)

	getCommons(matchlist1, matchlist2)

}

func getCommons(matchlist1, matchlist2 []string) []string {
	var returnList []string
	for _, match1 := range matchlist1 {
		for _, match2 := range matchlist2 {
			if match1 == match2 {
				returnList = append(returnList, match1)
				break
			}
		}
	}
	for i, v := range returnList {
		fmt.Printf("%v - %v\n", i, v)

	}
	if len(returnList) == 0 {
		fmt.Printf("There is no Match")
	}
	return nil
}

func getSummoners(client golio.Client) (*lol.Summoner, *lol.Summoner) {
	summonerName1 := os.Getenv("SUMMONER1")
	summonerName2 := os.Getenv("SUMMONER2")

	summoner1, err := client.Riot.Summoner.GetByName(summonerName1)
	summoner2, err := client.Riot.Summoner.GetByName(summonerName2)
	if err != nil {
		log.Fatal(err)
	}
	return summoner2, summoner1
}

func getMatchhistory(client golio.Client, summoner *lol.Summoner) []string {

	matches, err := client.Riot.Match.List(summoner.PUUID, 0, 100)
	if err != nil {
		log.Fatal(err)
	}

	return matches

}
