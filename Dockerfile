FROM golang:buster AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# https://johnnn.tech/q/very-slow-docker-build-with-go-sqlite3-cgo-enabled-package/
RUN go install github.com/mattn/go-sqlite3

COPY . ./
RUN make build-backend


FROM node:16 AS frontend

WORKDIR /app
COPY front ./

RUN npm ci
RUN npm run build


FROM python:3.8-buster

RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    net-tools \
    sqlite3
RUN pip3 install litecli

WORKDIR /app

COPY ./worker ./worker
RUN pip3 install -r ./worker/requirements.txt

COPY --from=backend /app/evexp ./evexp
COPY --from=frontend /app/dist ./front/dist

RUN ln -s /app/evexp /usr/bin/evexp

CMD ["evexp"]
