package htmlthings

type Condition struct {
	elements []*Element
}

func NewCondition(ele *Element) *Condition {
	cond := new(Condition)
	cond.elements = append(cond.elements, ele)

	return cond
}

func (cond *Condition) Include(ele *Element) *Condition {
	cond.elements = append(cond.elements, ele)
	return cond
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

func (ele *Element) Attr(key, val string) *Element {
	ele.attrs[key] = val
	return ele
}
