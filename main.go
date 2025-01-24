package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Player struct {
	ID         int
	Name       string
	Wins       int
	Byes       int
	DirectWins map[int]bool // Tracks which players this player has beaten directly
	Buchholz   int          // Buchholz score
}

type Match struct {
	Player1 Player
	Player2 Player
	Winner  int // 1 for Player1, 2 for Player2, 0 for draw
	IsBye   bool
}

type Tournament struct {
	Players    []Player
	Rounds     [][]Match
	MatchPairs map[string]bool // Track played matchups as "ID1-ID2"
}

func NewTournament(playerNames []string) *Tournament {
	players := make([]Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = Player{
			ID:         i + 1,
			Name:       name,
			Wins:       0,
			Byes:       0,
			DirectWins: make(map[int]bool), // Initialize map
			Buchholz:   0,
		}
	}
	return &Tournament{
		Players:    players,
		Rounds:     make([][]Match, 0),
		MatchPairs: make(map[string]bool),
	}
}

func (t *Tournament) PairPlayers() []Match {
	// Shuffle players within the same win group for randomization
	t.randomizeTiedPlayers()

	// Sort players by wins (descending)
	sort.SliceStable(t.Players, func(i, j int) bool {
		return t.Players[i].Wins > t.Players[j].Wins
	})

	matches := []Match{}
	used := make(map[int]bool)

	// Handle bye if the number of players is odd
	if len(t.Players)%2 != 0 {
		byePlayer := t.assignBye()
		fmt.Printf("Jogador %s ganha um bye esse round.\n", byePlayer.Name)
		matches = append(matches, Match{
			Player1: byePlayer,
			IsBye:   true,
		})
		used[byePlayer.ID] = true
	}

	for i := 0; i < len(t.Players); i++ {
		if used[t.Players[i].ID] {
			continue
		}

		for j := i + 1; j < len(t.Players); j++ {
			if used[t.Players[j].ID] {
				continue
			}

			matchKey := fmt.Sprintf("%d-%d", t.Players[i].ID, t.Players[j].ID)
			if t.MatchPairs[matchKey] {
				continue // Skip if these players have already faced each other
			}

			matches = append(matches, Match{
				Player1: t.Players[i],
				Player2: t.Players[j],
			})
			used[t.Players[i].ID] = true
			used[t.Players[j].ID] = true

			// Record this matchup
			t.MatchPairs[matchKey] = true
			t.MatchPairs[fmt.Sprintf("%d-%d", t.Players[j].ID, t.Players[i].ID)] = true
			break
		}
	}

	return matches
}

func (t *Tournament) randomizeTiedPlayers() {
	// Group players by wins
	winGroups := make(map[int][]Player)
	for _, player := range t.Players {
		winGroups[player.Wins] = append(winGroups[player.Wins], player)
	}

	// Shuffle players within each win group
	rand.Seed(time.Now().UnixNano())
	shuffledPlayers := []Player{}
	for _, group := range winGroups {
		rand.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
		shuffledPlayers = append(shuffledPlayers, group...)
	}

	// Replace tournament players with the shuffled result
	t.Players = shuffledPlayers
}

func (t *Tournament) assignBye() Player {
	// Select the player with the fewest wins who hasn't had a bye
	sort.SliceStable(t.Players, func(i, j int) bool {
		if t.Players[i].Byes == t.Players[j].Byes {
			return t.Players[i].Wins < t.Players[j].Wins
		}
		return t.Players[i].Byes < t.Players[j].Byes
	})

	for i := range t.Players {
		if t.Players[i].Byes == 0 {
			t.Players[i].Byes++
			t.Players[i].Wins++ // Award a win for the bye
			return t.Players[i]
		}
	}

	// If all players have had a bye, give another bye to the next eligible player
	t.Players[0].Byes++
	t.Players[0].Wins++
	return t.Players[0]
}

