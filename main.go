package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const indexHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Download App</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
            min-height: 100vh;
            background: linear-gradient(135deg, #dbeafe 0%, #e0e7ff 100%);
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 1rem;
        }
        
        .card {
            background: white;
            border-radius: 0.5rem;
            box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
            width: 100%;
            max-width: 28rem;
        }
        
        .card-header {
            text-align: center;
            padding: 1.5rem 1.5rem 0;
        }
        
        .icon-container {
            width: 4rem;
            height: 4rem;
            background-color: #dbeafe;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            margin: 0 auto 1rem;
        }
        
        .smartphone-icon {
            width: 2rem;
            height: 2rem;
            color: #2563eb;
        }
        
        .card-title {
            font-size: 1.5rem;
            font-weight: 700;
            color: #111827;
            margin-bottom: 0.5rem;
        }
        
        .card-description {
            color: #6b7280;
            line-height: 1.5;
        }
        
        .card-content {
            padding: 1.5rem;
        }
        
        .download-btn {
            width: 100%;
            background-color: #2563eb;
            color: white;
            border: none;
            border-radius: 0.375rem;
            padding: 0.75rem 1rem;
            font-size: 1.125rem;
            font-weight: 600;
            cursor: pointer;
            text-decoration: none;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-bottom: 1.5rem;
            transition: background-color 0.2s;
        }
        
        .download-btn:hover {
            background-color: #1d4ed8;
        }
        
        .download-icon {
            width: 1.25rem;
            height: 1.25rem;
            margin-right: 0.5rem;
        }
        
        .file-info {
            background-color: #f9fafb;
            border-radius: 0.5rem;
            padding: 1rem;
            margin-bottom: 1.5rem;
        }
        
        .file-info-header {
            display: flex;
            align-items: center;
            font-weight: 600;
            color: #111827;
            margin-bottom: 0.75rem;
        }
        
        .file-icon {
            width: 1rem;
            height: 1rem;
            margin-right: 0.5rem;
        }
        
        .file-details {
            font-size: 0.875rem;
            color: #6b7280;
        }
        
        .file-row {
            display: flex;
            justify-content: space-between;
            margin-bottom: 0.5rem;
        }
        
        .file-row:last-child {
            margin-bottom: 0;
        }
        
        .file-value {
            font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
            display: flex;
            align-items: center;
        }
        
        .harddrive-icon {
            width: 0.75rem;
            height: 0.75rem;
            margin-right: 0.25rem;
        }
        
        .warning-box {
            background-color: #fffbeb;
            border: 1px solid #fed7aa;
            border-radius: 0.5rem;
            padding: 1rem;
            display: flex;
            align-items: flex-start;
        }
        
        .warning-icon {
            width: 1.25rem;
            height: 1.25rem;
            background-color: #fef3c7;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            margin-right: 0.75rem;
            flex-shrink: 0;
            margin-top: 0.125rem;
        }
        
        .warning-icon span {
            color: #d97706;
            font-size: 0.75rem;
            font-weight: 700;
        }
        
        .warning-content h4 {
            font-weight: 600;
            color: #92400e;
            font-size: 0.875rem;
            margin-bottom: 0.25rem;
        }
        
        .warning-content p {
            color: #b45309;
            font-size: 0.875rem;
            line-height: 1.4;
        }
    </style>
</head>
<body>
    <div class="card">
        <div class="card-header">
            <div class="icon-container">
                <svg class="smartphone-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <rect x="5" y="2" width="14" height="20" rx="2" ry="2"></rect>
                    <line x1="12" y1="18" x2="12.01" y2="18"></line>
                </svg>
            </div>
            <h1 class="card-title">Download Hyperchat 2.0</h1>
            <p class="card-description">Click the button below to download the latest version of the app.</p>
        </div>
        
        <div class="card-content">
            <a href="/download" class="download-btn">
                <svg class="download-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path d="m3 16 4 4 4-4"></path>
                    <path d="M7 20V4"></path>
                    <path d="m13 8 4-4 4 4"></path>
                    <path d="M17 4v16"></path>
                </svg>
                Download APK
            </a>
            
            <div class="file-info">
                <h3 class="file-info-header">
                    <svg class="file-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                        <polyline points="14,2 14,8 20,8"></polyline>
                        <line x1="16" y1="13" x2="8" y2="13"></line>
                        <line x1="16" y1="17" x2="8" y2="17"></line>
                        <polyline points="10,9 9,9 8,9"></polyline>
                    </svg>
                    File Information
                </h3>
                <div class="file-details">
                    <div class="file-row">
                        <span>File:</span>
                        <span class="file-value">{{.FileName}}</span>
                    </div>
                    <div class="file-row">
                        <span>Size:</span>
                        <span class="file-value">
                            <svg class="harddrive-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <line x1="22" y1="12" x2="2" y2="12"></line>
                                <path d="M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"></path>
                                <line x1="6" y1="16" x2="6.01" y2="16"></line>
                                <line x1="10" y1="16" x2="10.01" y2="16"></line>
                            </svg>
                            {{.FileSize}}
                        </span>
                    </div>
                </div>
            </div>
            
            <div class="warning-box">
                <div class="warning-icon">
                    <span>!</span>
                </div>
                <div class="warning-content">
                    <h4>Installation Note</h4>
                    <p>Make sure to enable "Install from unknown sources" on your Android device.</p>
                </div>
            </div>
        </div>
    </div>
