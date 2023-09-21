FROM golang:1.21 as build

#ENV WORKING_DIR=/src
#ENV TARGET_DIR=/app
#ENV EXEC=vue-register
#ENV CONFIG_DIR=config
#ENV CONFIG_FILE=config.json

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./

# See https://mt165.co.uk/blog/static-link-go/
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'"

FROM scratch
RUN mkdir /app/tokens
COPY --from=build /src/.env /app/.env
COPY --from=build /src/config.json /app/config.json
COPY --from=build /src/tokens/* /app/tokens
COPY --from=build /src/vue-register /app/vue-register

EXPOSE 9000
CMD ["cd", "/app", "&&", "./vue-register", "serve"]
