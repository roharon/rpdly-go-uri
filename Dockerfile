FROM golang:1.15.6-alpine

WORKDIR /app
COPY . .

RUN apk update && apk add git
RUN go get github.com/cespare/reflex

EXPOSE 3000
RUN ENVIRONMENT=prod
CMD ["reflex", "-c", "reflex.conf"]