FROM alpine:latest as base

ENV USER=appuser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

FROM busybox

WORKDIR /app

COPY ./bin/main /app/main
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
RUN chmod -R 755 /app && chown -R appuser:appuser /app
USER appuser:appuser

CMD [ "/app/main" ]