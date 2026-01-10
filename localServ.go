package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// FILESYSTEM
func currentPath() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(cwd)
}
func getFromDir() {
	fmt.Print("\n")

	dirPath := "."

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			color.Green("%v (Folder)\n", entry.Name())
		} else {
			fileInfo, err := os.Stat(entry.Name())
			if err != nil {
				fmt.Println("Error")
				return
			}
			fmt.Printf("%v (%v bytes)\n", entry.Name(), fileInfo.Size())
		}
	}
	fmt.Print("\n")
}
func changeDir(path string) {
	os.Chdir(path)
}

func choosingDirectory(dirName string) {
	DirInfo, err := os.Stat(dirName)
	if err != nil {
		color.Red("Directory %s not found", dirName)
		return
	}
	isDir := DirInfo.IsDir()
	if !isDir {
		color.Red("%v NOT A DIRECTORY", dirName)
		return
	}
	color.Green("PROJECT FOLDER: %v (%v bytes)\n", dirName, DirInfo.Size())

	starting := true
	var port string
	reader := bufio.NewReader(os.Stdin)
	color.Yellow("Configure port(port: ....)")
	for starting {
		fmt.Print("STARTING>>")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error: ", err)
			break
		}
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		if strings.HasPrefix(input, "port: ") {
			port = strings.TrimSpace(strings.TrimPrefix(input, "port: "))
			if port == "" {
				color.Red("USING DEFAULT PORT 8080")
				port = "8080"
			}
			color.Yellow("Is this a React/Angular/Vue porject? (yes/no)")
			fmt.Print("STARTING>> ")
			spaInput, _ := reader.ReadString('\n')
			spaInput = strings.TrimSpace(strings.ToLower(spaInput))

			if spaInput == "yes" || spaInput == "y" {
				startSpaServer(dirName, port)
			} else if spaInput == "no" || spaInput == "n" {
				startServer(dirName, port)
			} else {
				color.Red("Please answer 'yes' or 'no'")
			}

		}

	}

}
func startSpaServer(dirName, port string) {
	absPath, _ := filepath.Abs(dirName)

	indexPath := filepath.Join(absPath, "index.html")
	if _, err := os.Stat(indexPath); err != nil {
		color.Red("ERROR: index.html not found in %s", absPath)
		color.Red("SPA applications require index.html in the root folder")
		return
	}

	color.Green("DETECTED SPA SIGNATURE")

	fs := http.FileServer(http.Dir(absPath))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		ip := r.RemoteAddr
		now := time.Now().Format("15:04:05")
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		color.Cyan("[%s] [%s] %s %s", now, ip, r.Method, r.URL.Path)

		if isStaticFile(path) {
			fs.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, indexPath)

	}))
	var info string = `
╔════════════════════════════════════════════════╗
║                SERVER IS WORKING!              ║
╚════════════════════════════════════════════════╝
Folder: %s
Port: %s
URL: http://%s:%s
Mode: LOCAL
`
	color.Green(info, absPath, port, "localhost", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		color.Red("SERVER ERROR: ", err)
	}

}
func isStaticFile(path string) bool {
	staticExtensions := []string{
		".js", ".css", ".json", ".png", ".jpg", ".jpeg", ".gif", ".svg",
		".ico", ".woff", ".woff2", ".ttf", ".eot", ".mp4", ".mp3", ".wav",
		".pdf", ".zip", ".txt", ".xml", ".webp", ".webm",
	}

	for _, ext := range staticExtensions {
		if strings.HasSuffix(strings.ToLower(path), ext) {
			return true
		}
	}

	// Пути, которые всегда статические
	staticPaths := []string{
		"/static/",
		"/assets/",
		"/public/",
		"/_next/",
		"/build/",
		"/dist/",
	}

	for _, staticPath := range staticPaths {
		if strings.HasPrefix(path, staticPath) {
			return true
		}
	}

	return false
}
func startServer(dirName, port string) {
	absPath, _ := filepath.Abs(dirName)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now().Format("15:04:05")
		if idx := strings.LastIndex(ip, ":"); idx != -1 {
			ip = ip[:idx]
		}
		color.Cyan("[%s] [%s] %s %s", now, ip, r.Method, r.URL.Path)

		http.FileServer(http.Dir(absPath)).ServeHTTP(w, r)

	}))
	var info string = `
╔════════════════════════════════════════════════╗
║                SERVER IS WORKING!              ║
╚════════════════════════════════════════════════╝
Folder: %s
Port: %s
URL: http://%s:%s
Mode: LOCAL
`
	color.Green(info, absPath, port, "localhost", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		color.Red("Error: %v", err)
	}
}

// MAIN FUNC
func localServer() {

	var greet = `
╔════════════════════════════════════════════════╗
║                  LOCAL SERVER                  ║ 
╚════════════════════════════════════════════════╝


╔═════════════════════════════════════════════════╗
║           PROJECT DIRECTORY SETUP               ║
╔═════════════════════════════════════════════════╗
║                                                 ║                             					 
║   Navigate to the parent directory              ║
║    (one level above your project folder)        ║
║                                                 ║	 
║   Then enter the name of your project directory ║
║    (Example: dir- ...)                          ║
║                                                 ║
╚═════════════════════════════════════════════════╝
`

	color.Yellow(greet)
	reader := bufio.NewReader(os.Stdin)

	running := true

	for running {
		fmt.Print("PROJECT DIRECTORY>>")
		input, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error", err)
			break
		}
		input = strings.TrimSpace(input)

		if input == "cancel" {
			color.Red("CANCELED")
			running = false
		} else if input == "pwd" {
			currentPath()
		} else if input == "ls" {
			getFromDir()
		} else if input[:3] == "cd " {
			changeDir(input[3:])
		} else if input[:5] == "dir- " {
			choosingDirectory(input[5:])
		} else if input == "" {
			continue
		} else if input[:3] != "" {
			color.Red("Unknown command")
			return
		}
	}
}
