FROM golang:1.19

# create directory app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

# copy all files to /app
COPY ./ /app

RUN go build -o joshgarden-app

CMD ./joshgarden-app