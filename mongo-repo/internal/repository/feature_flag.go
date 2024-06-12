package repository

import (
	"context"

	"github.com/novychok/go-samples/mongorepo/internal/entity"
)

type FeatureFlag interface {
	FindAll(ctx context.Context) ([]*entity.FeatureFlag, error)
	Create(ctx context.Context, create *entity.FeatureFlag) (entity.FeatureFlagID, error)
	Update(ctx context.Context, id entity.FeatureFlagID, update *entity.FeatureFlagUpsert) (*entity.FeatureFlag, error)
	Delete(ctx context.Context, id entity.FeatureFlagID) error
	GetByID(ctx context.Context, id entity.FeatureFlagID) (*entity.FeatureFlag, error)
	UpdateStatus(ctx context.Context, id entity.FeatureFlagID, toggle *entity.Toggle) error
}
