package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

const master = "[GameMaster]"

func main() {
	fmt.Println(master + "Welcome tom tic-tac-toe game!")
	time.Sleep(time.Second * 1)
	gameBoard := makingGameBoard()
	printGameBoard(gameBoard)
	expretionGame()
	gameManager(gameBoard)
}

func makingGameBoard() [][]string {
	gameBoard := [][]string{
		[]string{" ", "1", "2", "3", "4", "5"},
		[]string{"A", "-", "-", "-", "-", "-"},
		[]string{"B", "-", "-", "-", "-", "-"},
		[]string{"C", "-", "-", "-", "-", "-"},
		[]string{"D", "-", "-", "-", "-", "-"},
		[]string{"E", "-", "-", "-", "-", "-"},
	}

	return gameBoard
}

func expretionGame() {
	// time.Sleep(time.Second * 1)
	fmt.Println(master + "Choose box like A1")
	// time.Sleep(time.Second * 2)
	fmt.Println(master + "You -> 0: CPU -> X")
	// time.Sleep(time.Second * 2)
	fmt.Println(master + "Game start!!")
	// time.Sleep(time.Second * 2)
}

func printGameBoard(gameBoard [][]string) {
	fmt.Printf("-----------\n")
	fmt.Printf("%sCurrent gameboard is this!!\n", master)
	for i := range gameBoard {
		fmt.Printf("%s\n", strings.Join(gameBoard[i], " "))
	}

	fmt.Printf("-----------\n")
}

func gameManager(gameBoard [][]string) {
	userManager(gameBoard)
	cpuManager(gameBoard)
}

func userManager(gameBoard [][]string) {
	user_input := getInput()
	i, v := boxConverter(gameBoard, user_input)
	isCorrect := checkGameBoard(gameBoard[i][v], i, v)

	if isCorrect == false {
		fmt.Printf(master + "your input is invalid.\n")
		time.Sleep(time.Second * 2)
		fmt.Printf(master + "retry your input.\n")
		time.Sleep(time.Second * 2)
		userManager(gameBoard)
	}

	gameBoard[i][v] = "0"
	time.Sleep(time.Second * 2)
	printGameBoard(gameBoard)
	judgment(gameBoard)
}

func cpuManager(gameBoard [][]string) {
	cpuDeterminer(gameBoard)
	judgment(gameBoard)
	gameManager(gameBoard)
}

func getInput() string {
	fmt.Println(master + "your turn.")
	time.Sleep(time.Second * 2)
	fmt.Print(master + "choose your box -> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	user_input := scanner.Text()
	fmt.Println(master + "you choose " + user_input)

	if user_input == "exit" {
		endGame("draw")
	}

	return user_input
}

func checkGameBoard(targetBox string, i int, v int) bool {
	if targetBox != "-" {
		return false
	}
	if i == 0 || v == 0 {
		return false
	}

	return true
}

func boxConverter(gameBoard [][]string, user_input string) (int, int) {
	i := rowConverter(user_input[:1])
	v := columnConverter(user_input[1:])

	return i, v
}

func rowConverter(row string) int {
	i := 0

	switch row {
	case "A":
		i = 1
	case "B":
		i = 2
	case "C":
		i = 3
	case "D":
		i = 4
	case "E":
		i = 5
	default:
		i = 0
	}
	return i
}

func columnConverter(column string) int {
	v := 0

	switch column {
	case "1":
		v = 1
	case "2":
		v = 2
	case "3":
		v = 3
	case "4":
		v = 4
	case "5":
		v = 5
	default:
		v = 0
	}
	return v
}

func cpuDeterminer(gameBoard [][]string) [][]string {
	fmt.Println(master + "CPU is thinking ...")

	const letters = "ABCDE"
	couldLoop := true

	for couldLoop {
		// 乱数を生成
		rand.Seed(time.Now().Unix())
		i := rand.Intn(5)
		i++

		time.Sleep(time.Millisecond * 300)
		rand.Seed(time.Now().Unix())
		v := rand.Intn(5)
		v++

		if gameBoard[i][v] == "-" {
			gameBoard[i][v] = "X"
			fmt.Printf("%scpu choose %s%s\n", master, gameBoard[i][0], gameBoard[0][v])
			couldLoop = false
		}
	}
	time.Sleep(time.Second * 2)
	printGameBoard(gameBoard)

	return gameBoard
}

func judgment(gameBoard [][]string) {
	horizontalBoard := make([][]string, len(gameBoard))
	verticalBoard := make([][]string, len(gameBoard))
	diagonalBoard := make([][]string, len(gameBoard))
	reverseDiagonalBoard := make([][]string, len(gameBoard))

	winner := ""
	count := 0

	userWinningPattern := []string{"0", "0", "0", "0", "0"}
	cpuWinningPattern := []string{"x", "x", "x", "x", "x"}

	for i := range gameBoard {

		for v := range gameBoard {

			if i > 0 && v > 0 {
				horizontalBoard[i] = append(horizontalBoard[i], gameBoard[i][v])
				verticalBoard[i] = append(verticalBoard[i], gameBoard[v][i])

				if i == v {
					diagonalBoard[0] = append(diagonalBoard[0], gameBoard[i][v])
				}

				if i+v == 6 {
					reverseDiagonalBoard[0] = append(reverseDiagonalBoard[0], gameBoard[i][v])
				}

				count = count + strings.Count(gameBoard[i][v], "-")
			}
		}

		if reflect.DeepEqual(horizontalBoard[i], userWinningPattern) {
			winner = "user"
			endGame(winner)
		}

		if reflect.DeepEqual(verticalBoard[i], userWinningPattern) {
			winner = "user"
			endGame(winner)
		}

		if reflect.DeepEqual(diagonalBoard[0], userWinningPattern) {
			winner = "user"
			endGame(winner)
		}

		if reflect.DeepEqual(reverseDiagonalBoard[0], userWinningPattern) {
			winner = "user"
			endGame(winner)
		}

		if reflect.DeepEqual(horizontalBoard[i], cpuWinningPattern) {
			winner = "cpu"
			endGame(winner)
		}

		if reflect.DeepEqual(verticalBoard[i], cpuWinningPattern) {
			winner = "cpu"
			endGame(winner)
		}

		if reflect.DeepEqual(diagonalBoard[0], cpuWinningPattern) {
			winner = "cpu"
			endGame(winner)
		}

		if reflect.DeepEqual(reverseDiagonalBoard[0], cpuWinningPattern) {
			winner = "cpu"
			endGame(winner)
		}
	}

	if count == 0 {
		endGame("draw")
	}
}

func endGame(winner string) {

	if winner == "user" {
		fmt.Println(master + winner + " WIN !!")
	} else if winner == "cpu" {
		fmt.Println(master + winner + " WIN !!")
	} else {
		fmt.Println(master + "DRAW !!")
	}
	fmt.Println(master + "game end !!")
	os.Exit(0)
}
