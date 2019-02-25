package vn

func H(tagName string, params ...interface{}) Node {
	var key KeyIdentifier
	var children Children
	var textNode *TextNode
	attrs := &Attrs{}

	for _, param := range params {
		switch (param).(type) {
		case string:
			textNode = NewTextnode((param).(string))
		case KeyIdentifier:
			key = param.(KeyIdentifier)
		case *Props:
			attrs.Props = param.(*Props)
		case *Ev:
			attrs.Events = param.(*Ev)
		default:
			children = param.(Children)
		}
	}

	sanitizedAttrs := Sanitize(attrs)

	return NewNode(tagName, sanitizedAttrs, children, textNode, nil, key)
}
