// Code generated by ent, DO NOT EDIT.

package usersubscription

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// PackageID applies equality check predicate on the "package_id" field. It's identical to PackageIDEQ.
func PackageID(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPackageID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// UsageState applies equality check predicate on the "usage_state" field. It's identical to UsageStateEQ.
func UsageState(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsageState), v))
	})
}

// SubscriptionCredit applies equality check predicate on the "subscription_credit" field. It's identical to SubscriptionCreditEQ.
func SubscriptionCredit(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubscriptionCredit), v))
	})
}

// AddonCredit applies equality check predicate on the "addon_credit" field. It's identical to AddonCreditEQ.
func AddonCredit(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddonCredit), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// PackageIDEQ applies the EQ predicate on the "package_id" field.
func PackageIDEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPackageID), v))
	})
}

// PackageIDNEQ applies the NEQ predicate on the "package_id" field.
func PackageIDNEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPackageID), v))
	})
}

// PackageIDIn applies the In predicate on the "package_id" field.
func PackageIDIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPackageID), v...))
	})
}

// PackageIDNotIn applies the NotIn predicate on the "package_id" field.
func PackageIDNotIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPackageID), v...))
	})
}

// PackageIDGT applies the GT predicate on the "package_id" field.
func PackageIDGT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPackageID), v))
	})
}

// PackageIDGTE applies the GTE predicate on the "package_id" field.
func PackageIDGTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPackageID), v))
	})
}

// PackageIDLT applies the LT predicate on the "package_id" field.
func PackageIDLT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPackageID), v))
	})
}

// PackageIDLTE applies the LTE predicate on the "package_id" field.
func PackageIDLTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPackageID), v))
	})
}

// PackageIDIsNil applies the IsNil predicate on the "package_id" field.
func PackageIDIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPackageID)))
	})
}

// PackageIDNotNil applies the NotNil predicate on the "package_id" field.
func PackageIDNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPackageID)))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndAt), v))
	})
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndAt), v))
	})
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEndAt), v...))
	})
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEndAt), v...))
	})
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndAt), v))
	})
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndAt), v))
	})
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndAt), v))
	})
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndAt), v))
	})
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndAt)))
	})
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndAt)))
	})
}

// UsageStateEQ applies the EQ predicate on the "usage_state" field.
func UsageStateEQ(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsageState), v))
	})
}

// UsageStateNEQ applies the NEQ predicate on the "usage_state" field.
func UsageStateNEQ(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsageState), v))
	})
}

// UsageStateIn applies the In predicate on the "usage_state" field.
func UsageStateIn(vs ...string) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsageState), v...))
	})
}

// UsageStateNotIn applies the NotIn predicate on the "usage_state" field.
func UsageStateNotIn(vs ...string) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsageState), v...))
	})
}

// UsageStateGT applies the GT predicate on the "usage_state" field.
func UsageStateGT(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsageState), v))
	})
}

// UsageStateGTE applies the GTE predicate on the "usage_state" field.
func UsageStateGTE(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsageState), v))
	})
}

// UsageStateLT applies the LT predicate on the "usage_state" field.
func UsageStateLT(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsageState), v))
	})
}

// UsageStateLTE applies the LTE predicate on the "usage_state" field.
func UsageStateLTE(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsageState), v))
	})
}

// UsageStateContains applies the Contains predicate on the "usage_state" field.
func UsageStateContains(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsageState), v))
	})
}

// UsageStateHasPrefix applies the HasPrefix predicate on the "usage_state" field.
func UsageStateHasPrefix(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsageState), v))
	})
}

// UsageStateHasSuffix applies the HasSuffix predicate on the "usage_state" field.
func UsageStateHasSuffix(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsageState), v))
	})
}

// UsageStateIsNil applies the IsNil predicate on the "usage_state" field.
func UsageStateIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsageState)))
	})
}

// UsageStateNotNil applies the NotNil predicate on the "usage_state" field.
func UsageStateNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsageState)))
	})
}

// UsageStateEqualFold applies the EqualFold predicate on the "usage_state" field.
func UsageStateEqualFold(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsageState), v))
	})
}

// UsageStateContainsFold applies the ContainsFold predicate on the "usage_state" field.
func UsageStateContainsFold(v string) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsageState), v))
	})
}

// SubscriptionCreditEQ applies the EQ predicate on the "subscription_credit" field.
func SubscriptionCreditEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditNEQ applies the NEQ predicate on the "subscription_credit" field.
func SubscriptionCreditNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditIn applies the In predicate on the "subscription_credit" field.
func SubscriptionCreditIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSubscriptionCredit), v...))
	})
}

// SubscriptionCreditNotIn applies the NotIn predicate on the "subscription_credit" field.
func SubscriptionCreditNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSubscriptionCredit), v...))
	})
}

// SubscriptionCreditGT applies the GT predicate on the "subscription_credit" field.
func SubscriptionCreditGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditGTE applies the GTE predicate on the "subscription_credit" field.
func SubscriptionCreditGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditLT applies the LT predicate on the "subscription_credit" field.
func SubscriptionCreditLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditLTE applies the LTE predicate on the "subscription_credit" field.
func SubscriptionCreditLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubscriptionCredit), v))
	})
}

// SubscriptionCreditIsNil applies the IsNil predicate on the "subscription_credit" field.
func SubscriptionCreditIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSubscriptionCredit)))
	})
}

// SubscriptionCreditNotNil applies the NotNil predicate on the "subscription_credit" field.
func SubscriptionCreditNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSubscriptionCredit)))
	})
}

// AddonCreditEQ applies the EQ predicate on the "addon_credit" field.
func AddonCreditEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditNEQ applies the NEQ predicate on the "addon_credit" field.
func AddonCreditNEQ(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditIn applies the In predicate on the "addon_credit" field.
func AddonCreditIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAddonCredit), v...))
	})
}

// AddonCreditNotIn applies the NotIn predicate on the "addon_credit" field.
func AddonCreditNotIn(vs ...uint32) predicate.UserSubscription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAddonCredit), v...))
	})
}

// AddonCreditGT applies the GT predicate on the "addon_credit" field.
func AddonCreditGT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditGTE applies the GTE predicate on the "addon_credit" field.
func AddonCreditGTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditLT applies the LT predicate on the "addon_credit" field.
func AddonCreditLT(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditLTE applies the LTE predicate on the "addon_credit" field.
func AddonCreditLTE(v uint32) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddonCredit), v))
	})
}

// AddonCreditIsNil applies the IsNil predicate on the "addon_credit" field.
func AddonCreditIsNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAddonCredit)))
	})
}

// AddonCreditNotNil applies the NotNil predicate on the "addon_credit" field.
func AddonCreditNotNil() predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAddonCredit)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserSubscription) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserSubscription) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserSubscription) predicate.UserSubscription {
	return predicate.UserSubscription(func(s *sql.Selector) {
		p(s.Not())
	})
}
