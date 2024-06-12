package dao

import (
	"github.com/novychok/go-samples/mongorepo/internal/entity"
)

type FeatureFlagDAO struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Value       string `bson:"value"`
	Status      string `bson:"status"`
	Metadata    string `bson:"metadata"`
}

func ToDAO(featureFlag *entity.FeatureFlag) (*FeatureFlagDAO, error) {
	return &FeatureFlagDAO{
		ID:          string(featureFlag.ID),
		Name:        featureFlag.Name,
		Description: featureFlag.Description,
		Value:       featureFlag.Value,
		Status:      string(featureFlag.Status),
		Metadata:    featureFlag.Metadata,
	}, nil
}

func ToEntity(featureFlagDAO *FeatureFlagDAO) (*entity.FeatureFlag, error) {
	return &entity.FeatureFlag{
		ID:          entity.FeatureFlagID(featureFlagDAO.ID),
		Name:        featureFlagDAO.Name,
		Description: featureFlagDAO.Description,
		Value:       featureFlagDAO.Value,
		Status:      entity.StatusAction(featureFlagDAO.Status),
		Metadata:    featureFlagDAO.Metadata,
	}, nil
}
