package gorenderer

import (
	"github.com/romanoff/hyper-ui/template/ast"
	"testing"
)

func TestRenderText(t *testing.T) {
	tr := &ast.Tree{
		Root: &ast.ListNode{Nodes: []ast.Node{
			&ast.TextNode{
				NodeType: ast.NodeText,
				Content:  []byte("Hello world"),
			},
		}},
	}
	renderer := &Renderer{Tree: tr}
	renderer.Init()
	content, err := renderer.Render()
	if err != nil {
		t.Errorf("Expected not to get error while rendering text node, but got %v", err)
	}
	expected := `buffer := new(bytes.Buffer)
buffer.Write([]byte{'H','e','l','l','o',' ','w','o','r','l','d'})
return buffer.Bytes(), nil
`
	if expected != string(content) {
		t.Errorf("Expected to get:\n%v\n,but got:\n%v\n", expected, string(content))
	}
}
