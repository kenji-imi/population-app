.PHONY: init
init:
	go get -v -u golang.org/x/tools/cmd/goimports
	go get -v -u github.com/knq/xo

.PHONY: dep
dep:
	go get -v -u github.com/golang/dep/cmd/dep
	dep ensure -v

.PHONY: gen_xo_models
gen_xo_models:
	mkdir -p ./src/xo
	xo mysql://root:root@localhost:3309/populationdb -o ./src/xo

.PHONY: regen_xo_models
regen_xo_models:
	rm -f ./src/xo/*.xo.go ./src/xo/xo_db.xo.go
	xo mysql://root:root@localhost:3309/populationdb -o ./src/xo

.PHONY: clean
clean:
	echo "clean"

.PHONY: start_docker
start_docker:
	docker-compose up -d

.PHONY: stop_docker
stop_docker:
	docker-compose stop

.PHONY: down_docker
down_docker:
	docker-compose down -v
