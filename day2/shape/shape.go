package shape

import "errors"

type ShapeName string

const (
	EnumRock     ShapeName = "Rock"
	EnumPaper    ShapeName = "Paper"
	EnumScissors ShapeName = "Scissors"
)

type Shape interface {
	Name() ShapeName
	Value() int
	WithdrawsWith() Shape
	WinsFrom() Shape
	LoosesTo() Shape
}

func ShapeFromCharacter(character string) Shape {
	switch character {
	case "A":
		fallthrough
	case "X":
		return &Rock{}
	case "B":
		fallthrough
	case "Y":
		return &Paper{}
	case "C":
		fallthrough
	case "Z":
		return &Scissors{}
	default:
		panic(errors.New("this character isn't valid"))
	}

}
