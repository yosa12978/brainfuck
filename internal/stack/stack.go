package stack

type Braces struct {
	Start int
	End   int
}

func NewBraces(start int, end int) Braces {
	return Braces{start, end}
}

type Stack struct {
	sp    int
	elems []Braces
}

func NewStack(depth int) *Stack {
	return &Stack{
		sp:    0,
		elems: make([]Braces, depth),
	}
}

func (s *Stack) Push(i Braces) {
	s.elems[s.sp] = i
	s.sp++
}

func (s *Stack) Pop() Braces {
	s.sp--
	return s.elems[s.sp]
}

func (s *Stack) Top() Braces {
	return s.elems[s.sp-1]
}

func (s *Stack) IsEmpty() bool {
	return s.sp == 0
}
