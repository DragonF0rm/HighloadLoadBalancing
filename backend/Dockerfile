FROM golang:1.12.11-alpine AS build
LABEL maintainer="Uimin Maksim"
WORKDIR /backend
COPY . .
RUN yes| apk add git
RUN ./etc/build_cfg.py
RUN go build -o bin/backend main.go

FROM alpine
LABEL maintainer="Uimin Maksim"
COPY --from=build /backend/bin/backend /usr/bin/backend
COPY --from=build /backend/etc/backend.yaml /etc/backend.yaml
CMD ["/etc/backend.yaml"]
ENTRYPOINT ["/usr/bin/backend"]
