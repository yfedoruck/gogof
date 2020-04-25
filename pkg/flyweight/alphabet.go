package flyweight

type Letter interface {
	Construct(s string)
	GetSign() string
}

type Mark struct {
	Sign string
}

func (r *Mark) Construct(s string) {
	r.Sign = s
}

func (r Mark) GetSign() string {
	return r.Sign
}

type A struct {
	Mark
}
type B struct {
	Mark
}
type C struct {
	Mark
}

type Alphabet struct {
	Data map[string]Letter
}

func NewAlphabet() Alphabet {
	ab := Alphabet{}
	ab.Data = make(map[string]Letter)
	return ab
}

func (r *Alphabet) GetLetter(l string) Letter {
	letter, ok := r.Data[l]
	if ok {
		return letter
	}
	letter = r.build(l)
	r.Data[l] = letter
	return letter
}

func (r Alphabet) build(l string) Letter {
	var letter Letter
	switch l {
	case "a":
		letter = &A{}
	case "b":
		letter = &B{}
	case "c":
		letter = &C{}
	default:
		letter = &A{}
	}
	letter.Construct(l)

	return letter
}
