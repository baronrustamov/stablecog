// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/generationrealtime"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
)

// GenerationRealtimeDelete is the builder for deleting a GenerationRealtime entity.
type GenerationRealtimeDelete struct {
	config
	hooks    []Hook
	mutation *GenerationRealtimeMutation
}

// Where appends a list predicates to the GenerationRealtimeDelete builder.
func (grd *GenerationRealtimeDelete) Where(ps ...predicate.GenerationRealtime) *GenerationRealtimeDelete {
	grd.mutation.Where(ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GenerationRealtimeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GenerationRealtimeMutation](ctx, grd.sqlExec, grd.mutation, grd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GenerationRealtimeDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grd *GenerationRealtimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: generationrealtime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: generationrealtime.FieldID,
			},
		},
	}
	if ps := grd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, grd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	grd.mutation.done = true
	return affected, err
}

// GenerationRealtimeDeleteOne is the builder for deleting a single GenerationRealtime entity.
type GenerationRealtimeDeleteOne struct {
	grd *GenerationRealtimeDelete
}

// Exec executes the deletion query.
func (grdo *GenerationRealtimeDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{generationrealtime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GenerationRealtimeDeleteOne) ExecX(ctx context.Context) {
	grdo.grd.ExecX(ctx)
}
