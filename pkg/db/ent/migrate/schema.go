// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddonsColumns holds the columns for the "addons" table.
	AddonsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "credit", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "sort_order", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "enabled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AddonsTable holds the schema information for the "addons" table.
	AddonsTable = &schema.Table{
		Name:       "addons",
		Columns:    AddonsColumns,
		PrimaryKey: []*schema.Column{AddonsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "addon_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AddonsColumns[4]},
			},
		},
	}
	// ExchangesColumns holds the columns for the "exchanges" table.
	ExchangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "usage_type", Type: field.TypeString, Nullable: true, Default: "DefaultUsageType"},
		{Name: "credit", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "exchange_threshold", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "path", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// ExchangesTable holds the schema information for the "exchanges" table.
	ExchangesTable = &schema.Table{
		Name:       "exchanges",
		Columns:    ExchangesColumns,
		PrimaryKey: []*schema.Column{ExchangesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "exchange_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ExchangesColumns[4]},
			},
		},
	}
	// IgnoreIdsColumns holds the columns for the "ignore_ids" table.
	IgnoreIdsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "sample_col", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// IgnoreIdsTable holds the schema information for the "ignore_ids" table.
	IgnoreIdsTable = &schema.Table{
		Name:       "ignore_ids",
		Columns:    IgnoreIdsColumns,
		PrimaryKey: []*schema.Column{IgnoreIdsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ignoreid_ent_id",
				Unique:  true,
				Columns: []*schema.Column{IgnoreIdsColumns[4]},
			},
		},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PubsubMessagesColumns[4]},
			},
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[7]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[8]},
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "package_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "sort_order", Type: field.TypeUint32, Nullable: true},
		{Name: "package_type", Type: field.TypeString, Nullable: true, Default: "DefaultPackageType"},
		{Name: "credit", Type: field.TypeUint32, Nullable: true},
		{Name: "reset_type", Type: field.TypeString, Nullable: true, Default: "DefaultResetType"},
		{Name: "qps_limit", Type: field.TypeUint32, Nullable: true},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subscription_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SubscriptionsColumns[4]},
			},
		},
	}
	// UserCreditRecordsColumns holds the columns for the "user_credit_records" table.
	UserCreditRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "operation_type", Type: field.TypeString, Nullable: true, Default: "DefaultOperationType"},
		{Name: "credits_change", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "extra", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// UserCreditRecordsTable holds the schema information for the "user_credit_records" table.
	UserCreditRecordsTable = &schema.Table{
		Name:       "user_credit_records",
		Columns:    UserCreditRecordsColumns,
		PrimaryKey: []*schema.Column{UserCreditRecordsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "usercreditrecord_ent_id",
				Unique:  true,
				Columns: []*schema.Column{UserCreditRecordsColumns[4]},
			},
		},
	}
	// UserSubscriptionsColumns holds the columns for the "user_subscriptions" table.
	UserSubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "package_id", Type: field.TypeUUID, Nullable: true},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "usage_state", Type: field.TypeString, Nullable: true, Default: "DefaultUsageState"},
		{Name: "subscription_credit", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "addon_credit", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// UserSubscriptionsTable holds the schema information for the "user_subscriptions" table.
	UserSubscriptionsTable = &schema.Table{
		Name:       "user_subscriptions",
		Columns:    UserSubscriptionsColumns,
		PrimaryKey: []*schema.Column{UserSubscriptionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "usersubscription_ent_id",
				Unique:  true,
				Columns: []*schema.Column{UserSubscriptionsColumns[4]},
			},
		},
	}
	// UserSubscriptionChangesColumns holds the columns for the "user_subscription_changes" table.
	UserSubscriptionChangesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_subscription_id", Type: field.TypeUUID, Nullable: true},
		{Name: "old_package_id", Type: field.TypeUUID, Nullable: true},
		{Name: "new_package_id", Type: field.TypeUUID, Nullable: true},
	}
	// UserSubscriptionChangesTable holds the schema information for the "user_subscription_changes" table.
	UserSubscriptionChangesTable = &schema.Table{
		Name:       "user_subscription_changes",
		Columns:    UserSubscriptionChangesColumns,
		PrimaryKey: []*schema.Column{UserSubscriptionChangesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "usersubscriptionchange_ent_id",
				Unique:  true,
				Columns: []*schema.Column{UserSubscriptionChangesColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddonsTable,
		ExchangesTable,
		IgnoreIdsTable,
		PubsubMessagesTable,
		SubscriptionsTable,
		UserCreditRecordsTable,
		UserSubscriptionsTable,
		UserSubscriptionChangesTable,
	}
)

func init() {
}
