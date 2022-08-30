package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type choice struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type win struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Wins []string
}
type played struct {
	Player int `json:"player"`
}
type response struct {
	Results  string `json:"results"`
	Player   int    `json:"player"`
	Computer int    `json:"computer"`
}

var wins = []win{
	{ID: 1, Name: "Rock", Wins: []string{"Lizard", "Scissors"}},
	{ID: 2, Name: "Paper", Wins: []string{"Rock", "Spock"}},
	{ID: 3, Name: "Scissors", Wins: []string{"Paper", "Lizard"}},
	{ID: 4, Name: "Lizard", Wins: []string{"Spock", "Paper"}},
	{ID: 5, Name: "Spock", Wins: []string{"Rock", "Scissors"}},
}
var choices = []choice{
	{ID: 1, Name: "Rock"},
	{ID: 2, Name: "Paper"},
	{ID: 3, Name: "Scissors"},
	{ID: 4, Name: "Lizard"},
	{ID: 5, Name: "Spock"},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", index)
	router.GET("/choices", getChoices)
	router.GET("/choice", getChoice)
	router.POST("/play", play)

	router.Run("localhost:8080")
}
func index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to RPSLS")
}
func play(c *gin.Context) {
	var input played
	if err := c.BindJSON(&input); err != nil {
		return
	}
	rand.Seed(time.Now().UnixNano())
	played := wins[input.Player-1]
	computer := wins[rand.Intn(6-1)+1]
	resp := response{Player: played.ID, Computer: computer.ID}

	if computer.Name == played.Name {
		resp.Results = "tie"
	} else if computer.Name == played.Wins[0] || computer.Name == played.Wins[1] {
		resp.Results = "win"
	} else if played.Name == computer.Wins[0] || played.Name == computer.Wins[1] {
		resp.Results = "lose"
	}

	c.IndentedJSON(http.StatusCreated, resp)
}
func getChoices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, choices)
}
func getChoice(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	c.IndentedJSON(http.StatusOK, choices[rand.Intn(6-1)+1])
}
