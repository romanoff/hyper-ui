package css

import (
	"testing"
)

func TestMinifyCss(t *testing.T) {
	cssContent := `
.button {
  background-color: red;
}
`
	classesMap := map[string]string{
		"button": "a",
	}
	MinifyCss([]byte(cssContent), classesMap)
}
