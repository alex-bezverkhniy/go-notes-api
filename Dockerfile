FROM alpine:3.9
COPY ./go-notes-api /app/go-notes-api
RUN chmod +x /app/go-notes-api
ENV API_PORT 8080
EXPOSE ${API_PORT}
ENTRYPOINT [ "/app/go-notes-api" ]