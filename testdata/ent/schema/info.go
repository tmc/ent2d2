package schema

import (
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Info holds the schema definition for the Info entity.
type Info struct {
	ent.Schema
}

// Fields of the Info.
func (Info) Fields() []ent.Field {
	// content json.Rawmessage
	return []ent.Field{
		field.JSON("content", json.RawMessage{}),
	}
}

// Edges of the Info.
func (Info) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("info").Unique(),
	}
}
