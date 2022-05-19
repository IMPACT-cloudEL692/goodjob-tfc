FROM golang:1.8-alpine
ADD . /go/src/goodjob-app
RUN go install goodjob-app

FROM alpine:latest
COPY --from=0 /go/bin/goodjob-app .
ENV PORT 8080
CMD ["./goodjob-app"]
