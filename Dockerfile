FROM alpine:latest

WORKDIR /home/

COPY build/quotes_api .

EXPOSE ${PORT}

CMD ["./quotes_api"]
