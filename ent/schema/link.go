package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.String("fingerprint").NotEmpty().Unique().Immutable(),
		field.String("label").Optional(),
		field.String("url").NotEmpty(),
		field.String("created_by").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("collection", Collection.Type).Ref("links").Unique(),
	}
}
