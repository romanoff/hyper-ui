package css

import (
	"errors"
	"fmt"
	"github.com/aymerick/douceur/parser"
	"github.com/gorilla/css/scanner"
)

func MustMinifyCss(content string, classesMap map[string]string) (string, error) {
	return minifyCss(content, classesMap, true)
}

func MinifyCss(content string, classesMap map[string]string) (string, error) {
	return minifyCss(content, classesMap, false)
}

func minifyCss(content string, classesMap map[string]string, errorIfClassMissing bool) (string, error) {
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
				if errorIfClassMissing {
					return "", errors.New(fmt.Sprintf("unused class: .%v", token.Value))
				}
				selectors = append(selectors, selector)
				continue
			}
			selectors = append(selectors, "."+classesMap[token.Value])
		}
		rule.Selectors = selectors
	}
	return stylesheet.String(), nil
}
