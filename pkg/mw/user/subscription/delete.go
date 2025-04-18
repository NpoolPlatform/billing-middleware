package subscription

import (
	"context"
	"time"

	subscriptioncrud "github.com/NpoolPlatform/billing-middleware/pkg/crud/user/subscription"
	"github.com/NpoolPlatform/billing-middleware/pkg/db"
	"github.com/NpoolPlatform/billing-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteSubscription(ctx context.Context, cli *ent.Client) error {
	if _, err := subscriptioncrud.UpdateSet(
		cli.UserSubscription.UpdateOneID(*h.ID),
		&subscriptioncrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSubscription(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetSubscription(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteSubscription(_ctx, cli)
	})
}
