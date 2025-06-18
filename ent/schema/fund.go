package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Fund holds the schema definition for the Fund entity.
type Fund struct {
	ent.Schema
}

// Fields of the Fund.
func (Fund) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("Name").NotEmpty(),
	}
}

// Edges of the Fund.
func (Fund) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("deposits", Deposit.Type),
	}
}
