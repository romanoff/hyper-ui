package runtime

import (
	"github.com/romanoff/hyper-ui/template/ast"
)

func Render(tree *ast.Tree) ([]byte, error) {
	return []byte("Hello world"), nil
}
