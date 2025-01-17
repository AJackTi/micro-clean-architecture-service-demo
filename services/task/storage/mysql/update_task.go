package mysql

import (
	"context"
	"demo-service/services/task/entity"
	"github.com/pkg/errors"
)

func (store *mysqlStore) UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error {
	if err := store.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
