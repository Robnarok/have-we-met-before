package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/matches", matches)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
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
	return returnList
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

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("Template/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func matches(w http.ResponseWriter, r *http.Request) {
	APIKEY := os.Getenv("APIKEY")
	client := golio.NewClient(APIKEY,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	summoner1, summoner2 := getSummoners(*client)
	matchlist1 := getMatchhistory(*client, summoner1)
	matchlist2 := getMatchhistory(*client, summoner2)

	matches := getCommons(matchlist1, matchlist2)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range matches {
		fmt.Fprintf(w, "%v - %v\n", i, v)
	}
	if len(matches) == 0 {
		fmt.Fprintf(w, "There is no Match")
	} else {
		fmt.Fprintf(w, "You played %v matches together in the last 100 Games!",
			len(matches))
	}

}
