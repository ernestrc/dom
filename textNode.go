package dom

import (
	"github.com/ernestrc/dom/js"
)

func AsTextNode(v js.Value) *TextNode {
	if !v.Valid() {
		return nil
	}
	return &TextNode{NodeBase{v: v}}
}

type TextNode struct {
	NodeBase
}