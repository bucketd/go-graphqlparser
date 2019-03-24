package rules_test

import (
	"testing"

	"github.com/bucketd/go-graphqlparser/validation/rules"
)

func TestNoUndefinedVariables(t *testing.T) {
	tt := []ruleTestCase{}

	queryRuleTester(t, tt, rules.NoUndefinedVariables)
}
