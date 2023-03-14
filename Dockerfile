FROM golang:1.19
WORKDIR /app

COPY docs ./docs
COPY handlers ./handlers
COPY services ./services
COPY static ./static
COPY models ./models
COPY .env ./
COPY app.go ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go build -o /mbase

EXPOSE 3000

CMD [ "/mbase" ]