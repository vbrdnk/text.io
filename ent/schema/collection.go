package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.String("fingerprint").NotEmpty().Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.Bool("published").Default(false),
		field.String("description").NotEmpty(),
		field.String("created_by").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", Link.Type),
	}
}
