FROM golang:latest

WORKDIR /opt/cah

COPY . .

CMD ["go", "build"]

ENTRYPOINT ["./cah"]
