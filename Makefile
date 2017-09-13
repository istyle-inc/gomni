#! /usr/bin/make
test:
	@go test -tags=test gomni/... -cover
