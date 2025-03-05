package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"
	"github.com/google/uuid"
)

// UserCreditRecord holds the schema definition for the UserCreditRecord entity.
type UserCreditRecord struct {
	ent.Schema
}

func (UserCreditRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the UserCreditRecord.
func (UserCreditRecord) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("operation_type").
			Optional().
			Default(types.OperationType_DefaultOperationType.String()),
		field.
			Int32("credits_change").
			Optional().
			Default(0),
		field.
			String("extra").
			Optional().
			Default(""),
	}
}

// Edges of the UserCreditRecord.
func (UserCreditRecord) Edges() []ent.Edge {
	return nil
}
