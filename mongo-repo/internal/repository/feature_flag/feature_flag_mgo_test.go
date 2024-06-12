package feature_flag_test

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/novychok/go-samples/mongorepo/internal/entity"
	mocks "github.com/novychok/go-samples/mongorepo/internal/repository/feature_flag/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var expectedFlags = []*entity.FeatureFlag{
	{
		ID:          "1",
		Name:        "Feature1",
		Description: "This is feature 1",
		Value:       "true",
		Status:      "enabled",
		Metadata:    "metadata for feature 1",
	},
	{
		ID:          "2",
		Name:        "Feature2",
		Description: "This is feature 2",
		Value:       "42",
		Status:      "disabled",
		Metadata:    "metadata for feature 2",
	},
	{
		ID:          "3",
		Name:        "Feature3",
		Description: "This is feature 3",
		Value:       "value for feature 3",
		Status:      "disabled",
		Metadata:    "metadata for feature 3",
	},
}

var flag = &entity.FeatureFlag{
	ID:          "4",
	Name:        "Feature3",
	Description: "This is feature 3",
	Value:       "value for feature 3",
	Status:      "enabled",
	Metadata:    "metadata for feature 3",
}

func Test_FindAll(t *testing.T) {

	ctx := context.Background()

	t.Run("Find All Feature Flags", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().FindAll(ctx).Return(expectedFlags, nil)

		featureFlags, err := mockRepo.FindAll(ctx)

		var got string
		for _, flag := range expectedFlags {
			bytes, err := json.Marshal(flag)
			if err != nil {
				log.Println(err)
			}

			got += string(bytes)
		}

		var result string
		for _, flag := range featureFlags {
			bytes, err := json.Marshal(flag)
			if err != nil {
				log.Println(err)
			}

			result += string(bytes)
		}

		gotHash := md5.Sum([]byte(got))
		resultHash := md5.Sum([]byte(result))

		assert.NoError(t, err)
		assert.Equal(t, gotHash, resultHash)
	})
}

func Test_Create(t *testing.T) {

	ctx := context.Background()

	t.Run("Create Feature Flag", func(t *testing.T) {
		for _, flag := range expectedFlags {
			ctrl := gomock.NewController(t)
			mockRepo := mocks.NewMockFeatureFlag(ctrl)
			mockRepo.EXPECT().Create(ctx, flag).Return(flag.ID, nil)

			result, err := mockRepo.Create(ctx, flag)

			assert.NoError(t, err)
			assert.Equal(t, flag.ID, result)
		}
	})
}

func Test_Delete(t *testing.T) {

	ctx := context.Background()

	t.Run("Delete Feature Flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().Delete(ctx, flag.ID).Return(nil)

		err := mockRepo.Delete(ctx, flag.ID)
		assert.NoError(t, err)
	})
}

func Test_GetByID(t *testing.T) {

	ctx := context.Background()

	t.Run("Get By ID Feature Flag", func(t *testing.T) {
		for _, flag := range expectedFlags {
			ctrl := gomock.NewController(t)
			mockRepo := mocks.NewMockFeatureFlag(ctrl)
			mockRepo.EXPECT().GetByID(ctx, flag.ID).Return(flag, nil)

			result, err := mockRepo.GetByID(ctx, flag.ID)

			assert.NoError(t, err)
			assert.Equal(t, flag, result)
			assert.Equal(t, flag.Name, result.Name)
			assert.Equal(t, flag.Description, result.Description)
			assert.Equal(t, flag.Value, result.Value)
			assert.Equal(t, flag.Status, result.Status)
			assert.Equal(t, flag.Metadata, result.Metadata)
		}
	})

	t.Run("Feature Flag Not Found", func(t *testing.T) {
		idToCheck := entity.FeatureFlagID("4")

		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().GetByID(ctx, idToCheck).Return(nil, mongo.ErrNoDocuments)

		result, err := mockRepo.GetByID(ctx, idToCheck)

		assert.Error(t, err)
		assert.Equal(t, mongo.ErrNoDocuments, err)
		assert.Nil(t, result)
	})
}

func Test_Update(t *testing.T) {

	update := &entity.FeatureFlagUpsert{
		Name:        "chicki bamboni",
		Description: "burger shop",
		Value:       "100",
		Status:      "disabled",
		Metadata:    "cucumber",
	}

	expectedFlag := &entity.FeatureFlag{
		ID:          flag.ID,
		Name:        "chicki bamboni",
		Description: "burger shop",
		Value:       "100",
		Status:      "disabled",
		Metadata:    "cucumber",
	}

	ctx := context.Background()

	t.Run("Update Feature Flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().Update(ctx, flag.ID, update).Return(expectedFlag, nil)

		result, err := mockRepo.Update(ctx, flag.ID, update)

		assert.NoError(t, err)
		assert.Equal(t, expectedFlag, result)
		assert.Equal(t, expectedFlag.Name, result.Name)
		assert.Equal(t, expectedFlag.Description, result.Description)
		assert.Equal(t, expectedFlag.Value, result.Value)
		assert.Equal(t, expectedFlag.Status, result.Status)
		assert.Equal(t, expectedFlag.Metadata, result.Metadata)

	})

	t.Run("Feature Flag Not Found For Update", func(t *testing.T) {
		idToCheck := entity.FeatureFlagID("5")

		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().Update(ctx, idToCheck, update).Return(nil, mongo.ErrNoDocuments)

		result, err := mockRepo.Update(ctx, idToCheck, update)

		assert.Error(t, err)
		assert.Equal(t, mongo.ErrNoDocuments, err)
		assert.Nil(t, result)
	})
}

func Test_UpdateStatus(t *testing.T) {

	ctx := context.Background()

	toggle := &entity.Toggle{
		Status: "disabled",
	}

	t.Run("Update Status Feature Flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().UpdateStatus(ctx, flag.ID, toggle).Return(nil)

		err := mockRepo.UpdateStatus(ctx, flag.ID, toggle)

		assert.NoError(t, err)
	})

	t.Run("Feature Flag Not Found For Update Status", func(t *testing.T) {
		idToCheck := entity.FeatureFlagID("5")

		ctrl := gomock.NewController(t)
		mockRepo := mocks.NewMockFeatureFlag(ctrl)
		mockRepo.EXPECT().UpdateStatus(ctx, idToCheck, toggle).Return(mongo.ErrNoDocuments)

		err := mockRepo.UpdateStatus(ctx, idToCheck, toggle)

		assert.Error(t, err)
		assert.Equal(t, mongo.ErrNoDocuments, err)
	})
}
