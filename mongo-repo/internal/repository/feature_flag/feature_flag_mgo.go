package feature_flag

import (
	"context"

	"github.com/novychok/go-samples/mongorepo/internal/dao"
	"github.com/novychok/go-samples/mongorepo/internal/entity"
	"github.com/novychok/go-samples/mongorepo/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName = "feature_flags"

type mgo struct {
	db *mongo.Database
}

func (r *mgo) FindAll(ctx context.Context) ([]*entity.FeatureFlag, error) {

	cursor, err := r.db.Collection(collectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var featureFlags []*entity.FeatureFlag
	for cursor.Next(ctx) {
		var featureFlagDAO dao.FeatureFlagDAO
		if err := cursor.Decode(&featureFlagDAO); err != nil {
			return nil, err
		}

		result, err := dao.ToEntity(&featureFlagDAO)
		if err != nil {
			return nil, err
		}

		featureFlags = append(featureFlags, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return featureFlags, nil
}

func (r *mgo) Create(ctx context.Context, create *entity.FeatureFlag) (entity.FeatureFlagID, error) {

	featureFlagDAO, err := dao.ToDAO(create)
	if err != nil {
		return "", err
	}

	result, err := r.db.Collection(collectionName).InsertOne(ctx, featureFlagDAO)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(entity.FeatureFlagID), nil
}

func (r *mgo) Update(ctx context.Context, id entity.FeatureFlagID, update *entity.FeatureFlagUpsert) (*entity.FeatureFlag, error) {

	var updatedFeatureFlag dao.FeatureFlagDAO

	updateFlag := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: update.Name},
			{Key: "description", Value: update.Description},
			{Key: "value", Value: update.Value},
			{Key: "status", Value: update.Status},
			{Key: "metadata", Value: update.Metadata},
		}},
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.db.Collection(collectionName).FindOneAndUpdate(ctx, bson.M{"_id": id}, updateFlag, options).Decode(&updatedFeatureFlag)
	if err != nil {
		return nil, err
	}

	featureFlag, err := dao.ToEntity(&updatedFeatureFlag)
	if err != nil {
		return nil, err
	}

	return featureFlag, nil
}

func (r *mgo) Delete(ctx context.Context, id entity.FeatureFlagID) error {

	_, err := r.db.Collection(collectionName).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (r *mgo) GetByID(ctx context.Context, id entity.FeatureFlagID) (*entity.FeatureFlag, error) {

	var featureFlagDAO dao.FeatureFlagDAO

	err := r.db.Collection(collectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&featureFlagDAO)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	featureFlag, err := dao.ToEntity(&featureFlagDAO)
	if err != nil {
		return nil, err
	}

	return featureFlag, nil
}

func (r *mgo) UpdateStatus(ctx context.Context, id entity.FeatureFlagID, toggle *entity.Toggle) error {

	var updatedFeatureFlagStatus dao.FeatureFlagDAO

	updateStatus := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: toggle.Status},
		}},
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := r.db.Collection(collectionName).FindOneAndUpdate(ctx, bson.M{"_id": id}, updateStatus, options).Decode(&updatedFeatureFlagStatus)
	if err != nil {
		return err
	}
	return nil
}

func NewMongo(db *mongo.Database) repository.FeatureFlag {
	return &mgo{db: db}
}
