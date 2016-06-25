package css

import (
	// "fmt"
	"github.com/aymerick/douceur/parser"
)

func MinifyCss(content string, classesMap map[string]string) (string, error) {
	stylesheet, err := parser.Parse(string(content))
	if err != nil {
		return "", err
	}
	for _, rule := range stylesheet.Rules {
		rule.Selectors = []string{".a"}
	}
	return stylesheet.String(), nil
}
