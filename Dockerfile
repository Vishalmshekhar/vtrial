FROM golang:1.19

LABEL Vishal Local Test

WORKDIR /projects/data/

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o se.exe ./cmd/main.go

EXPOSE 4000

CMD [ "./se.exe" ]
