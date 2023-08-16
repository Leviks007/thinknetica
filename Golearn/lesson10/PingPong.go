package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type player struct {
	name  string
	score int
}

const (
	Ping     = "Ping"
	Pong     = "Pong"
	WinScore = 7
)

func pingPongGame(wg *sync.WaitGroup, ch chan string, playerNum int, players []*player) {
	defer wg.Done()
	for {
		msg := <-ch
		if msg == "stop" {
			return
		}

		fmt.Printf("%s: %s\n", players[playerNum].name, msg)
		time.Sleep(5 * time.Millisecond)

		if rand.Intn(100) < 20 {
			players[playerNum].score++
			fmt.Printf("%s: I won!\n", players[playerNum].name)
			ch <- "stop"
			return
		}
		if msg == Ping {
			ch <- Pong
		} else {
			ch <- Ping
		}

	}
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	players := []*player{
		{name: "Player 1", score: 0},
		{name: "Player 2", score: 0},
	}

	for players[0].score < WinScore && players[1].score < WinScore {
		for i := 0; i < len(players); i++ {
			wg.Add(1)
			go pingPongGame(&wg, ch, i, players)
		}

		for _, p := range players {
			fmt.Printf("%s: begin\n", p.name)
		}

		ch <- Ping

		wg.Wait()

		fmt.Println("Round finished!")
		for _, p := range players {
			fmt.Printf("%s score: %d\n", p.name, p.score)
		}
	}
	fmt.Println("Game finished!")
	for _, p := range players {
		if p.score == WinScore {
			fmt.Printf("\nWinner: %s\n", p.name)
		}
	}

}
