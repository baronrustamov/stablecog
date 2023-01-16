// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/negativeprompt"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
)

// NegativePromptUpdate is the builder for updating NegativePrompt entities.
type NegativePromptUpdate struct {
	config
	hooks    []Hook
	mutation *NegativePromptMutation
}

// Where appends a list predicates to the NegativePromptUpdate builder.
func (npu *NegativePromptUpdate) Where(ps ...predicate.NegativePrompt) *NegativePromptUpdate {
	npu.mutation.Where(ps...)
	return npu
}

// Mutation returns the NegativePromptMutation object of the builder.
func (npu *NegativePromptUpdate) Mutation() *NegativePromptMutation {
	return npu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (npu *NegativePromptUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, NegativePromptMutation](ctx, npu.sqlSave, npu.mutation, npu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (npu *NegativePromptUpdate) SaveX(ctx context.Context) int {
	affected, err := npu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (npu *NegativePromptUpdate) Exec(ctx context.Context) error {
	_, err := npu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npu *NegativePromptUpdate) ExecX(ctx context.Context) {
	if err := npu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (npu *NegativePromptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   negativeprompt.Table,
			Columns: negativeprompt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: negativeprompt.FieldID,
			},
		},
	}
	if ps := npu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, npu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{negativeprompt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	npu.mutation.done = true
	return n, nil
}

// NegativePromptUpdateOne is the builder for updating a single NegativePrompt entity.
type NegativePromptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NegativePromptMutation
}

// Mutation returns the NegativePromptMutation object of the builder.
func (npuo *NegativePromptUpdateOne) Mutation() *NegativePromptMutation {
	return npuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (npuo *NegativePromptUpdateOne) Select(field string, fields ...string) *NegativePromptUpdateOne {
	npuo.fields = append([]string{field}, fields...)
	return npuo
}

// Save executes the query and returns the updated NegativePrompt entity.
func (npuo *NegativePromptUpdateOne) Save(ctx context.Context) (*NegativePrompt, error) {
	return withHooks[*NegativePrompt, NegativePromptMutation](ctx, npuo.sqlSave, npuo.mutation, npuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (npuo *NegativePromptUpdateOne) SaveX(ctx context.Context) *NegativePrompt {
	node, err := npuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (npuo *NegativePromptUpdateOne) Exec(ctx context.Context) error {
	_, err := npuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (npuo *NegativePromptUpdateOne) ExecX(ctx context.Context) {
	if err := npuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (npuo *NegativePromptUpdateOne) sqlSave(ctx context.Context) (_node *NegativePrompt, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   negativeprompt.Table,
			Columns: negativeprompt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: negativeprompt.FieldID,
			},
		},
	}
	id, ok := npuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NegativePrompt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := npuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, negativeprompt.FieldID)
		for _, f := range fields {
			if !negativeprompt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != negativeprompt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := npuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &NegativePrompt{config: npuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, npuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{negativeprompt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	npuo.mutation.done = true
	return _node, nil
}
