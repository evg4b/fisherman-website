// nolint: dupl
package rules_test

import (
	"fisherman/internal/rules"
	"fisherman/internal/validation"
	"fisherman/testing/mocks"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitMessage_NotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		notEmpty bool
		hasError bool
	}{
		{name: "Active with empty string", hasError: true, message: "", notEmpty: true},
		{name: "Inactive with empty string", hasError: false, message: "", notEmpty: false},
		{name: "Active with not empty string", hasError: false, message: "message", notEmpty: true},
		{name: "Active with not empty string", hasError: false, message: "message", notEmpty: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mocks.NewExecutionContextMock(t).MessageMock.Return(tt.message, nil)

			rule := rules.CommitMessage{
				BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
				NotEmpty: tt.notEmpty,
			}

			err := rule.Check(ctx, ioutil.Discard)

			if tt.hasError {
				assert.EqualError(t, err, errorMessage(rules.CommitMessageType, "commit message should not be empty"))
				assert.IsType(t, &validation.Error{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommitMessage_HasPrefix(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		prefix   string
		hasError bool
	}{
		{name: "Active with empty string", hasError: true, message: "", prefix: "prefix-"},
		{name: "Inactive with empty string", hasError: false, message: "", prefix: ""},
		{name: "Active with string and prefix", hasError: false, message: "prefix-message", prefix: "prefix-"},
		{name: "Inactive with string and prefix", hasError: false, message: "prefix-message", prefix: ""},
		{name: "Active with string and other prefix", hasError: true, message: "other-prefix-message", prefix: "prefix-"},
		{name: "Inactive with string and other prefix", hasError: false, message: "other-prefix-message", prefix: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mocks.NewExecutionContextMock(t).MessageMock.Return(tt.message, nil)

			rule := rules.CommitMessage{
				BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
				Prefix:   tt.prefix,
			}

			err := rule.Check(ctx, ioutil.Discard)

			if tt.hasError {
				assert.EqualError(t, err, errorMessage(rules.CommitMessageType, "commit message should have prefix 'prefix-'"))
				assert.IsType(t, &validation.Error{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommitMessage_HasSuffix(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		suffix   string
		hasError bool
	}{
		{name: "Active with empty string", hasError: true, message: "", suffix: "-suffix"},
		{name: "Inactive with empty string", hasError: false, message: "", suffix: ""},
		{name: "Active with string and suffix", hasError: false, message: "message-suffix", suffix: "-suffix"},
		{name: "Inactive with string and suffix", hasError: false, message: "message-suffix", suffix: ""},
		{name: "Active with string and other suffix", hasError: true, message: "message-suffix-other", suffix: "-suffix"},
		{name: "Inactive with string and other suffix", hasError: false, message: "message-suffix-other", suffix: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mocks.NewExecutionContextMock(t).MessageMock.Return(tt.message, nil)

			rule := rules.CommitMessage{
				BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
				Suffix:   tt.suffix,
			}

			err := rule.Check(ctx, ioutil.Discard)

			if tt.hasError {
				assert.EqualError(t, err, errorMessage(rules.CommitMessageType, "commit message should have suffix '-suffix'"))
				assert.IsType(t, &validation.Error{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommitMessage_TestMessageRegexp(t *testing.T) {
	tests := []struct {
		name       string
		message    string
		expression string
		hasError   bool
	}{
		{name: "Inactive with empty string", hasError: false, message: "", expression: ""},
		{name: "Active with empty string", hasError: true, message: "", expression: "^[a-z]{5}$"},
		{name: "Active with correct matching", hasError: true, message: "message", expression: "^[a-z]{5}$"},
		{name: "Active with correct matching", hasError: false, message: "message", expression: "^[a-z]{7}$"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mocks.NewExecutionContextMock(t).MessageMock.Return(tt.message, nil)

			rule := rules.CommitMessage{
				BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
				Regexp:   tt.expression,
			}

			err := rule.Check(ctx, ioutil.Discard)

			if tt.hasError {
				assert.Error(t, err)
				assert.IsType(t, &validation.Error{}, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCommitMessage_TestMessageRegexp_MachingError(t *testing.T) {
	ctx := mocks.NewExecutionContextMock(t).MessageMock.Return("message", nil)

	rule := rules.CommitMessage{
		BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
		Regexp:   "[a-z]($",
	}

	err := rule.Check(ctx, ioutil.Discard)

	assert.Error(t, err)
}

func TestCommitMessage_Compile(t *testing.T) {
	rule := rules.CommitMessage{
		BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
		Prefix:   "Prefix{{var1}}",
		Suffix:   "Suffix{{var1}}",
		Regexp:   "Regexp{{var1}}",
		NotEmpty: true,
	}

	rule.Compile(map[string]interface{}{"var1": "VALUE"})

	assert.Equal(t, rules.CommitMessage{
		BaseRule: rules.BaseRule{Type: rules.CommitMessageType},
		Prefix:   "PrefixVALUE",
		Suffix:   "SuffixVALUE",
		Regexp:   "RegexpVALUE",
		NotEmpty: true,
	}, rule)
}
