package subscription

import (
	"entgo.io/ent/dialect/sql"

	subscriptioncrud "github.com/NpoolPlatform/billing-middleware/pkg/crud/subscription"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	entsubscription "github.com/NpoolPlatform/billing-middleware/pkg/db/ent/subscription"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.SubscriptionSelect
}

func (h *baseQueryHandler) selectSubscription(stm *ent.SubscriptionQuery) *ent.SubscriptionSelect {
	return stm.Select(entsubscription.FieldID)
}

func (h *baseQueryHandler) querySubscription(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.Subscription.Query().Where(entsubscription.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entsubscription.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entsubscription.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSubscription(stm)
	return nil
}

func (h *baseQueryHandler) querySubscriptions(cli *ent.Client) (*ent.SubscriptionSelect, error) {
	stm, err := subscriptioncrud.SetQueryConds(cli.Subscription.Query(), h.SubscriptionConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectSubscription(stm), nil
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
			t1.C(entsubscription.FieldAppID),
			t1.C(entsubscription.FieldPackageName),
			t1.C(entsubscription.FieldPrice),
			t1.C(entsubscription.FieldDescription),
			t1.C(entsubscription.FieldSortOrder),
			t1.C(entsubscription.FieldPackageType),
			t1.C(entsubscription.FieldCredit),
			t1.C(entsubscription.FieldResetType),
			t1.C(entsubscription.FieldQPSLimit),
			t1.C(entsubscription.FieldCreatedAt),
			t1.C(entsubscription.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
