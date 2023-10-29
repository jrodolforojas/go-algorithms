package exercises

type Play struct {
	player1 string
	player2 string
}

func (exe *Exercises) rockPaperScissors(game []Play) (winner string) {
	movements := map[string]string{"🪨": "✂️", "✂️": "📄", "📄": "🪨"}
	player1Wins := 0
	player2Wins := 0
	for _, play := range game {
		player1Movement := play.player1
		player2Movement := play.player2

		if movements[player1Movement] == player2Movement {
			player1Wins += 1
		}
		if movements[player2Movement] == player1Movement {
			player2Wins += 1
		}
	}

	if player1Wins > player2Wins {
		return "player 1"
	} else if player2Wins > player1Wins {
		return "player 2"
	} else {
		return "tie"
	}
}
