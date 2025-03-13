// Code generated by ent, DO NOT EDIT.

package usersubscriptionchange

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the usersubscriptionchange type in the database.
	Label = "user_subscription_change"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserSubscriptionID holds the string denoting the user_subscription_id field in the database.
	FieldUserSubscriptionID = "user_subscription_id"
	// FieldOldPackageID holds the string denoting the old_package_id field in the database.
	FieldOldPackageID = "old_package_id"
	// FieldNewPackageID holds the string denoting the new_package_id field in the database.
	FieldNewPackageID = "new_package_id"
	// Table holds the table name of the usersubscriptionchange in the database.
	Table = "user_subscription_changes"
)

// Columns holds all SQL columns for usersubscriptionchange fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldUserSubscriptionID,
	FieldOldPackageID,
	FieldNewPackageID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/billing-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultUserSubscriptionID holds the default value on creation for the "user_subscription_id" field.
	DefaultUserSubscriptionID func() uuid.UUID
	// DefaultOldPackageID holds the default value on creation for the "old_package_id" field.
	DefaultOldPackageID func() uuid.UUID
	// DefaultNewPackageID holds the default value on creation for the "new_package_id" field.
	DefaultNewPackageID func() uuid.UUID
)
