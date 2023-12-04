.PHONY: docker-build
docker-build:
	@docker build -t codegen .

.PHONY: docker-run
docker-run:
	@docker run -p 8080:8080 codegen