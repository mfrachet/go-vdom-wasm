package vn

type Children []*Vnode

func (children Children) createElement() {
	for _, childNode := range children {
		childNode.createElement()
	}
}
