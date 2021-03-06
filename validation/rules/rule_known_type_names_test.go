package rules_test

import (
	"testing"

	"github.com/bucketd/go-graphqlparser/graphql"
	"github.com/bucketd/go-graphqlparser/validation"
	"github.com/bucketd/go-graphqlparser/validation/rules"
)

func BenchmarkKnownTypeNames(b *testing.B) {
	query := `
		type B

		type SomeObject implements C {
			e(d: D): E
		}

		union SomeUnion = F | G

		interface SomeInterface {
			i(h: H): I
		}

		input SomeInput {
			j: J
		}

		directive @SomeDirective(k: K) on QUERY

		schema {
			query: L
			mutation: M
			subscription: N
		}
	`

	queryRuleBencher(b, ruleTestCase{query: query}, rules.KnownTypeNames)
}

func TestKnownTypeNames(t *testing.T) {
	t.Run("query document", func(t *testing.T) {
		tt := []ruleTestCase{
			{
				msg: "known type names are valid",
				query: `
					query Foo($var: String, $required: [String!]!) {
						user(id: 4) {
							pets { ... on Pet { name }, ...PetFields, ... { name } }
						}
					}

					fragment PetFields on Pet {
						name
					}
				`,
			},
			{
				msg: "unknown type names are invalid",
				query: `
					query Foo($var: JumbledUpLetters) {
						user(id: 4) {
							name
							pets { ... on Badger { name }, ...PetFields }
						}
					}

					fragment PetFields on Peettt {
						name
					}
				`,
				errs: (*graphql.Errors)(nil).
					Add(validation.UnknownTypeError("JumbledUpLetters", 0, 0)).
					Add(validation.UnknownTypeError("Badger", 0, 0)).
					Add(validation.UnknownTypeError("Peettt", 0, 0)),
			},
			// NOTE: It's not possible to use our parser and have a schema without the built-in
			// scalar types included. It's part of the spec, and a server would be pretty useless
			// without them, so we're not going to include this test, but for reference, it's name
			// is included below:
			//{
			//	msg: "references to standard scalars that are missing in schema",
			//},
		}

		queryRuleTester(t, tt, rules.KnownTypeNames)
	})

	t.Run("sdl document", func(t *testing.T) {
		tt := []ruleTestCase{
			{
				msg: "use standard scalars",
				query: `
					type Query {
						string: String
						int: Int
						float: Float
						boolean: Boolean
						id: ID
					}
				`,
			},
			{
				msg: "reference types defined inside the same document",
				query: `
					union SomeUnion = SomeObject | AnotherObject

					type SomeObject implements SomeInterface {
						someScalar(arg: SomeInputObject): SomeScalar
					}

					type AnotherObject {
						foo(arg: SomeInputObject): String
					}

					type SomeInterface {
						someScalar(arg: SomeInputObject): SomeScalar
					}

					input SomeInputObject {
						someScalar: SomeScalar
					}

					scalar SomeScalar

					type RootQuery {
						someInterface: SomeInterface
						someUnion: SomeUnion
						someScalar: SomeScalar
						someObject: SomeObject
					}

					schema {
						query: RootQuery
					}
				`,
			},
			{
				msg: "unknown type references",
				query: `
					type A
					type B

					type SomeObject implements C {
						e(d: D): E
					}

					union SomeUnion = F | G

					interface SomeInterface {
						i(h: H): I
					}

					input SomeInput {
						j: J
					}

					directive @SomeDirective(k: K) on QUERY

					schema {
						query: L
						mutation: M
						subscription: N
					}
				`,
				errs: (*graphql.Errors)(nil).
					Add(validation.UnknownTypeError("C", 0, 0)).
					Add(validation.UnknownTypeError("D", 0, 0)).
					Add(validation.UnknownTypeError("E", 0, 0)).
					Add(validation.UnknownTypeError("F", 0, 0)).
					Add(validation.UnknownTypeError("G", 0, 0)).
					Add(validation.UnknownTypeError("H", 0, 0)).
					Add(validation.UnknownTypeError("I", 0, 0)).
					Add(validation.UnknownTypeError("J", 0, 0)).
					Add(validation.UnknownTypeError("K", 0, 0)).
					Add(validation.UnknownTypeError("L", 0, 0)).
					Add(validation.UnknownTypeError("M", 0, 0)).
					Add(validation.UnknownTypeError("N", 0, 0)),
			},
			{
				msg: "doesn't consider non-type definitions",
				query: `
					query Foo { __typename }
					fragment Foo on Query { __typename }
					directive @Foo on QUERY

					type Query {
						foo: Foo
					}
				`,
				errs: (*graphql.Errors)(nil).
					Add(validation.UnknownTypeError("Foo", 0, 0)),
			},
			{
				msg:    "reference standard scalars inside extension document",
				schema: mustBuildSchema(nil, []byte(`type Foo`)),
				query: `
					type SomeType {
						string: String
						int: Int
						float: Float
						boolean: Boolean
						id: ID
					}
				`,
			},
			{
				msg:    "reference types inside extension document",
				schema: mustBuildSchema(nil, []byte(`type Foo`)),
				query: `
					type QueryRoot {
						foo: Foo
						bar: Bar
					}

					scalar Bar

					type NotSchema {
						query: QueryRoot
					}
				`,
			},
			{
				msg:    "unknown type references inside extension document",
				schema: mustBuildSchema(nil, []byte(`type A`)),
				query: `
					type B

					type SomeObject implements C {
						e(d: D): E
					}

					union SomeUnion = F | G

					interface SomeInterface {
						i(h: H): I
					}

					input SomeInput {
						j: J
					}

					directive @SomeDirective(k: K) on QUERY

					type NotSchema {
						query: L
						mutation: M
						subscription: N
					}
				`,
				errs: (*graphql.Errors)(nil).
					Add(validation.UnknownTypeError("C", 0, 0)).
					Add(validation.UnknownTypeError("D", 0, 0)).
					Add(validation.UnknownTypeError("E", 0, 0)).
					Add(validation.UnknownTypeError("F", 0, 0)).
					Add(validation.UnknownTypeError("G", 0, 0)).
					Add(validation.UnknownTypeError("H", 0, 0)).
					Add(validation.UnknownTypeError("I", 0, 0)).
					Add(validation.UnknownTypeError("J", 0, 0)).
					Add(validation.UnknownTypeError("K", 0, 0)).
					Add(validation.UnknownTypeError("L", 0, 0)).
					Add(validation.UnknownTypeError("M", 0, 0)).
					Add(validation.UnknownTypeError("N", 0, 0)),
			},
		}

		sdlRuleTester(t, tt, rules.KnownTypeNames)
	})
}
