// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"requirements/ent/implementation"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Implementation is the model entity for the Implementation schema.
type Implementation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImplementationQuery when eager-loading is set.
	Edges        ImplementationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ImplementationEdges holds the relations/edges for other nodes in the graph.
type ImplementationEdges struct {
	// Requirements holds the value of the requirements edge.
	Requirements []*Requirement `json:"requirements,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RequirementsOrErr returns the Requirements value or an error if the edge
// was not loaded in eager-loading.
func (e ImplementationEdges) RequirementsOrErr() ([]*Requirement, error) {
	if e.loadedTypes[0] {
		return e.Requirements, nil
	}
	return nil, &NotLoadedError{edge: "requirements"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Implementation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case implementation.FieldURL, implementation.FieldDescription:
			values[i] = new(sql.NullString)
		case implementation.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Implementation fields.
func (i *Implementation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case implementation.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case implementation.FieldURL:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[j])
			} else if value.Valid {
				i.URL = value.String
			}
		case implementation.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Implementation.
// This includes values selected through modifiers, order, etc.
func (i *Implementation) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryRequirements queries the "requirements" edge of the Implementation entity.
func (i *Implementation) QueryRequirements() *RequirementQuery {
	return NewImplementationClient(i.config).QueryRequirements(i)
}

// Update returns a builder for updating this Implementation.
// Note that you need to call Implementation.Unwrap() before calling this method if this Implementation
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Implementation) Update() *ImplementationUpdateOne {
	return NewImplementationClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Implementation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Implementation) Unwrap() *Implementation {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Implementation is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Implementation) String() string {
	var builder strings.Builder
	builder.WriteString("Implementation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("url=")
	builder.WriteString(i.URL)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Implementations is a parsable slice of Implementation.
type Implementations []*Implementation
