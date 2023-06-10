name = tl

.PHONY: build
build: main.go
	@echo "**** build start ****"
	make clean && go build -o $(name) $^
	@echo "****  build end  ****"

.PHONY: prod
prod: main.go
	@echo "production build start"
	go build -o $(name) -ldflags '-s -w' $^
	@echo "production build end"

.PHONY: clean
clean:
	@echo "**** start cleaning ****"
	go clean -n -x && rm -rf $(name).exe $(name)
	@echo "**** cleaning end  ****"
