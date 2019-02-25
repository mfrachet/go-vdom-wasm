package vn

func H(tagName string, attrs *Attrs, variadicChildren interface{}, params ...interface{}) Node {
	sanitizedAttrs := Sanitize(attrs)
	var key KeyIdentifier
	var children Children
	var textNode *TextNode

	if len(params) == 1 {
		key = params[0].(KeyIdentifier)
	}

	switch (variadicChildren).(type) {
	case string:
		textNode = NewTextnode((variadicChildren).(string))
	default:
		children = variadicChildren.(Children)
	}

	return NewNode(tagName, sanitizedAttrs, children, textNode, nil, key)
}
