package exchange

import (
	"context"

	"github.com/NpoolPlatform/billing-middleware/pkg/db"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/billing/mw/v1/credit/exchange"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.ExchangeSelect
	infos    []*npool.Exchange
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetExchange(ctx context.Context) (*npool.Exchange, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryExchange(cli); err != nil {
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

func (h *Handler) GetExchanges(ctx context.Context) ([]*npool.Exchange, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryExchanges(cli); err != nil {
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

func (h *Handler) GetExchangesCount(ctx context.Context) (uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmCount, err = handler.queryExchanges(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		return handler.scan(_ctx)
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return handler.total, nil
}
