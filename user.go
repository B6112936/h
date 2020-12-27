package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserEmail").NotEmpty().
			Unique(),
		field.String("UserName").NotEmpty(),
		
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("preemption", Preemption.Type).StorageKey(edge.Column("UserID")),
		edge.From("Role_ID", Role.Type).
			Ref("RoID").Unique(),
	}
}

