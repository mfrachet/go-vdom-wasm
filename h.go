package vn

func H(tagName string, attrs *Attrs, children interface{}) *Vnode {
	switch (children).(type) {
	case string:
		return &Vnode{tagName, attrs, &TextNode{(children).(string), nil}, nil}
	default:
		return &Vnode{tagName, attrs, children.(Children), nil}
	}
}
