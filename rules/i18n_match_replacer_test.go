package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringAddr(s string) *string {
	return &s
}

func TestGetJavaMatchReplacer(t *testing.T) {
	testCases := []struct {
		match    []string
		expected *string
	}{
		{match: []string{`StringProvider.getValue("key1", "default value 1")`, "key1", "value 1"}, expected: stringAddr("StringValues.INSTANCE.getKey1()")},
		{match: []string{`StringProvider.getValue("key4","default value 4")`, "key4", "value 4"}, expected: stringAddr("StringValues.INSTANCE.getKey4()")},
		{match: []string{``, "", ""}, expected: nil},
	}

	replacer := i18nMatchReplacer{}
	for _, testCase := range testCases {
		result := replacer.GetJavaMatchReplacer(testCase.match)
		assert.Equal(t, testCase.expected, result)
	}
}
