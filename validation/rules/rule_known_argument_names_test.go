package rules

import (
	"testing"
)

func TestKnownArgumentNames(t *testing.T) {
	tt := []ruleTestCase{}

	queryRuleTester(t, tt, knownArgumentNames)
}