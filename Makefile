GO_MODULE := proto-course

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen
else
	rm -fR ./protogen
endif

.PHONY: protoc
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \
	./proto/first/*.proto \
	./proto/jobsearch/*.proto \

.PHONY: build
build: clean protoc tidy

.PHONY: run
run:
	@go run *.go

.PHONY: execute
execute: build run
