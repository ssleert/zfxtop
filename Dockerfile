FROM golang:1.19
WORKDIR /app
COPY . .
RUN "/app/scripts/build.sh"
CMD ["/app/zfxtop"]
