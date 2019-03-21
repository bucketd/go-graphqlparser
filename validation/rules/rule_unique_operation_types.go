package rules

import (
	"github.com/bucketd/go-graphqlparser/ast"
	"github.com/bucketd/go-graphqlparser/graphql"
	"github.com/bucketd/go-graphqlparser/validation"
)

// uniqueOperationTypes ...
func uniqueOperationTypes(w *validation.Walker) {
	w.AddSchemaDefinitionLeaveEventHandler(func(ctx *validation.Context, def *ast.SchemaDefinition) {
		def.RootOperationTypeDefinitions.ForEach(func(rotd ast.RootOperationTypeDefinition, i int) {
			// NOTE: Can't be extending here.
			switch rotd.OperationType {
			case ast.OperationDefinitionKindQuery:
				if ctx.SDLContext.QueryTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(rotd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.QueryTypeDefined = true
			case ast.OperationDefinitionKindMutation:
				if ctx.SDLContext.MutationTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(rotd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.MutationTypeDefined = true
			case ast.OperationDefinitionKindSubscription:
				if ctx.SDLContext.SubscriptionTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(rotd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.SubscriptionTypeDefined = true
			}
		})
	})

	w.AddSchemaExtensionLeaveEventHandler(func(ctx *validation.Context, ext *ast.SchemaExtension) {
		ext.OperationTypeDefinitions.ForEach(func(otd ast.OperationTypeDefinition, i int) {
			switch otd.OperationType {
			case ast.OperationDefinitionKindQuery:
				if ctx.Schema.QueryType != nil {
					ctx.AddError(existedOperationTypeMessage(otd.OperationType.String(), 0, 0))
				} else if ctx.SDLContext.QueryTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(otd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.QueryTypeDefined = true
			case ast.OperationDefinitionKindMutation:
				if ctx.Schema.MutationType != nil {
					ctx.AddError(existedOperationTypeMessage(otd.OperationType.String(), 0, 0))
				} else if ctx.SDLContext.MutationTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(otd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.MutationTypeDefined = true
			case ast.OperationDefinitionKindSubscription:
				if ctx.Schema.SubscriptionType != nil {
					ctx.AddError(existedOperationTypeMessage(otd.OperationType.String(), 0, 0))
				} else if ctx.SDLContext.SubscriptionTypeDefined {
					ctx.AddError(duplicateOperationTypeMessage(otd.OperationType.String(), 0, 0))
				}

				ctx.SDLContext.SubscriptionTypeDefined = true
			}
		})
	})
}

// duplicateOperationTypeMessage ...
func duplicateOperationTypeMessage(operation string, line, col int) graphql.Error {
	return graphql.NewError(
		"There can be only one " + operation + " type in schema.",
		// TODO: Location.
	)
}

// existedOperationTypeMessage ...
func existedOperationTypeMessage(operation string, line, col int) graphql.Error {
	return graphql.NewError(
		"Type for " + operation + " already defined in the schema. It cannot be redefined.",
		// TODO: Location.
	)
}
