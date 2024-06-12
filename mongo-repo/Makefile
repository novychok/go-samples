.PHONY: gen
mock:
	mockgen -destination=internal/repository/feature_flag/mocks/mock_feature_flag.go \
		-package=feature_flag github.com/novychok/go-samples/mongorepo/internal/repository FeatureFlag

test:
	@go test -v ./...