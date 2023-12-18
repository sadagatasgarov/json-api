FROM golang:1.21.5-alpine3.18 as builder

WORKDIR /app

COPY . .

#RUN go mod download

COPY *.go .

RUN go build -o ./app/gobank
CMD ["chmod", "+x", "./app/gobank" ]


FROM scratch

COPY --from=builder /app/app/gobank /app/gobank

EXPOSE 4000


CMD [ "./app/gobank" ]