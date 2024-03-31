DOCKER_TAG_FEATURE_A=feature-a:latest
DOCKER_TAG_FEATURE_B=feature-b:latest

.PHONY: build-feature-a

# feature-a の Docker イメージをビルド
docker-build-feature-a:
	docker build -t $(DOCKER_TAG_FEATURE_A) -f cmd/saas-a/feature-a/Dockerfile .
