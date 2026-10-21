package simpleGoApp

type Data struct {
	Elements []int
}

func NewData() *Data {
	return &Data{
		Elements: make([]int, 0),
	}
}
