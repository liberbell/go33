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

func (d *Data) RemoveElement(element int) {
	index := -1
	for i, e := range d.Elements {
		if e = element {
			index = i
			break
		}
	}
}