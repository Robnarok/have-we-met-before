# have-we-met-before

> **WARNING**
> This complete tool is still work in Progress. It may not work     

This is yet another short Programm in GO for League of Legends. Its purpose is
to find out if you already played a game with someone... It only checks the last
few Matches and prints MatchIDs of games you both played.

## Key Features

- Nothing, everything is still work in Progress

## Upcoming Features and Ideas

- Lists all games two players have common in their Match History
- Accepts the Names via Webinterface, Commandline or API
- ???

## How to Use ?

### [Recommended] with Docker

``` bash
# Clone this Repository
$ git clone https://github.com/Robnarok/have-we-met-before.git

# Go into the Repository
$ cd ./have-we-met-before

# Create .env file from template
$ cp template.env .env

# Edit .env file to set the APIKEY Variable
...

# Run with Docker-compose
$ docker-compose up --build -d

# Visit Web Interface for the Application and start using
```

### Without Docker

``` bash

# Clone this Repository
$ git clone https://github.com/Robnarok/have-we-met-before.git

# Go into the Repository
$ cd ./have-we-met-before

# Set Environment Variables
$ export APIKEY="XXX"
$ export SUMMONER1="XXX"
$ export SUMMONER2="XXX"

# Get the dependencies
$ go get .

# Run the Application
$ go run .

```

## Credits

- [Golang](https://go.dev/)
- [Golio](https://github.com/KnutZuidema/golio)
- ...

## Contribution

If you have some cool Ideas, do hesitate to open an Issue with a great idea.
This complete Project is only for practice, its intention is not to be commonly
used.. But if you actually use it: just let me know :)

If there is a bug or maybe a big issue, Please open an Github Issue - ill be
happy to fix it, if i have the time to do it
