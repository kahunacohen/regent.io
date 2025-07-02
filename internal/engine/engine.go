package engine

import "fmt"

type Policy struct {
	Name    string
	Content []byte
}

type Evaluator struct {
	policies []Policy
}

// Compiles rego policies here, create prepared querries.
func NewEvaluator(policies []Policy) *Evaluator {
	return &Evaluator{
		policies: policies,
	}
}

// Actually evaluate a policy
func (e *Evaluator) EvaluatePolicy(name string) {
	for _, policy := range e.policies {
		fmt.Println(policy.Name)
	}
}