</body>
</html>
`

// NEW: App opener HTML template
const openAppHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Opening HyperChat App...</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            margin: 0;
            padding: 20px;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            color: white;
            text-align: center;
        }
        
        .container {
            max-width: 400px;
            background: rgba(255, 255, 255, 0.1);
            backdrop-filter: blur(10px);
            border-radius: 20px;
            padding: 40px;
            box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
        }
        
        .icon {
            font-size: 4rem;
            margin-bottom: 20px;
        }
        
        h1 {
            margin: 0 0 20px 0;
            font-size: 1.5rem;
            font-weight: 600;
        }
        
        .status {
            margin: 20px 0;
            font-size: 1.1rem;
        }
        
        .download-btn {
            display: inline-block;
            background: #4CAF50;
            color: white;
            padding: 15px 30px;
            text-decoration: none;
            border-radius: 50px;
            font-weight: 600;
            margin-top: 20px;
            transition: background 0.3s;
        }
        
        .download-btn:hover {
            background: #45a049;
        }
        
        .spinner {
            border: 3px solid rgba(255, 255, 255, 0.3);
            border-top: 3px solid white;
            border-radius: 50%;
            width: 30px;
            height: 30px;
            animation: spin 1s linear infinite;
            margin: 20px auto;
        }
        
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
        
        .hidden { display: none; }
    </style>
</head>
<body>
    <div class="container">
        <div class="icon">🩺</div>
        <h1>Opening HyperChat App...</h1>
        
        <div id="loading" class="status">
            <div class="spinner"></div>
            <p>Attempting to open the app...</p>
        </div>
        
        <div id="fallback" class="status hidden">
            <p>📱 App not installed or couldn't open?</p>
            <a href="https://hyperchat.up.railway.app/" class="download-btn">
                Download HyperChat App
            </a>
        </div>
    </div>

    <script>
        function tryOpenApp() {
            const userAgent = navigator.userAgent || navigator.vendor || window.opera;
            
            const deepLink = 'hyperchat://open';
            const intentLink = 'intent://open#Intent;scheme=hyperchat;package=com.samess.hyperchat_app;S.browser_fallback_url=https%3A%2F%2Fhyperchat.up.railway.app%2F;end';
            
            let appOpened = false;
            
            if (/android/i.test(userAgent)) {
                try {
                    window.location.href = intentLink;
                    appOpened = true;
                } catch (e) {
                    console.log('Intent failed, trying deep link');
                    try {
                        window.location.href = deepLink;
                        appOpened = true;
                    } catch (e2) {
                        console.log('Deep link failed');
                    }
                }
            } else {
                try {
                    window.location.href = deepLink;
                    appOpened = true;
                } catch (e) {
                    console.log('Deep link failed');
                }
            }
            
            setTimeout(() => {
                if (!appOpened || !document.hidden) {
                    document.getElementById('loading').classList.add('hidden');
                    document.getElementById('fallback').classList.remove('hidden');
                }
            }, 3000);
            
            document.addEventListener('visibilitychange', () => {
                if (document.hidden) {
                    appOpened = true;
                }
            });
        }
        
        tryOpenApp();
        
        window.addEventListener('load', () => {
            setTimeout(tryOpenApp, 500);
        });
    </script>
</body>
</html>
`

type PageData struct {
	FileName string
	FileSize string
}

func formatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return "63.6 MB"
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Parse templates
	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		log.Fatal("Error parsing index template:", err)
	}

	openAppTmpl, err := template.New("openApp").Parse(openAppHTML)
	if err != nil {
		log.Fatal("Error parsing open-app template:", err)
	}

	// APK file path - change this to your APK file path
	apkPath := "./hyperchat.apk"

	// Home page (unchanged)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Get file info
		fileInfo, err := os.Stat(apkPath)
		var fileName, fileSize string

		if err != nil {
			fileName = "hyperchat.apk (file not found)"
			fileSize = "Unknown"
		} else {
			fileName = fileInfo.Name()
			fileSize = formatFileSize(fileInfo.Size())
		}

		data := PageData{
			FileName: fileName,
			FileSize: fileSize,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})

	// NEW: App opener route
	r.Get("/open-app", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if err := openAppTmpl.Execute(w, nil); err != nil {
			http.Error(w, "Error rendering open-app template", http.StatusInternalServerError)
			return
		}
	})

	// Download endpoint (unchanged)
	r.Get("/download", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://drive.google.com/drive/folders/1q3ELPL61wIZ-FOHCptv8Cw7z5cR6W_UI", http.StatusFound)
	})

	// Health check (unchanged)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("APK file path: %s", apkPath)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
