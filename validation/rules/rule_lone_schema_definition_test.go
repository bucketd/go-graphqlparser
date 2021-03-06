package rules_test

import (
	"testing"

	"github.com/bucketd/go-graphqlparser/graphql"
	"github.com/bucketd/go-graphqlparser/validation"
)

func TestLoneSchemaDefinition(t *testing.T) {
	tt := []ruleTestCase{
		{
			msg: "no schema",
			query: `
				type Foo { checkEnumValueUniqueness: String }
			`,
			errs: nil,
		},
		{
			msg: "one schema definition",
			query: `
				schema { query: Foo }
				type Foo { checkEnumValueUniqueness: String }
			`,
			errs: nil,
		},
		{
			msg: "multiple schema definitions",
			query: `
				schema { query: Foo }
				type Foo { checkEnumValueUniqueness: String }
				schema { mutation: Foo }
				schema { subscription: Foo }
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.SchemaDefinitionNotAloneError(0, 0)).
				Add(validation.SchemaDefinitionNotAloneError(0, 0)),
		},
		{
			msg:    "define schema in schema extension",
			schema: &graphql.Schema{},
			query: `
				schema {
					query: Foo
				}
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.CanNotDefineSchemaWithinExtensionError(0, 0)),
		},
		{
			msg: "redefine schema in schema extension",
			schema: mustBuildSchema(nil, []byte(`
				schema {
					query: Foo
				}
		
				type Foo
			`)),
			query: `
				schema {
					mutation: Foo
				}
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.CanNotDefineSchemaWithinExtensionError(0, 0)),
		},
		{
			msg: "redefine implicit schema in schema extension",
			schema: mustBuildSchema(nil, []byte(`
				type Query {
					fooField: Foo
				}
		
				type Foo {
					foo: String
				}
			`)),
			query: `
				schema {
					mutation: Foo
				}
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.CanNotDefineSchemaWithinExtensionError(0, 0)),
		},
		{
			msg: "extend schema in schema extension",
			schema: mustBuildSchema(nil, []byte(`
				schema {
					query: Foo
				}
		
				type Foo
			`)),
			query: `
				extend schema {
					mutation: Foo
				}
			`,
			errs: nil,
		},
	}

	sdlRuleTester(t, tt, func(w *validation.Walker) {})
}
