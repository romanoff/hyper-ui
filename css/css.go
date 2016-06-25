package css

import (
	"fmt"
	"github.com/gorilla/css/scanner"
)

func MinifyCss(content []byte, classesMap map[string]string) ([]byte, error) {
	s := scanner.New(string(content))
	for {
		token := s.Next()
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}
		fmt.Println(token)
	}
	return nil, nil
}
