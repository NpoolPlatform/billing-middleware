package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// IgnoreID holds the schema definition for the IgnoreID entity.
type IgnoreID struct {
	ent.Schema
}

func (IgnoreID) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the IgnoreID.
func (IgnoreID) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("sample_col").
			Optional().
			Default(""),
	}
}

// Edges of the IgnoreID.
func (IgnoreID) Edges() []ent.Edge {
	return nil
}
