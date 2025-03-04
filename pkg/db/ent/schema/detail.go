package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// Detail holds the schema definition for the Detail entity.
type Detail struct {
	ent.Schema
}

func (Detail) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Detail.
func (Detail) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("sample_col").
			Optional().
			Default(""),
	}
}

// Edges of the Detail.
func (Detail) Edges() []ent.Edge {
	return nil
}
