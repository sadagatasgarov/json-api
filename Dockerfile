FROM golang:1.21-alpine3.17 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

COPY *.go .

RUN go build -o ./app/gobank
CMD ["chmod", "+x", "./app/gobank" ]


FROM scratch

COPY --from=builder /app/app/gobank /app/gobank

EXPOSE 3000


CMD [ "./app/gobank" ]