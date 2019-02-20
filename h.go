package vn

func H(tagName string, attrs *Attrs, children interface{}) *Vnode {
	sanitizedAttrs := Sanitize(attrs)

	switch (children).(type) {
	case string:
		return NewVNode(tagName, sanitizedAttrs, nil, &TextNode{(children).(string), nil}, nil)
	default:
		return NewVNode(tagName, sanitizedAttrs, children.(Children), nil, nil)
	}
}
