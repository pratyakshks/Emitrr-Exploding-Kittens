package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
var upgrader = websocket.Upgrader{}

type GameState struct {
	Username   string   `json:"username"`
	Deck       []string `json:"deck"`
	Defuses    int      `json:"defuses"`
	Won        int      `json:"won"`
	InProgress bool     `json:"inProgress"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan map[string]interface{})

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Helper to shuffle cards
func shuffleDeck() []string {
	cards := []string{"😼", "🙅‍♂️", "🔀", "💣", "😼"}
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

// API to start/restart a game
func startGame(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	ctx := context.Background()

	state := GameState{
		Username:   username,
		Deck:       shuffleDeck(),
		Defuses:    0,
		Won:        0,
		InProgress: true,
	}

	stateJSON, _ := json.Marshal(state)
	rdb.Set(ctx, username, stateJSON, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(stateJSON)
}

// API to draw a card
func drawCard(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	ctx := context.Background()

	val, err := rdb.Get(ctx, username).Result()
	if err != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var state GameState
	json.Unmarshal([]byte(val), &state)

	if len(state.Deck) == 0 || !state.InProgress {
		http.Error(w, "Game is over", http.StatusBadRequest)
		return
	}

	card := state.Deck[0]
	state.Deck = state.Deck[1:]

	if card == "💣" {
		if state.Defuses > 0 {
			state.Defuses--
		} else {
			state.InProgress = false
			card = "Game Over 💣"
		}
	} else if card == "🙅‍♂️" {
		state.Defuses++
	} else if card == "🔀" {
		state.Deck = shuffleDeck()
	} else if len(state.Deck) == 0 {
		state.Won++
		state.InProgress = false
		card = "You Won 🎉"
	}

	stateJSON, _ := json.Marshal(state)
	rdb.Set(ctx, username, stateJSON, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"card": "%s", "state": %s}`, card, stateJSON)))

	// Update leaderboard
	updateLeaderboard(state.Username, state.Won)
}

// WebSocket handler for real-time updates
func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg map[string]interface{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			break
		}
		broadcast <- msg
	}
}

// Broadcast leaderboard updates
func updateLeaderboard(username string, points int) {
	for client := range clients {
		client.WriteJSON(map[string]interface{}{
			"username": username,
			"points":   points,
		})
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/start/{username}", startGame).Methods("POST")
	r.HandleFunc("/draw/{username}", drawCard).Methods("POST")
	r.HandleFunc("/ws", wsHandler)

	go func() {
		for {
			msg := <-broadcast
			for client := range clients {
				client.WriteJSON(msg)
			}
		}
	}()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
