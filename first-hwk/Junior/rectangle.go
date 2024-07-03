package Junior

type Rectangle struct {
	Length, Width int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}
