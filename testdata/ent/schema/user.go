package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").Optional(),
		field.Int("spouse_id").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spouse", User.Type).Field("spouse_id").Unique(),
		// add children/parent edge:
		edge.To("children", User.Type),
		edge.From("parent", User.Type).Ref("children").Field("parent_id").Unique(),
		// add pets/owner:
		edge.To("pets", Pet.Type),
		// cards:
		edge.To("card", Card.Type).Unique(),
		// posts:
		edge.To("posts", Post.Type),
		// metadata:
		edge.To("metadata", Metadata.Type),
		// info
		edge.To("info", Info.Type),
	}
}
