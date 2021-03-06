package rules_test

import (
	"testing"

	"github.com/bucketd/go-graphqlparser/graphql"
	"github.com/bucketd/go-graphqlparser/validation"
	"github.com/bucketd/go-graphqlparser/validation/rules"
)

func TestProvidedRequiredArgumentsOnDirectives(t *testing.T) {
	tt := []ruleTestCase{
		{
			msg: "ignores unknown directives",
			query: `
			{
				dog @unknown
			}
			`,
		},
		{
			msg: "with directives of valid types",
			query: `
			{
				dog @include(if: true) {
				  name
				}
				human @skip(if: false) {
				  name
				}
			}
			`,
		},
		{
			msg: "with directive with missing types",
			query: `
			{
				dog @include {
				  name @skip
				}
			}
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("include", "if", "Boolean!", 0, 0)).
				Add(validation.MissingDirectiveArgError("skip", "if", "Boolean!", 0, 0)),
		},
		{
			msg: "missing optional args on directive defined inside SDL",
			query: `
			type Query {
				foo: String @test
			}
			directive @test(arg1: String, arg2: String! = "") on FIELD_DEFINITION
			`,
		},
		{
			msg: "missing arg on directive defined inside SDL",
			query: `
			type Query {
				foo: String @test
			}
			directive @test(arg: String!) on FIELD_DEFINITION
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("test", "arg", "String!", 0, 0)),
		},
		{
			msg: "missing arg on standard directive",
			query: `
			type Query {
				foo: String @include
			}
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("include", "if", "Boolean!", 0, 0)),
		},
		{
			msg: "missing arg on overridden standard directive",
			query: `
			type Query {
				foo: String @deprecated
			  }
			directive @deprecated(reason: String!) on FIELD
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("deprecated", "reason", "String!", 0, 0)),
		},
		{
			msg: "missing arg on directive defined in schema extension",
			schema: mustBuildSchema(nil, []byte(`
			type Query {
				foo: String
			}
			`)),
			query: `
			directive @test(arg: String!) on OBJECT
			extend type Query  @test
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("test", "arg", "String!", 0, 0)),
		},
		{
			msg: "missing arg on directive used in schema extension",
			schema: mustBuildSchema(nil, []byte(`
			directive @test(arg: String!) on OBJECT
			type Query {
				foo: String
			}
			`)),
			query: `
			extend type Query @test
			`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingDirectiveArgError("test", "arg", "String!", 0, 0)),
		},
	}

	sdlRuleTester(t, tt, rules.ProvidedRequiredArgumentsOnDirectives)
}

func TestProvidedRequiredArguments(t *testing.T) {
	tt := []ruleTestCase{
		{
			msg: "ignores unknown arguments",
			query: `
		{
			dog {
			  isHousetrained(unknownArgument: true)
			}
		}
		`,
		},
		{
			msg: "arg on optional arg",
			query: `
		{
			dog {
			  isHousetrained(atOtherHomes: true)
			}
		}
		`,
		},
		{
			msg: "no arg on optional arg",
			query: `
		{
			dog {
			  isHousetrained
			}
		}
		`,
		},
		{
			msg: "no arg on non-null field with default",
			query: `
		{
			complicatedArgs {
			  nonNullFieldWithDefault
			}
		}
		`,
		},
		{
			msg: "multiple args",
			query: `
		{
			complicatedArgs {
			  multipleReqs(req1: 1, req2: 2)
			}
		}
		`,
		},
		{
			msg: "multiple args reverse order",
			query: `
		{
			complicatedArgs {
			  multipleReqs(req2: 2, req1: 1)
			}
		}
		`,
		},
		{
			msg: "no args on multiple optional",
			query: `
		{
			complicatedArgs {
			  multipleReqs(req2: 2, req1: 1)
			}
		}
		`,
		},
		{
			msg: "one arg on multiple optional",
			query: `
		{
			complicatedArgs {
			  multipleOpts(opt1: 1)
			}
		}
		`,
		},
		{
			msg: "second arg on multiple optional",
			query: `
		{
			complicatedArgs {
			  multipleOpts(opt2: 1)
			}
		}
		`,
		},
		{
			msg: "multiple reqs on mixedList",
			query: `
		{
			complicatedArgs {
			  multipleOptAndReq(req1: 3, req2: 4)
			}
		}
		`,
		},
		{
			msg: "multiple reqs and one opt on mixedList",
			query: `
		{
			complicatedArgs {
			  multipleOptAndReq(req1: 3, req2: 4, opt1: 5)
			}
		}
		`,
		},
		{
			msg: "all reqs and opts on mixedList",
			query: `
		{
			complicatedArgs {
			  multipleOptAndReq(req1: 3, req2: 4, opt1: 5, opt2: 6)
			}
		}
		`,
		},
		{
			msg: "missing one non-nullable argument",
			query: `
		{
			complicatedArgs {
			  multipleReqs(req2: 2)
			}
		}
		`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingFieldArgError("multipleReqs", "req1", "Int!", 0, 0)),
		},
		{
			msg: "missing multiple non-nullable arguments",
			query: `
		{
			complicatedArgs {
			  multipleReqs
			}
		}
		`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingFieldArgError("multipleReqs", "req1", "Int!", 0, 0)).
				Add(validation.MissingFieldArgError("multipleReqs", "req2", "Int!", 0, 0)),
		},
		{
			msg: "incorrect value and missing argument",
			query: `
		{
			complicatedArgs {
			  multipleReqs(req1: "one")
			}
		}
		`,
			errs: (*graphql.Errors)(nil).
				Add(validation.MissingFieldArgError("multipleReqs", "req2", "Int!", 0, 0)),
		},
	}

	_ = tt
	// queryRuleTester(t, tt, rules.ProvidedRequiredArguments)
}
