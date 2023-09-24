DOCKER_USERNAME ?= robtcallahan
FRONTEND_NAME ?= fincal-frontend
BACKEND_NAME ?= fincal-backend
DATABASE_NAME ?= fincal-database

ecr_base := 941928517230.dkr.ecr.us-east-1.amazonaws.com
services = fincal-frontend fincal-backend fincal-database

define newline

endef

.PHONY: frontend
.SILENT: ecr-push prune list

run: build up
runnc: buildnc upnc

try:

build:
	docker-compose build

up:
	docker-compose up

stop:
	docker-compose down

buildnc:
	docker-compose build --no-cache

upnc:
	docker-compose up --force-recreate

down:
	docker-compose stop

frontend:
	cd ./services/frontend && docker build --tag ${FRONTEND_NAME} .

backend:
	cd ./services/backend && docker build --tag ${BACKEND_NAME} .

database:
	cd ./services/database && docker build --tag ${DATABASE_NAME} .

ecr-push:
	echo fincal-database ; \
	docker tag fincal-database $(ecr_base)/fincal-database ; \
	docker push $(ecr_base)/fincal-database ; \
#	for service in $(services); do \
#		echo $$service ; \
#		docker tag $$service $(ecr_base)/$$service ; \
#		docker push $(ecr_base)/$$service ; \
#	done
#	$(foreach service,$(services),echo docker push $(ecr_base)/$(service); $(newline))

list:
	docker image ls
	echo ""
	docker container ls -a
	echo ""
	docker volume ls

prune:
	docker container prune --force
	echo ""
	docker image prune --force
	echo ""
	docker volume rm -f fincal_db-data