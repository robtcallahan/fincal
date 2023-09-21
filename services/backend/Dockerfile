FROM golang:1.21 as build

#ENV WORKING_DIR=/src
#ENV TARGET_DIR=/app
#ENV EXEC=vue-register
#ENV CONFIG_FILE=config.json
#ENV TOKENS_DIR=tokens

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./

# See https://mt165.co.uk/blog/static-link-go/
RUN go build -ldflags "-linkmode 'external' -extldflags '-static'"

FROM scratch
WORKDIR /app/backend
COPY --from=build /src/.env .
COPY --from=build /src/config.json .
COPY --from=build /src/tokens/* ./tokens
COPY --from=build /src/vue-register .

EXPOSE 9000
CMD ["./vue-register", "serve"]
