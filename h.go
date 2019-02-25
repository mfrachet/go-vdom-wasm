package vn

func H(tagName string, attrs *Attrs, children interface{}, params ...string) Node {
	sanitizedAttrs := Sanitize(attrs)
	var key *string

	if len(params) == 1 {
		key = &params[0]
	}

	switch (children).(type) {
	case string:
		return NewNode(tagName, sanitizedAttrs, nil, NewTextnode((children).(string)), nil, key)
	default:
		return NewNode(tagName, sanitizedAttrs, children.(Children), nil, nil, key)
	}
}
