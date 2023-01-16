// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/prompt"
)

// PromptDelete is the builder for deleting a Prompt entity.
type PromptDelete struct {
	config
	hooks    []Hook
	mutation *PromptMutation
}

// Where appends a list predicates to the PromptDelete builder.
func (pd *PromptDelete) Where(ps ...predicate.Prompt) *PromptDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PromptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PromptMutation](ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PromptDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PromptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: prompt.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prompt.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PromptDeleteOne is the builder for deleting a single Prompt entity.
type PromptDeleteOne struct {
	pd *PromptDelete
}

// Exec executes the deletion query.
func (pdo *PromptDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{prompt.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PromptDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
