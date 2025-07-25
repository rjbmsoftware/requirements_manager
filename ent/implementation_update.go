// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"requirements/ent/implementation"
	"requirements/ent/predicate"
	"requirements/ent/requirement"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ImplementationUpdate is the builder for updating Implementation entities.
type ImplementationUpdate struct {
	config
	hooks    []Hook
	mutation *ImplementationMutation
}

// Where appends a list predicates to the ImplementationUpdate builder.
func (iu *ImplementationUpdate) Where(ps ...predicate.Implementation) *ImplementationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetURL sets the "url" field.
func (iu *ImplementationUpdate) SetURL(s string) *ImplementationUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iu *ImplementationUpdate) SetNillableURL(s *string) *ImplementationUpdate {
	if s != nil {
		iu.SetURL(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *ImplementationUpdate) SetDescription(s string) *ImplementationUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *ImplementationUpdate) SetNillableDescription(s *string) *ImplementationUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// AddRequirementIDs adds the "requirements" edge to the Requirement entity by IDs.
func (iu *ImplementationUpdate) AddRequirementIDs(ids ...uuid.UUID) *ImplementationUpdate {
	iu.mutation.AddRequirementIDs(ids...)
	return iu
}

// AddRequirements adds the "requirements" edges to the Requirement entity.
func (iu *ImplementationUpdate) AddRequirements(r ...*Requirement) *ImplementationUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iu.AddRequirementIDs(ids...)
}

// Mutation returns the ImplementationMutation object of the builder.
func (iu *ImplementationUpdate) Mutation() *ImplementationMutation {
	return iu.mutation
}

// ClearRequirements clears all "requirements" edges to the Requirement entity.
func (iu *ImplementationUpdate) ClearRequirements() *ImplementationUpdate {
	iu.mutation.ClearRequirements()
	return iu
}

// RemoveRequirementIDs removes the "requirements" edge to Requirement entities by IDs.
func (iu *ImplementationUpdate) RemoveRequirementIDs(ids ...uuid.UUID) *ImplementationUpdate {
	iu.mutation.RemoveRequirementIDs(ids...)
	return iu
}

// RemoveRequirements removes "requirements" edges to Requirement entities.
func (iu *ImplementationUpdate) RemoveRequirements(r ...*Requirement) *ImplementationUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iu.RemoveRequirementIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImplementationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImplementationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImplementationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImplementationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImplementationUpdate) check() error {
	if v, ok := iu.mutation.URL(); ok {
		if err := implementation.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Implementation.url": %w`, err)}
		}
	}
	return nil
}

func (iu *ImplementationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(implementation.Table, implementation.Columns, sqlgraph.NewFieldSpec(implementation.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.SetField(implementation.FieldURL, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(implementation.FieldDescription, field.TypeString, value)
	}
	if iu.mutation.RequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedRequirementsIDs(); len(nodes) > 0 && !iu.mutation.RequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RequirementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implementation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImplementationUpdateOne is the builder for updating a single Implementation entity.
type ImplementationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImplementationMutation
}

// SetURL sets the "url" field.
func (iuo *ImplementationUpdateOne) SetURL(s string) *ImplementationUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (iuo *ImplementationUpdateOne) SetNillableURL(s *string) *ImplementationUpdateOne {
	if s != nil {
		iuo.SetURL(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *ImplementationUpdateOne) SetDescription(s string) *ImplementationUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *ImplementationUpdateOne) SetNillableDescription(s *string) *ImplementationUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// AddRequirementIDs adds the "requirements" edge to the Requirement entity by IDs.
func (iuo *ImplementationUpdateOne) AddRequirementIDs(ids ...uuid.UUID) *ImplementationUpdateOne {
	iuo.mutation.AddRequirementIDs(ids...)
	return iuo
}

// AddRequirements adds the "requirements" edges to the Requirement entity.
func (iuo *ImplementationUpdateOne) AddRequirements(r ...*Requirement) *ImplementationUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iuo.AddRequirementIDs(ids...)
}

// Mutation returns the ImplementationMutation object of the builder.
func (iuo *ImplementationUpdateOne) Mutation() *ImplementationMutation {
	return iuo.mutation
}

// ClearRequirements clears all "requirements" edges to the Requirement entity.
func (iuo *ImplementationUpdateOne) ClearRequirements() *ImplementationUpdateOne {
	iuo.mutation.ClearRequirements()
	return iuo
}

// RemoveRequirementIDs removes the "requirements" edge to Requirement entities by IDs.
func (iuo *ImplementationUpdateOne) RemoveRequirementIDs(ids ...uuid.UUID) *ImplementationUpdateOne {
	iuo.mutation.RemoveRequirementIDs(ids...)
	return iuo
}

// RemoveRequirements removes "requirements" edges to Requirement entities.
func (iuo *ImplementationUpdateOne) RemoveRequirements(r ...*Requirement) *ImplementationUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return iuo.RemoveRequirementIDs(ids...)
}

// Where appends a list predicates to the ImplementationUpdate builder.
func (iuo *ImplementationUpdateOne) Where(ps ...predicate.Implementation) *ImplementationUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImplementationUpdateOne) Select(field string, fields ...string) *ImplementationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Implementation entity.
func (iuo *ImplementationUpdateOne) Save(ctx context.Context) (*Implementation, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImplementationUpdateOne) SaveX(ctx context.Context) *Implementation {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImplementationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImplementationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImplementationUpdateOne) check() error {
	if v, ok := iuo.mutation.URL(); ok {
		if err := implementation.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Implementation.url": %w`, err)}
		}
	}
	return nil
}

func (iuo *ImplementationUpdateOne) sqlSave(ctx context.Context) (_node *Implementation, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(implementation.Table, implementation.Columns, sqlgraph.NewFieldSpec(implementation.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Implementation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implementation.FieldID)
		for _, f := range fields {
			if !implementation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != implementation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.SetField(implementation.FieldURL, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(implementation.FieldDescription, field.TypeString, value)
	}
	if iuo.mutation.RequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedRequirementsIDs(); len(nodes) > 0 && !iuo.mutation.RequirementsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RequirementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   implementation.RequirementsTable,
			Columns: implementation.RequirementsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(requirement.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Implementation{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{implementation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
