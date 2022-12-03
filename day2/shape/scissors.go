package shape

type Scissors struct {
}

func (scissors *Scissors) Name() ShapeName {
	return EnumScissors
}

func (scissors *Scissors) Value() int {
	return 3
}

func (scissors *Scissors) WithdrawsWith() Shape {
	return &Scissors{}
}

func (scissors *Scissors) WinsFrom() Shape {
	return &Paper{}
}

func (scissors *Scissors) LoosesTo() Shape {
	return &Rock{}
}
