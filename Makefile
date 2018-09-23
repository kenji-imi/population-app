MYSQL_HOST=127.0.0.1
MYSQL_PORT=3309
MYSQL_USER=root
MYSQL_PASS=root

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
	rm -f ./src/xo/xo_db.xo.go
	rm -f ./src/xo/*.xo.go
	xo mysql://root:root@localhost:3309/populationdb -o ./src/xo

.PHONY: gen_plantuml
gen_plantuml:
	mysql -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) --password=$(MYSQL_PASS) -e "drop database if exists dump_schema"
	mysql -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) --password=$(MYSQL_PASS) -e "create database dump_schema"
	cat db/init/1_create_table.sql | grep -v -E '^USE ' | mysql -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) --password=$(MYSQL_PASS) dump_schema
	rm -f db/plantuml/table_class_definitions.puml
	xo mysql://$(MYSQL_USER):$(MYSQL_PASS)@$(MYSQL_HOST):$(MYSQL_PORT)/dump_schema --int32-type int32 --uint32-type uint32 --template-path xo/plantuml_templates -o db/plantuml/table_class_definitions.puml --single-file 2>&1 >/dev/null | true

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
