package subscription

import (
	"entgo.io/ent/dialect/sql"

	subscriptioncrud "github.com/NpoolPlatform/billing-middleware/pkg/crud/user/subscription"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	entsubscription "github.com/NpoolPlatform/billing-middleware/pkg/db/ent/usersubscription"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.UserSubscriptionSelect
}

func (h *queryHandler) queryJoin() {
	if h.stmSelect != nil {
		h.baseQueryHandler.queryJoin()
	}
}

func (h *baseQueryHandler) selectSubscription(stm *ent.UserSubscriptionQuery) *ent.UserSubscriptionSelect {
	return stm.Select(entsubscription.FieldID)
}

func (h *baseQueryHandler) querySubscription(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.UserSubscription.Query().Where(entsubscription.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsubscription.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsubscription.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSubscription(stm)
	return nil
}

func (h *baseQueryHandler) querySubscriptions(cli *ent.Client) error {
	stm, err := subscriptioncrud.SetQueryConds(cli.UserSubscription.Query(), h.SubscriptionConds)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.stmSelect = h.selectSubscription(stm)
	return nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entsubscription.Table)
	s.Join(t1).
		On(
			s.C(entsubscription.FieldID),
			t1.C(entsubscription.FieldID),
		).
		AppendSelect(
			t1.C(entsubscription.FieldEntID),
			t1.C(entsubscription.FieldAppID),
			t1.C(entsubscription.FieldUserID),
			t1.C(entsubscription.FieldPackageID),
			t1.C(entsubscription.FieldStartAt),
			t1.C(entsubscription.FieldEndAt),
			t1.C(entsubscription.FieldUsageState),
			t1.C(entsubscription.FieldSubscriptionCredit),
			t1.C(entsubscription.FieldAddonCredit),
			t1.C(entsubscription.FieldCreatedAt),
			t1.C(entsubscription.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
