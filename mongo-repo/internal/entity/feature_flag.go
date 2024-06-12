package entity

type StatusAction string

type FeatureFlagID ID

type ID string

func (id ID) FeatureObjectID() string {
	return string(id)
}

const (
	Enable StatusAction = "enabled"
	Diable StatusAction = "disabled"
)

type FeatureFlag struct {
	ID          FeatureFlagID
	Name        string
	Description string
	Value       string
	Status      StatusAction
	Metadata    string
}

type FeatureFlagUpsert struct {
	Name        string
	Description string
	Value       string
	Status      StatusAction
	Metadata    string
}

type Toggle struct {
	Status StatusAction
}
