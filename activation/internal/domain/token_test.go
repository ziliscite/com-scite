package domain

import (
	"errors"
	"testing"
)

func TestValidateTokenPlaintext(t *testing.T) {
	tests := []struct {
		tokenPlaintext string
		expectedError  error
	}{
		{
			tokenPlaintext: "",
			expectedError:  errors.New("must be provided"),
		},
		{
			tokenPlaintext: "shorttoken",
			expectedError:  errors.New("must be 26 bytes long"),
		},
		{
			tokenPlaintext: "thisisaverylongtokenthatexceedsthelength",
			expectedError:  errors.New("must be 26 bytes long"),
		},
		{
			tokenPlaintext: "XFNYUZ4CKGLOPYJYFZJZI7QQJE",
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.tokenPlaintext, func(t *testing.T) {
			err := ValidateTokenPlaintext(tt.tokenPlaintext)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
			}
		})
	}
}
