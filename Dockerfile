ARG BUILDER_TAG="1.22-bookworm"
ARG BUILD_TAG=static-debian12

FROM golang:${BUILDER_TAG} AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -a -o /out/protovision .

FROM gcr.io/distroless/${BUILD_TAG}
LABEL author="Jean-Philippe Ruijs <jean-philippe@ruijs.fr>"
WORKDIR /app
COPY --from=build /out/protovision .
COPY *.json .
USER nonroot:nonroot

CMD [ "/app/protovision" ]