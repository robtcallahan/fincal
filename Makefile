DOCKER_USERNAME ?= robtcallahan
FRONTEND_NAME ?= fincal-frontend
BACKEND_NAME ?= fincal-backend
DATABASE_NAME ?= fincal-database

.PHONY: frontend

up:
	docker-compose up

down:
	docker-compose stop

frontend:
	cd ./services/frontend && docker build --tag ${FRONTEND_NAME} .

backend:
	cd ./services/frontend && docker build --tag ${BACKEND_NAME} .

database:
	cd ./services/database && docker build --tag ${DATABASE_NAME} .
