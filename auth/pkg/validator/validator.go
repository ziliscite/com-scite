package validator

import "sync"

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

func (v *Validator) Errors() map[string]string {
	return v.errs
}
