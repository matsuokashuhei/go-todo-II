// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) *TagQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TagQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TagQuery {
	return t
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TaskQuery) CollectFields(ctx context.Context, satisfies ...string) *TaskQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		t = t.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return t
}

func (t *TaskQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TaskQuery {
	return t
}