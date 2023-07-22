default: testacc
.PHONY: clean docker-down docker-up testacc test-acceptance

clean: docker-down

# Pulls down the compose stack removing the ephemeral volumes as well.
docker-down:
	cd docker; docker compose down -v

# Brings up the compose stack in daemon mode.
docker-up:
	cd docker; docker compose up -d

# Brings up the compose stack then runs the acceptance tests.
testacc: clean docker-up test-acceptance clean

# Run acceptance tests
test-acceptance:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
