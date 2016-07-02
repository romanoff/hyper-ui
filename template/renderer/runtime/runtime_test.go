package runtime

import (
	"github.com/romanoff/hyper-ui/template/ast"
	"testing"
)

func TestRenderText(t *testing.T) {
	tr := &ast.Tree{}
	_, err := Render(tr)
	if err != nil {
		t.Errorf("Expected not to get error while rendering text node, but got %v", err)
	}
}
