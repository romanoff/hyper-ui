package runtime

import (
	"github.com/romanoff/hyper-ui/template/ast"
	"testing"
)

func TestRenderText(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.TextNode{Content: []byte("Hello world")},
		}},
	}
	content, err := Render(tr)
	if err != nil {
		t.Errorf("Expected not to get error while rendering text node, but got %v", err)
	}
	expected := "Hello world"
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}
