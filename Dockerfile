FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN git clone https://github.com/CharlesMuchogo/point_of_sale.git
WORKDIR /build/point_of_sale
COPY .env .
COPY go.mod .
COPY go.sum .
RUN go build -o main
EXPOSE 8080
ENTRYPOINT [ "/build/point_of_sale/main" ]


