package shape

type Paper struct {
}

func (paper *Paper) Name() ShapeName {
	return EnumPaper
}

func (paper *Paper) Value() int {
	return 2
}

func (paper *Paper) WithdrawsWith() Shape {
	return &Paper{}
}

func (paper *Paper) WinsFrom() Shape {
	return &Rock{}
}

func (paper *Paper) LoosesTo() Shape {
	return &Scissors{}
}
