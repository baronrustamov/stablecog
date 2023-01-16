// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/upscalerealtime"
)

// UpscaleRealtimeCreate is the builder for creating a UpscaleRealtime entity.
type UpscaleRealtimeCreate struct {
	config
	mutation *UpscaleRealtimeMutation
	hooks    []Hook
}

// Mutation returns the UpscaleRealtimeMutation object of the builder.
func (urc *UpscaleRealtimeCreate) Mutation() *UpscaleRealtimeMutation {
	return urc.mutation
}

// Save creates the UpscaleRealtime in the database.
func (urc *UpscaleRealtimeCreate) Save(ctx context.Context) (*UpscaleRealtime, error) {
	return withHooks[*UpscaleRealtime, UpscaleRealtimeMutation](ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UpscaleRealtimeCreate) SaveX(ctx context.Context) *UpscaleRealtime {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UpscaleRealtimeCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UpscaleRealtimeCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UpscaleRealtimeCreate) check() error {
	return nil
}

func (urc *UpscaleRealtimeCreate) sqlSave(ctx context.Context) (*UpscaleRealtime, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UpscaleRealtimeCreate) createSpec() (*UpscaleRealtime, *sqlgraph.CreateSpec) {
	var (
		_node = &UpscaleRealtime{config: urc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: upscalerealtime.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: upscalerealtime.FieldID,
			},
		}
	)
	return _node, _spec
}

// UpscaleRealtimeCreateBulk is the builder for creating many UpscaleRealtime entities in bulk.
type UpscaleRealtimeCreateBulk struct {
	config
	builders []*UpscaleRealtimeCreate
}

// Save creates the UpscaleRealtime entities in the database.
func (urcb *UpscaleRealtimeCreateBulk) Save(ctx context.Context) ([]*UpscaleRealtime, error) {
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UpscaleRealtime, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UpscaleRealtimeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UpscaleRealtimeCreateBulk) SaveX(ctx context.Context) []*UpscaleRealtime {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UpscaleRealtimeCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UpscaleRealtimeCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
