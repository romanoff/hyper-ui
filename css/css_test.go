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
	minifiedConent, err := MinifyCss(cssContent, classesMap)
	if err != nil {
		t.Errorf("Expected to get no error, but got %v\n", err)
	}
	expectedContent := `.a {
  background-color: red;
}`
	if minifiedConent != expectedContent {
		t.Errorf("Expected to get:\n %v\n, but got: %v\n", expectedContent, minifiedConent)
	}
	classesMap = map[string]string{
		"unknown": "a",
	}
	_, err = MustMinifyCss(cssContent, classesMap)
	if err == nil {
		t.Errorf("Expected to get error due to missing css class, but got nil")
	}
	if err.Error() != "unused class: .button" {
		t.Errorf("Expected to get error message:\n%v\n,but got:\n%v\n", "unused class: .button", err.Error())
	}
}
