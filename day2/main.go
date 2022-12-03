package main

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/AndersonBargas/aoc2022/shape"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sumOfAllMatchesPartOne := 0
	sumOfAllMatchesPartTwo := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		characters := strings.Split(line, " ")
		opponentShape, playerShape := transformCharactersIntoShapes(characters[0], characters[1])
		sumOfAllMatchesPartOne += matchForThePartOne(opponentShape, playerShape)
		sumOfAllMatchesPartTwo += matchForThePartTwo(opponentShape, characters[1])
	}

	readFile.Close()

	println(sumOfAllMatchesPartOne)
	println(sumOfAllMatchesPartTwo)
}

func transformCharactersIntoShapes(opponentCharacter, playerCharacter string) (shape.Shape, shape.Shape) {
	opponentShape := shape.ShapeFromCharacter(opponentCharacter)
	playerShape := shape.ShapeFromCharacter(playerCharacter)
	return opponentShape, playerShape
}

func matchForThePartOne(opponentShape, playerShape shape.Shape) int {
	playerScore := 0

	playerScore += playerShape.Value()

	if opponentShape.Name() == playerShape.Name() {
		playerScore += 3
	}

	if opponentShape.Name() == shape.EnumRock && playerShape.Name() == shape.EnumPaper {
		playerScore += 6
	}

	if opponentShape.Name() == shape.EnumPaper && playerShape.Name() == shape.EnumScissors {
		playerScore += 6
	}

	if opponentShape.Name() == shape.EnumScissors && playerShape.Name() == shape.EnumRock {
		playerScore += 6
	}

	return playerScore
}

func matchForThePartTwo(opponentShape shape.Shape, playerStrategy string) int {
	switch playerStrategy {
	case "X": // player looses
		playerShape := opponentShape.WinsFrom()
		return matchForThePartOne(opponentShape, playerShape)
	case "Y": // player withdraw
		playerShape := opponentShape.WithdrawsWith()
		return matchForThePartOne(opponentShape, playerShape)
	case "Z": // player wins
		playerShape := opponentShape.LoosesTo()
		return matchForThePartOne(opponentShape, playerShape)
	default:
		panic(errors.New("this player strategy ins't valid"))
	}
}
