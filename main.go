package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
)

// Бот состоящий из массы и скорости
type Bot struct {
	mass     float64
	velocity float64
}

type IBot interface {
	getKineticEnergy() float64
}

func (r *Bot) getKineticEnergy() float64 {
	return (r.mass * math.Pow(r.velocity, 2)) * 0.5
}

func example(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mass, err := strconv.ParseFloat(vars["mass"], 64)
	if err != nil {
		log.Println("Error_1")
	}
	velocity, err := strconv.ParseFloat(vars["velocity"], 64)
	if err != nil {
		log.Println("Error_2")
	}
	bot := Bot{mass, velocity}
	jsonResponse, jsonError := json.Marshal(bot.getKineticEnergy())
	if jsonError != nil {
		log.Println("Error_3")
	}
	w.Header().Set("Counter-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/example/{mass:[0-9]+}/{velocity[0-9]+}/", example)
	log.Fatal(http.ListenAndServe(":8080", r))
}
