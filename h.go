package vn

func H(tagName string, attrs *Attrs, params ...interface{}) Node {
	sanitizedAttrs := Sanitize(attrs)
	var key KeyIdentifier
	var children Children
	var textNode *TextNode

	for _, param := range params {
		switch (param).(type) {
		case string:
			textNode = NewTextnode((param).(string))
		case KeyIdentifier:
			key = param.(KeyIdentifier)
		default:
			children = param.(Children)
		}
	}

	return NewNode(tagName, sanitizedAttrs, children, textNode, nil, key)
}
