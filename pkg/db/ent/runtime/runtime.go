// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/addon"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/detail"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/exchange"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/ignoreid"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/pubsubmessage"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/schema"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/subscription"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usercreditrecord"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usersubscription"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addonMixin := schema.Addon{}.Mixin()
	addon.Policy = privacy.NewPolicies(addonMixin[0], schema.Addon{})
	addon.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := addon.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	addonMixinFields0 := addonMixin[0].Fields()
	_ = addonMixinFields0
	addonMixinFields1 := addonMixin[1].Fields()
	_ = addonMixinFields1
	addonFields := schema.Addon{}.Fields()
	_ = addonFields
	// addonDescCreatedAt is the schema descriptor for created_at field.
	addonDescCreatedAt := addonMixinFields0[0].Descriptor()
	// addon.DefaultCreatedAt holds the default value on creation for the created_at field.
	addon.DefaultCreatedAt = addonDescCreatedAt.Default.(func() uint32)
	// addonDescUpdatedAt is the schema descriptor for updated_at field.
	addonDescUpdatedAt := addonMixinFields0[1].Descriptor()
	// addon.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	addon.DefaultUpdatedAt = addonDescUpdatedAt.Default.(func() uint32)
	// addon.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	addon.UpdateDefaultUpdatedAt = addonDescUpdatedAt.UpdateDefault.(func() uint32)
	// addonDescDeletedAt is the schema descriptor for deleted_at field.
	addonDescDeletedAt := addonMixinFields0[2].Descriptor()
	// addon.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	addon.DefaultDeletedAt = addonDescDeletedAt.Default.(func() uint32)
	// addonDescEntID is the schema descriptor for ent_id field.
	addonDescEntID := addonMixinFields1[1].Descriptor()
	// addon.DefaultEntID holds the default value on creation for the ent_id field.
	addon.DefaultEntID = addonDescEntID.Default.(func() uuid.UUID)
	// addonDescAppID is the schema descriptor for app_id field.
	addonDescAppID := addonFields[0].Descriptor()
	// addon.DefaultAppID holds the default value on creation for the app_id field.
	addon.DefaultAppID = addonDescAppID.Default.(func() uuid.UUID)
	// addonDescPrice is the schema descriptor for price field.
	addonDescPrice := addonFields[1].Descriptor()
	// addon.DefaultPrice holds the default value on creation for the price field.
	addon.DefaultPrice = addonDescPrice.Default.(decimal.Decimal)
	// addonDescCredit is the schema descriptor for credit field.
	addonDescCredit := addonFields[2].Descriptor()
	// addon.DefaultCredit holds the default value on creation for the credit field.
	addon.DefaultCredit = addonDescCredit.Default.(uint32)
	// addonDescSortOrder is the schema descriptor for sort_order field.
	addonDescSortOrder := addonFields[3].Descriptor()
	// addon.DefaultSortOrder holds the default value on creation for the sort_order field.
	addon.DefaultSortOrder = addonDescSortOrder.Default.(uint32)
	// addonDescEnabled is the schema descriptor for enabled field.
	addonDescEnabled := addonFields[4].Descriptor()
	// addon.DefaultEnabled holds the default value on creation for the enabled field.
	addon.DefaultEnabled = addonDescEnabled.Default.(bool)
	// addonDescDescription is the schema descriptor for description field.
	addonDescDescription := addonFields[5].Descriptor()
	// addon.DefaultDescription holds the default value on creation for the description field.
	addon.DefaultDescription = addonDescDescription.Default.(string)
	detailMixin := schema.Detail{}.Mixin()
	detail.Policy = privacy.NewPolicies(detailMixin[0], schema.Detail{})
	detail.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := detail.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	detailMixinFields0 := detailMixin[0].Fields()
	_ = detailMixinFields0
	detailMixinFields1 := detailMixin[1].Fields()
	_ = detailMixinFields1
	detailFields := schema.Detail{}.Fields()
	_ = detailFields
	// detailDescCreatedAt is the schema descriptor for created_at field.
	detailDescCreatedAt := detailMixinFields0[0].Descriptor()
	// detail.DefaultCreatedAt holds the default value on creation for the created_at field.
	detail.DefaultCreatedAt = detailDescCreatedAt.Default.(func() uint32)
	// detailDescUpdatedAt is the schema descriptor for updated_at field.
	detailDescUpdatedAt := detailMixinFields0[1].Descriptor()
	// detail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	detail.DefaultUpdatedAt = detailDescUpdatedAt.Default.(func() uint32)
	// detail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	detail.UpdateDefaultUpdatedAt = detailDescUpdatedAt.UpdateDefault.(func() uint32)
	// detailDescDeletedAt is the schema descriptor for deleted_at field.
	detailDescDeletedAt := detailMixinFields0[2].Descriptor()
	// detail.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	detail.DefaultDeletedAt = detailDescDeletedAt.Default.(func() uint32)
	// detailDescEntID is the schema descriptor for ent_id field.
	detailDescEntID := detailMixinFields1[1].Descriptor()
	// detail.DefaultEntID holds the default value on creation for the ent_id field.
	detail.DefaultEntID = detailDescEntID.Default.(func() uuid.UUID)
	// detailDescSampleCol is the schema descriptor for sample_col field.
	detailDescSampleCol := detailFields[0].Descriptor()
	// detail.DefaultSampleCol holds the default value on creation for the sample_col field.
	detail.DefaultSampleCol = detailDescSampleCol.Default.(string)
	exchangeMixin := schema.Exchange{}.Mixin()
	exchange.Policy = privacy.NewPolicies(exchangeMixin[0], schema.Exchange{})
	exchange.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := exchange.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	exchangeMixinFields0 := exchangeMixin[0].Fields()
	_ = exchangeMixinFields0
	exchangeMixinFields1 := exchangeMixin[1].Fields()
	_ = exchangeMixinFields1
	exchangeFields := schema.Exchange{}.Fields()
	_ = exchangeFields
	// exchangeDescCreatedAt is the schema descriptor for created_at field.
	exchangeDescCreatedAt := exchangeMixinFields0[0].Descriptor()
	// exchange.DefaultCreatedAt holds the default value on creation for the created_at field.
	exchange.DefaultCreatedAt = exchangeDescCreatedAt.Default.(func() uint32)
	// exchangeDescUpdatedAt is the schema descriptor for updated_at field.
	exchangeDescUpdatedAt := exchangeMixinFields0[1].Descriptor()
	// exchange.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	exchange.DefaultUpdatedAt = exchangeDescUpdatedAt.Default.(func() uint32)
	// exchange.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	exchange.UpdateDefaultUpdatedAt = exchangeDescUpdatedAt.UpdateDefault.(func() uint32)
	// exchangeDescDeletedAt is the schema descriptor for deleted_at field.
	exchangeDescDeletedAt := exchangeMixinFields0[2].Descriptor()
	// exchange.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	exchange.DefaultDeletedAt = exchangeDescDeletedAt.Default.(func() uint32)
	// exchangeDescEntID is the schema descriptor for ent_id field.
	exchangeDescEntID := exchangeMixinFields1[1].Descriptor()
	// exchange.DefaultEntID holds the default value on creation for the ent_id field.
	exchange.DefaultEntID = exchangeDescEntID.Default.(func() uuid.UUID)
	// exchangeDescAppID is the schema descriptor for app_id field.
	exchangeDescAppID := exchangeFields[0].Descriptor()
	// exchange.DefaultAppID holds the default value on creation for the app_id field.
	exchange.DefaultAppID = exchangeDescAppID.Default.(func() uuid.UUID)
	// exchangeDescUsageType is the schema descriptor for usage_type field.
	exchangeDescUsageType := exchangeFields[1].Descriptor()
	// exchange.DefaultUsageType holds the default value on creation for the usage_type field.
	exchange.DefaultUsageType = exchangeDescUsageType.Default.(string)
	// exchangeDescCredit is the schema descriptor for credit field.
	exchangeDescCredit := exchangeFields[2].Descriptor()
	// exchange.DefaultCredit holds the default value on creation for the credit field.
	exchange.DefaultCredit = exchangeDescCredit.Default.(uint32)
	// exchangeDescExchangeThreshold is the schema descriptor for exchange_threshold field.
	exchangeDescExchangeThreshold := exchangeFields[3].Descriptor()
	// exchange.DefaultExchangeThreshold holds the default value on creation for the exchange_threshold field.
	exchange.DefaultExchangeThreshold = exchangeDescExchangeThreshold.Default.(uint32)
	ignoreidMixin := schema.IgnoreID{}.Mixin()
	ignoreid.Policy = privacy.NewPolicies(ignoreidMixin[0], schema.IgnoreID{})
	ignoreid.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := ignoreid.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	ignoreidMixinFields0 := ignoreidMixin[0].Fields()
	_ = ignoreidMixinFields0
	ignoreidMixinFields1 := ignoreidMixin[1].Fields()
	_ = ignoreidMixinFields1
	ignoreidFields := schema.IgnoreID{}.Fields()
	_ = ignoreidFields
	// ignoreidDescCreatedAt is the schema descriptor for created_at field.
	ignoreidDescCreatedAt := ignoreidMixinFields0[0].Descriptor()
	// ignoreid.DefaultCreatedAt holds the default value on creation for the created_at field.
	ignoreid.DefaultCreatedAt = ignoreidDescCreatedAt.Default.(func() uint32)
	// ignoreidDescUpdatedAt is the schema descriptor for updated_at field.
	ignoreidDescUpdatedAt := ignoreidMixinFields0[1].Descriptor()
	// ignoreid.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ignoreid.DefaultUpdatedAt = ignoreidDescUpdatedAt.Default.(func() uint32)
	// ignoreid.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ignoreid.UpdateDefaultUpdatedAt = ignoreidDescUpdatedAt.UpdateDefault.(func() uint32)
	// ignoreidDescDeletedAt is the schema descriptor for deleted_at field.
	ignoreidDescDeletedAt := ignoreidMixinFields0[2].Descriptor()
	// ignoreid.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	ignoreid.DefaultDeletedAt = ignoreidDescDeletedAt.Default.(func() uint32)
	// ignoreidDescEntID is the schema descriptor for ent_id field.
	ignoreidDescEntID := ignoreidMixinFields1[1].Descriptor()
	// ignoreid.DefaultEntID holds the default value on creation for the ent_id field.
	ignoreid.DefaultEntID = ignoreidDescEntID.Default.(func() uuid.UUID)
	// ignoreidDescSampleCol is the schema descriptor for sample_col field.
	ignoreidDescSampleCol := ignoreidFields[0].Descriptor()
	// ignoreid.DefaultSampleCol holds the default value on creation for the sample_col field.
	ignoreid.DefaultSampleCol = ignoreidDescSampleCol.Default.(string)
	pubsubmessageMixin := schema.PubsubMessage{}.Mixin()
	pubsubmessage.Policy = privacy.NewPolicies(pubsubmessageMixin[0], schema.PubsubMessage{})
	pubsubmessage.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := pubsubmessage.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	pubsubmessageMixinFields0 := pubsubmessageMixin[0].Fields()
	_ = pubsubmessageMixinFields0
	pubsubmessageMixinFields1 := pubsubmessageMixin[1].Fields()
	_ = pubsubmessageMixinFields1
	pubsubmessageFields := schema.PubsubMessage{}.Fields()
	_ = pubsubmessageFields
	// pubsubmessageDescCreatedAt is the schema descriptor for created_at field.
	pubsubmessageDescCreatedAt := pubsubmessageMixinFields0[0].Descriptor()
	// pubsubmessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	pubsubmessage.DefaultCreatedAt = pubsubmessageDescCreatedAt.Default.(func() uint32)
	// pubsubmessageDescUpdatedAt is the schema descriptor for updated_at field.
	pubsubmessageDescUpdatedAt := pubsubmessageMixinFields0[1].Descriptor()
	// pubsubmessage.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pubsubmessage.DefaultUpdatedAt = pubsubmessageDescUpdatedAt.Default.(func() uint32)
	// pubsubmessage.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pubsubmessage.UpdateDefaultUpdatedAt = pubsubmessageDescUpdatedAt.UpdateDefault.(func() uint32)
	// pubsubmessageDescDeletedAt is the schema descriptor for deleted_at field.
	pubsubmessageDescDeletedAt := pubsubmessageMixinFields0[2].Descriptor()
	// pubsubmessage.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	pubsubmessage.DefaultDeletedAt = pubsubmessageDescDeletedAt.Default.(func() uint32)
	// pubsubmessageDescEntID is the schema descriptor for ent_id field.
	pubsubmessageDescEntID := pubsubmessageMixinFields1[1].Descriptor()
	// pubsubmessage.DefaultEntID holds the default value on creation for the ent_id field.
	pubsubmessage.DefaultEntID = pubsubmessageDescEntID.Default.(func() uuid.UUID)
	// pubsubmessageDescMessageID is the schema descriptor for message_id field.
	pubsubmessageDescMessageID := pubsubmessageFields[0].Descriptor()
	// pubsubmessage.DefaultMessageID holds the default value on creation for the message_id field.
	pubsubmessage.DefaultMessageID = pubsubmessageDescMessageID.Default.(string)
	// pubsubmessageDescState is the schema descriptor for state field.
	pubsubmessageDescState := pubsubmessageFields[1].Descriptor()
	// pubsubmessage.DefaultState holds the default value on creation for the state field.
	pubsubmessage.DefaultState = pubsubmessageDescState.Default.(string)
	// pubsubmessageDescRespToID is the schema descriptor for resp_to_id field.
	pubsubmessageDescRespToID := pubsubmessageFields[2].Descriptor()
	// pubsubmessage.DefaultRespToID holds the default value on creation for the resp_to_id field.
	pubsubmessage.DefaultRespToID = pubsubmessageDescRespToID.Default.(func() uuid.UUID)
	// pubsubmessageDescUndoID is the schema descriptor for undo_id field.
	pubsubmessageDescUndoID := pubsubmessageFields[3].Descriptor()
	// pubsubmessage.DefaultUndoID holds the default value on creation for the undo_id field.
	pubsubmessage.DefaultUndoID = pubsubmessageDescUndoID.Default.(func() uuid.UUID)
	// pubsubmessageDescArguments is the schema descriptor for arguments field.
	pubsubmessageDescArguments := pubsubmessageFields[4].Descriptor()
	// pubsubmessage.DefaultArguments holds the default value on creation for the arguments field.
	pubsubmessage.DefaultArguments = pubsubmessageDescArguments.Default.(string)
	subscriptionMixin := schema.Subscription{}.Mixin()
	subscription.Policy = privacy.NewPolicies(subscriptionMixin[0], schema.Subscription{})
	subscription.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := subscription.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	subscriptionMixinFields0 := subscriptionMixin[0].Fields()
	_ = subscriptionMixinFields0
	subscriptionMixinFields1 := subscriptionMixin[1].Fields()
	_ = subscriptionMixinFields1
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionMixinFields0[0].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() uint32)
	// subscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	subscriptionDescUpdatedAt := subscriptionMixinFields0[1].Descriptor()
	// subscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	subscription.DefaultUpdatedAt = subscriptionDescUpdatedAt.Default.(func() uint32)
	// subscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	subscription.UpdateDefaultUpdatedAt = subscriptionDescUpdatedAt.UpdateDefault.(func() uint32)
	// subscriptionDescDeletedAt is the schema descriptor for deleted_at field.
	subscriptionDescDeletedAt := subscriptionMixinFields0[2].Descriptor()
	// subscription.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	subscription.DefaultDeletedAt = subscriptionDescDeletedAt.Default.(func() uint32)
	// subscriptionDescEntID is the schema descriptor for ent_id field.
	subscriptionDescEntID := subscriptionMixinFields1[1].Descriptor()
	// subscription.DefaultEntID holds the default value on creation for the ent_id field.
	subscription.DefaultEntID = subscriptionDescEntID.Default.(func() uuid.UUID)
	// subscriptionDescAppID is the schema descriptor for app_id field.
	subscriptionDescAppID := subscriptionFields[0].Descriptor()
	// subscription.DefaultAppID holds the default value on creation for the app_id field.
	subscription.DefaultAppID = subscriptionDescAppID.Default.(func() uuid.UUID)
	// subscriptionDescPackageName is the schema descriptor for package_name field.
	subscriptionDescPackageName := subscriptionFields[1].Descriptor()
	// subscription.DefaultPackageName holds the default value on creation for the package_name field.
	subscription.DefaultPackageName = subscriptionDescPackageName.Default.(string)
	// subscriptionDescPrice is the schema descriptor for price field.
	subscriptionDescPrice := subscriptionFields[2].Descriptor()
	// subscription.DefaultPrice holds the default value on creation for the price field.
	subscription.DefaultPrice = subscriptionDescPrice.Default.(decimal.Decimal)
	// subscriptionDescDescription is the schema descriptor for description field.
	subscriptionDescDescription := subscriptionFields[3].Descriptor()
	// subscription.DefaultDescription holds the default value on creation for the description field.
	subscription.DefaultDescription = subscriptionDescDescription.Default.(string)
	// subscriptionDescSortOrder is the schema descriptor for sort_order field.
	subscriptionDescSortOrder := subscriptionFields[4].Descriptor()
	// subscription.DefaultSortOrder holds the default value on creation for the sort_order field.
	subscription.DefaultSortOrder = subscriptionDescSortOrder.Default.(func() uint32)
	// subscriptionDescPackageType is the schema descriptor for package_type field.
	subscriptionDescPackageType := subscriptionFields[5].Descriptor()
	// subscription.DefaultPackageType holds the default value on creation for the package_type field.
	subscription.DefaultPackageType = subscriptionDescPackageType.Default.(string)
	// subscriptionDescCredit is the schema descriptor for credit field.
	subscriptionDescCredit := subscriptionFields[6].Descriptor()
	// subscription.DefaultCredit holds the default value on creation for the credit field.
	subscription.DefaultCredit = subscriptionDescCredit.Default.(func() uint32)
	// subscriptionDescResetType is the schema descriptor for reset_type field.
	subscriptionDescResetType := subscriptionFields[7].Descriptor()
	// subscription.DefaultResetType holds the default value on creation for the reset_type field.
	subscription.DefaultResetType = subscriptionDescResetType.Default.(string)
	// subscriptionDescQPSLimit is the schema descriptor for qps_limit field.
	subscriptionDescQPSLimit := subscriptionFields[8].Descriptor()
	// subscription.DefaultQPSLimit holds the default value on creation for the qps_limit field.
	subscription.DefaultQPSLimit = subscriptionDescQPSLimit.Default.(func() uint32)
	usercreditrecordMixin := schema.UserCreditRecord{}.Mixin()
	usercreditrecord.Policy = privacy.NewPolicies(usercreditrecordMixin[0], schema.UserCreditRecord{})
	usercreditrecord.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := usercreditrecord.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	usercreditrecordMixinFields0 := usercreditrecordMixin[0].Fields()
	_ = usercreditrecordMixinFields0
	usercreditrecordMixinFields1 := usercreditrecordMixin[1].Fields()
	_ = usercreditrecordMixinFields1
	usercreditrecordFields := schema.UserCreditRecord{}.Fields()
	_ = usercreditrecordFields
	// usercreditrecordDescCreatedAt is the schema descriptor for created_at field.
	usercreditrecordDescCreatedAt := usercreditrecordMixinFields0[0].Descriptor()
	// usercreditrecord.DefaultCreatedAt holds the default value on creation for the created_at field.
	usercreditrecord.DefaultCreatedAt = usercreditrecordDescCreatedAt.Default.(func() uint32)
	// usercreditrecordDescUpdatedAt is the schema descriptor for updated_at field.
	usercreditrecordDescUpdatedAt := usercreditrecordMixinFields0[1].Descriptor()
	// usercreditrecord.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usercreditrecord.DefaultUpdatedAt = usercreditrecordDescUpdatedAt.Default.(func() uint32)
	// usercreditrecord.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usercreditrecord.UpdateDefaultUpdatedAt = usercreditrecordDescUpdatedAt.UpdateDefault.(func() uint32)
	// usercreditrecordDescDeletedAt is the schema descriptor for deleted_at field.
	usercreditrecordDescDeletedAt := usercreditrecordMixinFields0[2].Descriptor()
	// usercreditrecord.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	usercreditrecord.DefaultDeletedAt = usercreditrecordDescDeletedAt.Default.(func() uint32)
	// usercreditrecordDescEntID is the schema descriptor for ent_id field.
	usercreditrecordDescEntID := usercreditrecordMixinFields1[1].Descriptor()
	// usercreditrecord.DefaultEntID holds the default value on creation for the ent_id field.
	usercreditrecord.DefaultEntID = usercreditrecordDescEntID.Default.(func() uuid.UUID)
	// usercreditrecordDescAppID is the schema descriptor for app_id field.
	usercreditrecordDescAppID := usercreditrecordFields[0].Descriptor()
	// usercreditrecord.DefaultAppID holds the default value on creation for the app_id field.
	usercreditrecord.DefaultAppID = usercreditrecordDescAppID.Default.(func() uuid.UUID)
	// usercreditrecordDescUserID is the schema descriptor for user_id field.
	usercreditrecordDescUserID := usercreditrecordFields[1].Descriptor()
	// usercreditrecord.DefaultUserID holds the default value on creation for the user_id field.
	usercreditrecord.DefaultUserID = usercreditrecordDescUserID.Default.(func() uuid.UUID)
	// usercreditrecordDescOperationType is the schema descriptor for operation_type field.
	usercreditrecordDescOperationType := usercreditrecordFields[2].Descriptor()
	// usercreditrecord.DefaultOperationType holds the default value on creation for the operation_type field.
	usercreditrecord.DefaultOperationType = usercreditrecordDescOperationType.Default.(string)
	// usercreditrecordDescCreditsChange is the schema descriptor for credits_change field.
	usercreditrecordDescCreditsChange := usercreditrecordFields[3].Descriptor()
	// usercreditrecord.DefaultCreditsChange holds the default value on creation for the credits_change field.
	usercreditrecord.DefaultCreditsChange = usercreditrecordDescCreditsChange.Default.(uint32)
	// usercreditrecordDescExtra is the schema descriptor for extra field.
	usercreditrecordDescExtra := usercreditrecordFields[4].Descriptor()
	// usercreditrecord.DefaultExtra holds the default value on creation for the extra field.
	usercreditrecord.DefaultExtra = usercreditrecordDescExtra.Default.(string)
	usersubscriptionMixin := schema.UserSubscription{}.Mixin()
	usersubscription.Policy = privacy.NewPolicies(usersubscriptionMixin[0], schema.UserSubscription{})
	usersubscription.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := usersubscription.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	usersubscriptionMixinFields0 := usersubscriptionMixin[0].Fields()
	_ = usersubscriptionMixinFields0
	usersubscriptionMixinFields1 := usersubscriptionMixin[1].Fields()
	_ = usersubscriptionMixinFields1
	usersubscriptionFields := schema.UserSubscription{}.Fields()
	_ = usersubscriptionFields
	// usersubscriptionDescCreatedAt is the schema descriptor for created_at field.
	usersubscriptionDescCreatedAt := usersubscriptionMixinFields0[0].Descriptor()
	// usersubscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersubscription.DefaultCreatedAt = usersubscriptionDescCreatedAt.Default.(func() uint32)
	// usersubscriptionDescUpdatedAt is the schema descriptor for updated_at field.
	usersubscriptionDescUpdatedAt := usersubscriptionMixinFields0[1].Descriptor()
	// usersubscription.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usersubscription.DefaultUpdatedAt = usersubscriptionDescUpdatedAt.Default.(func() uint32)
	// usersubscription.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usersubscription.UpdateDefaultUpdatedAt = usersubscriptionDescUpdatedAt.UpdateDefault.(func() uint32)
	// usersubscriptionDescDeletedAt is the schema descriptor for deleted_at field.
	usersubscriptionDescDeletedAt := usersubscriptionMixinFields0[2].Descriptor()
	// usersubscription.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	usersubscription.DefaultDeletedAt = usersubscriptionDescDeletedAt.Default.(func() uint32)
	// usersubscriptionDescEntID is the schema descriptor for ent_id field.
	usersubscriptionDescEntID := usersubscriptionMixinFields1[1].Descriptor()
	// usersubscription.DefaultEntID holds the default value on creation for the ent_id field.
	usersubscription.DefaultEntID = usersubscriptionDescEntID.Default.(func() uuid.UUID)
	// usersubscriptionDescAppID is the schema descriptor for app_id field.
	usersubscriptionDescAppID := usersubscriptionFields[0].Descriptor()
	// usersubscription.DefaultAppID holds the default value on creation for the app_id field.
	usersubscription.DefaultAppID = usersubscriptionDescAppID.Default.(func() uuid.UUID)
	// usersubscriptionDescUserID is the schema descriptor for user_id field.
	usersubscriptionDescUserID := usersubscriptionFields[1].Descriptor()
	// usersubscription.DefaultUserID holds the default value on creation for the user_id field.
	usersubscription.DefaultUserID = usersubscriptionDescUserID.Default.(func() uuid.UUID)
	// usersubscriptionDescPackageID is the schema descriptor for package_id field.
	usersubscriptionDescPackageID := usersubscriptionFields[2].Descriptor()
	// usersubscription.DefaultPackageID holds the default value on creation for the package_id field.
	usersubscription.DefaultPackageID = usersubscriptionDescPackageID.Default.(func() uuid.UUID)
	// usersubscriptionDescOrderID is the schema descriptor for order_id field.
	usersubscriptionDescOrderID := usersubscriptionFields[3].Descriptor()
	// usersubscription.DefaultOrderID holds the default value on creation for the order_id field.
	usersubscription.DefaultOrderID = usersubscriptionDescOrderID.Default.(func() uuid.UUID)
	// usersubscriptionDescStartAt is the schema descriptor for start_at field.
	usersubscriptionDescStartAt := usersubscriptionFields[4].Descriptor()
	// usersubscription.DefaultStartAt holds the default value on creation for the start_at field.
	usersubscription.DefaultStartAt = usersubscriptionDescStartAt.Default.(uint32)
	// usersubscriptionDescEndAt is the schema descriptor for end_at field.
	usersubscriptionDescEndAt := usersubscriptionFields[5].Descriptor()
	// usersubscription.DefaultEndAt holds the default value on creation for the end_at field.
	usersubscription.DefaultEndAt = usersubscriptionDescEndAt.Default.(uint32)
	// usersubscriptionDescUsageState is the schema descriptor for usage_state field.
	usersubscriptionDescUsageState := usersubscriptionFields[6].Descriptor()
	// usersubscription.DefaultUsageState holds the default value on creation for the usage_state field.
	usersubscription.DefaultUsageState = usersubscriptionDescUsageState.Default.(string)
	// usersubscriptionDescSubscriptionCredit is the schema descriptor for subscription_credit field.
	usersubscriptionDescSubscriptionCredit := usersubscriptionFields[7].Descriptor()
	// usersubscription.DefaultSubscriptionCredit holds the default value on creation for the subscription_credit field.
	usersubscription.DefaultSubscriptionCredit = usersubscriptionDescSubscriptionCredit.Default.(uint32)
	// usersubscriptionDescAddonCredit is the schema descriptor for addon_credit field.
	usersubscriptionDescAddonCredit := usersubscriptionFields[8].Descriptor()
	// usersubscription.DefaultAddonCredit holds the default value on creation for the addon_credit field.
	usersubscription.DefaultAddonCredit = usersubscriptionDescAddonCredit.Default.(uint32)
}

const (
	Version = "v0.12.0" // Version of ent codegen.
)
