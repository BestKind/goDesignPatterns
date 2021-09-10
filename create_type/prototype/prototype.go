package prototype

type Example struct {
	Content string
}

func (e *Example) Clone() *Example {
	res:= *e
	return &res
}
