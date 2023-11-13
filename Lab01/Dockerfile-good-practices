##

## STEP 1 - BUILD

##

# specify the base image to be used for the application

FROM golang:1.21-alpine AS build

# create a working directory inside the image

WORKDIR /app

# copy Go modules and dependencies to image

COPY server/go.mod server/go.sum ./

# download Go modules and dependencies

RUN go mod download

# copy directory files i.e all files ending with .go

COPY server/*.go .

# compile application

RUN go build -o /godocker

##

## STEP 2 - DEPLOY

##

FROM scratch

WORKDIR /

COPY --from=build /godocker /godocker

EXPOSE 3000/tcp

ENTRYPOINT ["/godocker"]