
usecase_path=internal/usecase
handler_path=internal/handler


.PHONY: build-mocks
build-mocks:
	rm -rf ./**/mocks
	mockery --dir=$(usecase_path) --name=InfoRepository --filename=infos.go --output=$(usecase_path)/mocks --outpkg=mocks
	mockery --dir=$(usecase_path) --name=TechnologyRepository --filename=technology.go --output=$(usecase_path)/mocks --outpkg=mocks

	mockery --dir=$(handler_path) --name=UpsertTechnologyUseCase --filename=upsert_technology.go --output=$(handler_path)/mocks --outpkg=mocks
	mockery --dir=$(handler_path) --name=FindTechnologyUseCase --filename=find_technology.go --output=$(handler_path)/mocks --outpkg=mocks
	mockery --dir=$(handler_path) --name=InfosUseCase --filename=infos.go --output=$(handler_path)/mocks --outpkg=mocks


.PHONY: run-test-coverage
run-test-coverage:
	go test ./...  -coverprofile=coverage.out