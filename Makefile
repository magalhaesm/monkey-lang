RESET := \033[0m
CYAN  := \033[1;36m
CHECK := \342\234\224
LOG   := printf "[$(CYAN)$(CHECK)$(RESET)] %s\n"

all:
	@$(LOG) "Running the application..."
	@go run .

test:
	@$(LOG) "Running all tests..."
	@go test ./... -v
