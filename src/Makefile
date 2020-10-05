GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

# Shows the variables we just set.
# By prepending `@` we prevent Make
# from printing the command before the
# stdout of the execution.
show:
	@echo "SRC  = $(GO_SRC_DIRS)"
	@echo "TEST = $(GO_TEST_DIRS)"

test: $(GO_TEST_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go test -v ; \
		popd > /dev/null ; \
	done;

fmt: $(GO_SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		echo $$dir ; \
		echo $(MAKE)  ; \
		echo $(LD)  ; \
		echo $(CC)  ; \
		echo $(@)  ; \
		echo $(%)  ; \
		echo $(?)  ; \
		popd > /dev/null ; \
	done;

generate.c: test.o
	@echo make generate

test.o: test.c
	@echo make test