func (t *Tournament) RecordResults(matches []Match) {
	for _, match := range matches {
		if match.IsBye {
			continue // Byes are already recorded
		}

		switch match.Winner {
		case 1:
			for i := range t.Players {
				if t.Players[i].ID == match.Player1.ID {
					t.Players[i].Wins++
					t.Players[i].DirectWins[match.Player2.ID] = true // Player1 beats Player2
				}
			}
		case 2:
			for i := range t.Players {
				if t.Players[i].ID == match.Player2.ID {
					t.Players[i].Wins++
					t.Players[i].DirectWins[match.Player1.ID] = true // Player2 beats Player1
				}
			}
		}
	}
	t.Rounds = append(t.Rounds, matches)
}

func (t *Tournament) RecalculateBuchholz() {
	// Recalculate Buchholz score for each player
	for i := range t.Players {
		t.Players[i].Buchholz = 0 // Reset Buchholz score
		for j := range t.Players {
			if i != j && t.Players[i].DirectWins[t.Players[j].ID] {
				t.Players[i].Buchholz += t.Players[j].Wins
			}
		}
	}
}

func (t *Tournament) PrintStandings() {
	// Recalculate Buchholz after the tournament
	t.RecalculateBuchholz()

	// Sort players by wins and then Buchholz
	sort.SliceStable(t.Players, func(i, j int) bool {
		if t.Players[i].Wins == t.Players[j].Wins {
			// Tiebreaker: First check direct match result
			if t.Players[i].DirectWins[t.Players[j].ID] {
				return true // Player i wins due to direct match
			}
			if t.Players[j].DirectWins[t.Players[i].ID] {
				return false // Player j wins due to direct match
			}

			// Tiebreaker: Use Buchholz score if wins and direct matches are tied
			if t.Players[i].Buchholz != t.Players[j].Buchholz {
				return t.Players[i].Buchholz > t.Players[j].Buchholz // Higher Buchholz score wins
			}

			// If still tied, use player ID for tie-breaking
			return t.Players[i].ID < t.Players[j].ID
		}
		return t.Players[i].Wins > t.Players[j].Wins
	})

	fmt.Println("Classificação:")
	for _, player := range t.Players {
		fmt.Printf("ID: %d, Name: %s, Wins: %d, Byes: %d, Buchholz: %d\n", player.ID, player.Name, player.Wins, player.Byes, player.Buchholz)
	}
}

func main() {
	playerNames := []string{"Anthony", "Renata", "Henrique", "Sofia", "Thiagão", "Arthur", "Renatinho"}
	tournament := NewTournament(playerNames)

	for round := 1; round <= 3; round++ { // Simulate 3 rounds
		fmt.Printf("Round %d:\n", round)
		matches := tournament.PairPlayers()

		if len(matches) == 0 {
			fmt.Println("Fim do torneio.")
			break
		}

		// Print all the matches of the round
		for i, match := range matches {
			if match.IsBye {
				fmt.Printf("Partida %d: %s gets a bye.\n", i+1, match.Player1.Name)
			} else {
				fmt.Printf("Partida %d: %s vs %s\n", i+1, match.Player1.Name, match.Player2.Name)
			}
		}

		// Now ask for the result after printing all matches
		for i, match := range matches {
			if match.IsBye {
				continue // Skip byes as we don't need to ask for a result
			}

			// Clear and specific prompt for the user
			fmt.Printf("Informe o resultado da partida %d: %s vs %s\n", i+1, match.Player1.Name, match.Player2.Name)
			fmt.Print("Informe o resultado: (1 for Player1, 2 for Player2, 0 for draw): ")
			fmt.Scan(&matches[i].Winner)
		}

		// Record the results and print standings after each round
		tournament.RecordResults(matches)
		tournament.PrintStandings()
	}

	// Final standings after recalculating Buchholz scores
	fmt.Println("\nClassificação final:")
	tournament.PrintStandings()
}
