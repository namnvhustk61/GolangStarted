package biz

import (
	"context"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoaitemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoaitemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank
	}
	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
