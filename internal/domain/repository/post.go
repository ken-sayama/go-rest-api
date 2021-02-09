package repository

import (
	"context"

	"github.com/ken-sayama/todo-app/internal/domain/entity"
)

type Post interface {
	Get(ctx context.Context, id int32) (entity.Post, error)
}
