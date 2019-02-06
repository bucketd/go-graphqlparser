package rules

import (
	"testing"

	"github.com/bucketd/go-graphqlparser/graphql"
	"github.com/bucketd/go-graphqlparser/language"
	"github.com/bucketd/go-graphqlparser/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// schemaDocument ...
	schemaDocument = `
		schema {
			query: Query
		}

		type Query {
			foo: String!
		}
	`
)

// ruleTestCase ...
type ruleTestCase struct {
	msg   string
	query string
	errs  *graphql.Errors
}

// queryRuleTester ...
func queryRuleTester(t *testing.T, tt []ruleTestCase, fn validation.VisitFunc) {
	var schema graphql.Schema

	schemaParser := language.NewParser([]byte(schemaDocument))

	schemaDoc, err := schemaParser.Parse()
	require.NoError(t, err, "failed to parse schema document")

	errs := validation.ValidateSDL(schemaDoc, &schema)
	require.True(t, errs.Len() == 0, "found errors validation schema")

	for _, tc := range tt {
		parser := language.NewParser([]byte(tc.query))

		doc, err := parser.Parse()
		if err != nil {
			require.NoError(t, err)
		}

		walker := validation.NewWalker([]validation.VisitFunc{fn})

		errs := validation.ValidateWithWalker(doc, &schema, walker)
		assert.Equal(t, tc.errs, errs, tc.msg)
	}
}