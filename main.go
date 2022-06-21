package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/matches", matches)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Infof("Webserver is now Running!")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("Error Starting the HTTP Server : ", err)
			return
		}
	}()
	<-done
	log.Printf("Server stopped!")
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

func getSummoners(client golio.Client, summonerName1, summonerName2 string) (*lol.Summoner, *lol.Summoner, error) {
	summoner1, err := client.Riot.Summoner.GetByName(summonerName1)
	if err != nil {
		return nil, nil, fmt.Errorf("getSummoners(%v): %v", summonerName1, err)
	}
	summoner2, err := client.Riot.Summoner.GetByName(summonerName2)
	if err != nil {
		return nil, nil, fmt.Errorf("getSummoners(%v): %v", summonerName2, err)
	}
	return summoner1, summoner2, nil
}

func getMatchhistory(client golio.Client, summoner *lol.Summoner) ([]string, error) {
	matches, err := client.Riot.Match.List(summoner.PUUID, 0, 100)
	if err != nil {
		return nil, fmt.Errorf("getMatchhistory: %v", err)
	}
	return matches, nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("Template/index.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Warningf("Error executing template: %v", err)
		return
	}
}

func matches(w http.ResponseWriter, r *http.Request) {
	APIKEY := os.Getenv("APIKEY")
	client := golio.NewClient(APIKEY,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	summonerName1 := r.FormValue("summ1")
	summonerName2 := r.FormValue("summ2")
	log.Infof("Search for %v and %vs mutal Matches", summonerName1, summonerName2)
	summoner1, summoner2, err := getSummoners(*client, summonerName1, summonerName2)
	if err != nil {
		log.Warningf("matches: %v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	matchlist1, err := getMatchhistory(*client, summoner1)
	if err != nil {
		log.Warningf("matches: %v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	matchlist2, err := getMatchhistory(*client, summoner2)
	if err != nil {
		log.Warningf("matches: %v", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	matches := getCommons(matchlist1, matchlist2)
	err = r.ParseForm()
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
