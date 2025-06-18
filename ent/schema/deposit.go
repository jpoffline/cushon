package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Deposit holds the schema definition for the Deposit entity.
type Deposit struct {
	ent.Schema
}

// Fields of the Deposit.
func (Deposit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Float("amount"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Deposit.
func (Deposit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fund", Fund.Type).
			Ref("deposits").
			Required(),
		edge.From("customer", Customer.Type).
			Ref("deposits").
			Required(),
	}
}
