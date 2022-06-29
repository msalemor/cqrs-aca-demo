FROM golang:alpine as app-builder
RUN apk update
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates

# add a user here because addgroup and adduser are not available in scratch
ENV USER=appuser
ENV UID=10001

RUN apk update && apk upgrade && \
    apk add --no-cache \
    git ca-certificates tzdata

RUN update-ca-certificates

RUN adduser \                           
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

#RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

WORKDIR /workspace

COPY . .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build \
    -ldflags '-w -s -extldflags "-static"' \
    -a -o /out/app main.go

FROM scratch as final
WORKDIR /workspace

ENV USER=appuser                        
COPY --from=0 /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /etc/passwd /etc/passwd 
COPY --from=0 /etc/group /etc/group
COPY --from=0 /out/app ./
# Copy the app to final
#COPY --from=app-builder /myapp /myapp
# copy the certs to final
#COPY --from=app-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# copy the passwords to final
#COPY --from=app-builder /etc/passwd /etc/passwd

EXPOSE 80
EXPOSE 443

USER appuser
#USER nobody

ENTRYPOINT ["./app"]
