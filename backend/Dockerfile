FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM busybox
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

ADD bin/main ./
ADD queries/get-issues.gql ./queries/get-issues.gql

EXPOSE 80/tcp
EXPOSE 443/tcp

CMD ["/main"]