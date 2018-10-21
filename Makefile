# Go parameters
GOCMD=go

.SILENT:
install:
		dep ensure

run:
		$(GOCMD) run cmd/sudoku/sudoku.go

test:
		$(GOCMD) test ./... -cover -v
