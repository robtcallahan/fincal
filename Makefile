DOCKER_USERNAME ?= robtcallahan
BACKEND_NAME ?= fincal-backend
FRONTEND_NAME ?= fincal-frontend

.PHONY: frontend

up:
	docker-compose up

down:
	docker-compose stop

frontend:
	cd ./frontend && docker build --tag ${FRONTEND_NAME} .

backend:
	docker build --tag ${BACKEND_NAME} .

server_backend:
	docker run -p 9000:9000${BACKEND_NAME}

server_frontend:
	docker run -p 8000:5173 ${FRONTEND_NAME}