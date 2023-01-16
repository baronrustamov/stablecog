// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/upscalerealtime"
)

// UpscaleRealtimeDelete is the builder for deleting a UpscaleRealtime entity.
type UpscaleRealtimeDelete struct {
	config
	hooks    []Hook
	mutation *UpscaleRealtimeMutation
}

// Where appends a list predicates to the UpscaleRealtimeDelete builder.
func (urd *UpscaleRealtimeDelete) Where(ps ...predicate.UpscaleRealtime) *UpscaleRealtimeDelete {
	urd.mutation.Where(ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UpscaleRealtimeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UpscaleRealtimeMutation](ctx, urd.sqlExec, urd.mutation, urd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UpscaleRealtimeDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UpscaleRealtimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: upscalerealtime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upscalerealtime.FieldID,
			},
		},
	}
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urd.mutation.done = true
	return affected, err
}

// UpscaleRealtimeDeleteOne is the builder for deleting a single UpscaleRealtime entity.
type UpscaleRealtimeDeleteOne struct {
	urd *UpscaleRealtimeDelete
}

// Exec executes the deletion query.
func (urdo *UpscaleRealtimeDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{upscalerealtime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UpscaleRealtimeDeleteOne) ExecX(ctx context.Context) {
	urdo.urd.ExecX(ctx)
}
