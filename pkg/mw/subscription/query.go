package subscription

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/billing-middleware/pkg/db"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/billing/v1"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/subscription"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.SubscriptionSelect
	infos    []*npool.Subscription
	total    uint32
}

func (h *queryHandler) queryJoin() {
	if h.stmSelect != nil {
		h.baseQueryHandler.queryJoin()
		return
	}
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, _ := decimal.NewFromString(info.Price)
		info.Price = amount.String()
		info.PackageType = types.PackageType(types.PackageType_value[info.PackageTypeStr])
		info.ResetType = types.ResetType(types.ResetType_value[info.ResetTypeStr])
	}
}

func (h *Handler) GetSubscription(ctx context.Context) (*npool.Subscription, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySubscription(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetSubscriptions(ctx context.Context) ([]*npool.Subscription, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.querySubscriptions(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	handler.formalize()
	return handler.infos, nil
}

func (h *Handler) GetSubscriptionsCount(ctx context.Context) (uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmCount, err = handler.querySubscriptions(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		return nil
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return handler.total, nil
}
