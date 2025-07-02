package engine

import (
	"testing"
)

func TestEngine(t *testing.T) {
	evaluator := NewEvaluator([]Policy{
		{
			Name: "policy1",
			Content: []byte(`
			package example
			allow = true	
		`),
		},
	})

	evaluator.EvaluatePolicy("policy1")
}
