package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Requirement holds the schema definition for the Requirement entity.
type Requirement struct {
	ent.Schema
}

// Fields of the Requirement.
func (Requirement) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("path").NotEmpty(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("description").Default(""),
	}
}

// Edges of the Requirement.
func (Requirement) Edges() []ent.Edge {
	return nil
}
