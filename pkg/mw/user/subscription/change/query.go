package change

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/billing-middleware/pkg/db"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/user/subscription/change"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.UserSubscriptionChangeSelect
	infos    []*npool.SubscriptionChange
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

func (h *Handler) GetSubscriptionChange(ctx context.Context) (*npool.SubscriptionChange, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySubscriptionChange(cli); err != nil {
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
	return handler.infos[0], nil
}

func (h *Handler) GetSubscriptionChanges(ctx context.Context) ([]*npool.SubscriptionChange, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.querySubscriptionChanges(cli); err != nil {
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
	return handler.infos, nil
}

func (h *Handler) GetSubscriptionChangesCount(ctx context.Context) (uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmCount, err = handler.querySubscriptionChanges(cli); err != nil {
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
