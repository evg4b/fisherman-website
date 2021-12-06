package guards_test

import (
	. "fisherman/pkg/guards"
	"testing"

	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
)

func TestShouldBeDefined(t *testing.T) {
	tests := []struct {
		name    string
		object  interface{}
		message string
		err     string
	}{
		{
			name:    "should not panic for zero",
			message: "unknown",
			object:  0,
		},
		{
			name:    "should not panic for defined empty string",
			message: "unknown",
			object:  "",
		},
		{
			name:    "should not panic for defined empty struct",
			message: "unknown",
			object:  struct{}{},
		},
		{
			name:    "should not panic for defined empty slice",
			message: "unknown",
			object:  []string{},
		},
		{
			name:    "should panic for nil value",
			message: "value is null",
			err:     "value is null",
			object:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.err) > 0 {
				assert.PanicsWithError(t, tt.err, func() {
					ShouldBeDefined(tt.object, tt.message)
				})
			} else {
				assert.NotPanics(t, func() {
					ShouldBeDefined(tt.object, tt.message)
				})
			}
		})
	}
}

func TestNoError(t *testing.T) {
	t.Run("panic for error", func(t *testing.T) {
		err := errors.New("Test err")
		assert.PanicsWithError(t, err.Error(), func() {
			NoError(err)
		})
	})
	t.Run("does not panic for nil", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NoError(nil)
		})
	})
}
