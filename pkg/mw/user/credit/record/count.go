package record

import (
	"context"

	"github.com/NpoolPlatform/billing-middleware/pkg/db"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
)

type countHandler struct {
	*baseQueryHandler
	total uint32
}

func (h *Handler) CountRecords(ctx context.Context) (uint32, error) {
	handler := &countHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err = handler.queryRecords(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_total, err := handler.stmSelect.Count(_ctx)
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
