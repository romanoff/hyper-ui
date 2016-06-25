package css

import (
	"github.com/aymerick/douceur/parser"
	"github.com/gorilla/css/scanner"
)

func MinifyCss(content string, classesMap map[string]string) (string, error) {
	stylesheet, err := parser.Parse(string(content))
	if err != nil {
		return "", err
	}
	for _, rule := range stylesheet.Rules {
		selectors := make([]string, 0, 0)
		for _, selector := range rule.Selectors {
			s := scanner.New(selector)
			token := s.Next()
			if token.Type != scanner.TokenChar || token.Value != "." {
				selectors = append(selectors, selector)
				continue
			}
			token = s.Next()
			if token.Type != scanner.TokenIdent {
				selectors = append(selectors, selector)
				continue
			}
			if classesMap[token.Value] == "" {
				continue
			}
			selectors = append(selectors, "."+classesMap[token.Value])
		}
		rule.Selectors = selectors
	}
	return stylesheet.String(), nil
}
