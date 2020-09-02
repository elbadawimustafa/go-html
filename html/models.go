package html


type Page struct {
	elements []*Element
}

type Properties map[string]interface{}

type Element struct{
	tag string
	priority int8
	properties Properties
	content string
	children []*Element
}


func (e *Element) AddChild(t, c string, p int8) *Element{
	ch := &Element{
		tag: t,
		priority: p,
		content: c,
	}

	e.children = append(e.children, ch)
	return ch
}

func (e *Element) SetProperty(k string, v interface{}){
	e.properties[k] = v
}

