# Go parameters
GOCMD=go

.SILENT:
run:
		$(GOCMD) run cmd/sudoku/sudoku.go

test:
		$(GOCMD) test ./... -cover -v
