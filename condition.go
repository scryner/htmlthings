package htmlthings

type Condition struct {
	elements []*Element
}

func NewCondition(ele *Element) *Condition {
	cond := new(Condition)
	cond.elements = append(cond.elements, ele)

	return cond
}

func (cond *Condition) AddSubElement(ele *Element) {
	cond.elements = append(cond.elements, ele)
}

type Element struct {
	name  string
	attrs map[string]string
}

func NewElement(name string) *Element {
	return &Element{
		name:  name,
		attrs: make(map[string]string),
	}
}

func (ele *Element) AddAttribute(key, val string) {
	ele.attrs[key] = val
}
