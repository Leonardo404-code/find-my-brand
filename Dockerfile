FROM golang:1.22-alpine AS base
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/server/main.go
ENV PORT=3000
ENV ENV=development
ENV API_KEY=eULfrL76DKhS48ZjbhW2Ny6y
ENV GOEMAIL_PASSWORD="zvit hlbt fxpd uoij"
ENV GOEMAIL_USER=leonardobispo.dev@gmail.com
EXPOSE $PORT
CMD ["./main"]