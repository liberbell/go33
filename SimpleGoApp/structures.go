package simpleGoApp

type Data struct {
	Elements []int
}

func NewData() *Data {
	return &Data{
		Elements: make([]int, 0),
	}
}

func (d *Data) AddElement(element int) {
	d.Elements = append(d.Elements, element)
}
