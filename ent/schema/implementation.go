package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Implementation holds the schema definition for the Requirement entity.
type Implementation struct {
	ent.Schema
}

// Fields of the Implementation.
func (Implementation) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").NotEmpty(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("description").Default(""),
	}
}

// Edges of the Implementation.
func (Implementation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("requirements", Requirement.Type).Ref("implementations"),

		edge.From("products", Product.Type).Ref("implementationsProduct").Unique(),
	}
}
