package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	itemStatusDeleted := model.ItemStatusDeleted
	if err := s.db.Table(model.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": itemStatusDeleted.String(),
		}).Error; err != nil {
		return err
	}
	return nil
}
