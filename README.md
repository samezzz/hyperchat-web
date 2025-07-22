# APK Download Server

A simple and elegant Go web server for hosting and distributing Android APK files. Built with Chi router and featuring a modern, responsive web interface.

## Features

- 🚀 Simple and fast Go web server
- 📱 Modern, responsive UI design
- 📊 Automatic file size detection and formatting
- 🔒 Proper APK MIME type headers for secure downloads
- 🎯 Health check endpoint for monitoring
- ⚡ Lightweight with minimal dependencies
- 🌐 Cross-platform deployment ready

## Demo

![APK Download Server Screenshot](screenshot.png)

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Your Android APK file

### Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/apk-download-server.git
cd apk-download-server
```

2. Install dependencies:

```bash
go mod tidy
```

3. Place your APK file in the project directory and name it `app.apk` (or update the `apkPath` in `main.go`)

4. Run the server:

```bash
go run main.go
```

5. Open your browser and navigate to `http://localhost:8080`

## Configuration

### Environment Variables

- `PORT` - Server port (default: 8080)

### Customization

To use a different APK file, update the `apkPath` variable in `main.go`:

```go
// Change this to your APK file path
apkPath := "./your-app.apk"
```

## API Endpoints

| Endpoint    | Method | Description                |
| ----------- | ------ | -------------------------- |
| `/`         | GET    | Main download page with UI |
| `/download` | GET    | Direct APK file download   |
| `/health`   | GET    | Health check endpoint      |

## Deployment

### Docker

1. Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/app.apk .

EXPOSE 8080
CMD ["./main"]
```

2. Build and run:

```bash
docker build -t apk-server .
docker run -p 8080:8080 apk-server
```

### Heroku

1. Create a `Procfile`:

```
web: ./main
```

2. Deploy:

```bash
heroku create your-app-name
git push heroku main
```

### Google Cloud Run

1. Deploy directly from source:

```bash
gcloud run deploy apk-server --source . --platform managed --region us-central1 --allow-unauthenticated
```

### DigitalOcean App Platform

1. Connect your GitHub repository
2. Choose "Web Service"
3. Set build command: `go build -o main .`
4. Set run command: `./main`
5. Deploy!

## Project Structure

```
apk-download-server/
├── main.go          # Main server code
├── go.mod           # Go module file
├── go.sum           # Go dependencies
├── app.apk          # Your APK file (not in git)
├── Dockerfile       # Docker configuration
├── Procfile         # Heroku configuration
└── README.md        # This file
```

## Development

### Running in Development

```bash
# Install air for hot reloading (optional)
go install github.com/cosmtrek/air@latest

# Run with hot reloading
air

# Or run normally
go run main.go
```

### Building for Production

```bash
# Build binary
go build -o apk-server

# Run binary
./apk-server
```

## Security Considerations

- The server serves files without authentication - ensure you're comfortable with public access
- Consider adding rate limiting for production use
- Use HTTPS in production (most hosting platforms provide this automatically)
- Regularly update your APK file and dependencies

## Customizing the UI

The HTML template is embedded in `main.go`. To customize:

1. Modify the `indexHTML` constant in `main.go`
2. Update CSS styles within the `<style>` tags
3. The template uses Go's `html/template` package with these variables:
   - `{{.FileName}}` - Name of the APK file
   - `{{.FileSize}}` - Formatted file size

## Troubleshooting

### Common Issues

**APK file not found**

- Ensure your APK file exists in the specified path
- Check file permissions
- Update `apkPath` variable if using a different location

**Port already in use**

- Change the port using the `PORT` environment variable
- Kill existing processes using the port: `lsof -ti:8080 | xargs kill`

**Build errors**

- Ensure you're using Go 1.21 or higher: `go version`
- Run `go mod tidy` to resolve dependencies

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit
