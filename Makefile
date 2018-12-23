# Go parameters
GOCMD=go

.SILENT:
install:
		dep ensure

local:
		$(GOCMD) run cmd/sudoku/sudoku.go

test:
		$(GOCMD) test ./... -cover -v -coverprofile=coverage.out

cover:
		$(GOCMD) tool cover -html=coverage.out
