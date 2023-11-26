FROM golang

WORKDIR /app

ARG CACHEBUST=1

RUN git clone -b develop https://naikit20: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@github.com/naikit20/https://github.com/NaiKit20/GORM_Hex.git .

RUN go mod download

ENV GIN_MODE=release

COPY ./ ./

RUN go install github.com/cosmtrek/air@latest

# Don't forget to add .air.toml .gitignore
RUN air init

EXPOSE 8080

CMD ["air"]