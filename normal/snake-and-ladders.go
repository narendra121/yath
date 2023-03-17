package normal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var ladders = [][]int{{1, 4, 8, 21, 28, 50, 71, 80}, {38, 14, 30, 42, 76, 67, 92, 99}}
var snakes = [][]int{{32, 36, 48, 62, 88, 95, 97}, {10, 6, 26, 18, 24, 56, 78}}

func GetPosition(inp int, usrPos int) int {
	temp := usrPos
	usrPos += inp
	snake := 0
	ladder := 0
	if usrPos > 100 {
		fmt.Println("value exceeded max limit of 100")
		return temp
	}
	for j := 0; j < len(ladders[0]); j++ {
		if ladders[0][j] == usrPos {
			ladder = ladders[1][j]
		}
	}
	for j := 0; j < len(snakes[0]); j++ {
		if snakes[0][j] == usrPos {
			snake = snakes[1][j]
		}
	}
	if snake > 0 {
		return snake
	} else if ladder > 0 {
		return ladder
	} else {
		return usrPos
	}
}

func StartSnakesAndLaddersGame() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time
	people := 0
	// generate a random integer between 1 and 6 (inclusive)
	fmt.Println("Enter Number of people: ")
	fmt.Scan(&people)
	dropperCount := people
	ScoreBoard := make([][]int, 2)
	for i := range ScoreBoard {
		ScoreBoard[i] = make([]int, people)
	}
	k := 0
	for j := 0; j < people; j++ {
		k += 1
		ScoreBoard[0][j] = k
	}

	fmt.Println(ScoreBoard)

	for {
		for j := 0; j < people; j++ {
			inp := rand.Intn(6) + 1
			check := ""
			if dropperCount == 1 {
				for j := 0; j < people; j++ {
					if ScoreBoard[1][j] >= 0 {
						fmt.Println()

						fmt.Println("THE WINNER IS ", ScoreBoard[0][j])
						return
					}
				}
			}
			for !strings.EqualFold(check, "y") || !strings.EqualFold(check, "d") || !strings.EqualFold(check, "s") {
				fmt.Println()
				if ScoreBoard[1][j] < 0 {
					fmt.Print("Player:", ScoreBoard[0][j], " Dropped from game")
					break
				}
				fmt.Print("Player:", ScoreBoard[0][j], " Please type  y  to roll the dice or d to display score  or s to drop from the game:")
				fmt.Scan(&check)
				if strings.EqualFold(check, "y") {
					fmt.Println("dice value is: ", inp)
					ScoreBoard[1][j] = GetPosition(inp, ScoreBoard[1][j])

					fmt.Println(ScoreBoard[0][j], " current position is: ", ScoreBoard[1][j])
					if ScoreBoard[1][j] == 100 {
						fmt.Println("THE WINNER IS ", ScoreBoard[0][j])
						return
					}
					break
				} else if strings.EqualFold(check, "d") {
					fmt.Println(ScoreBoard)
				} else if strings.EqualFold(check, "s") {
					ScoreBoard[1][j] = -1
					fmt.Println("Droppeing Player: ", ScoreBoard[0][j])
					dropperCount -= 1
				} else {
					fmt.Println("wrong input")

				}
			}
		}

	}
}
