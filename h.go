package vn

func H(tagName string, attrs *Attrs, children interface{}) *Vnode {
	switch (children).(type) {
	case string:
		return NewVNode(tagName, attrs, nil, &TextNode{(children).(string), nil}, nil)
	default:
		return NewVNode(tagName, attrs, children.(Children), nil, nil)
	}
}
