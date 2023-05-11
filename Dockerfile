FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN git clone https://github.com/CharlesMuchogo/point_of_sale.git
WORKDIR /build/pos
COPY .env .
COPY go.mod .
RUN go build -o main
EXPOSE 8080
ENTRYPOINT [ "/build/pos/main" ]

