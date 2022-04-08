FROM golang:latest

WORKDIR /app

COPY ./cats-service/go.mod ./
COPY ./cats-service/go.sum ./

RUN go mod download

COPY ./cats-service/*.go ./

RUN go build -o /docker-cats
COPY ./waitforit.sh .
RUN chmod +x wait-for-it.sh
EXPOSE 8080


CMD ./wait-for-it.sh ${MYSQL_HOST}:${MYSQL_PORT} -- /docker-cats