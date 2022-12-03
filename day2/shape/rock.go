package shape

type Rock struct {
}

func (rock *Rock) Name() ShapeName {
	return EnumRock
}

func (rock *Rock) Value() int {
	return 1
}

func (rock *Rock) WithdrawsWith() Shape {
	return &Rock{}
}

func (rock *Rock) WinsFrom() Shape {
	return &Scissors{}
}

func (rock *Rock) LoosesTo() Shape {
	return &Paper{}
}
