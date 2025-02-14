package validator

import (
	"encoding/json"
	"sync"
)

type Validator struct {
	errs map[string]string
	mu   sync.Mutex
}

func New() *Validator {
	return &Validator{
		errs: make(map[string]string),
		mu:   sync.Mutex{},
	}
}

func (v *Validator) Valid() bool {
	return len(v.errs) == 0
}

func (v *Validator) AddError(key, message string) {
	v.mu.Lock()
	if _, exists := v.errs[key]; !exists {
		v.errs[key] = message
	}
	v.mu.Unlock()
}

// Error to match error interface, so that I can return the validation error through grpc.
func (v *Validator) Error() string {
	byt, _ := json.Marshal(v.errs)
	return string(byt)
}
