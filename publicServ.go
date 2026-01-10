package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

func PublicIp() (string, error) {
	services := []string{
		"https://api.ipify.org",
		"https://ifconfig.me/ip",
		"https://icanhazip.com",
		"https://checkip.amazonaws.com",
		"https://ipinfo.io/ip",
	}

	client := &http.Client{Timeout: 5 * time.Second}

	for _, url := range services {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			ip, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				return string(ip), nil
			}
		}
	}
	return "", fmt.Errorf("COULD NOT GET PUBLIC IP")
}
func choosingDirectory2(dirName string) {
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
				startSpaServer2(dirName, port)
				return
			} else if spaInput == "no" || spaInput == "n" {
				startServer2(dirName, port)
				return
			} else {
				color.Red("Please answer 'yes' or 'no'")
			}

		}

	}

}
func startSpaServer2(dirName, port string) {
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
	publicIp, error := PublicIp()
	var info string = `
╔════════════════════════════════════════════════╗
║                SERVER IS WORKING!              ║
╚════════════════════════════════════════════════╝
Folder: %s
Port: %s
URL: http://%s:%s
Mode: PUBLIC
`
	color.Green(info, absPath, port, publicIp, port)
	addr := ":" + port
	if error != nil {
		color.Red("COULD NOT GET PUBLIC IP")

	}
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		color.Red("SERVER ERROR: ", err)
	}

}
func startServer2(dirName, port string) {
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
	publicIp, error := PublicIp()
	if error != nil {
		color.Green("SERVER STARTED")
		color.Red("COULD NOT GET PUBLIC IP")
	}
	var info string = `
╔════════════════════════════════════════════════╗
║                SERVER IS WORKING!              ║
╚════════════════════════════════════════════════╝
Folder: %s
Port: %s
URL: http://%s:%s
Mode: PUBLIC
`
	color.Green(info, absPath, port, publicIp, port)
	addr := ":" + port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		color.Red("Error: %v", err)
	}
}
func public() {

	var greet = `
╔════════════════════════════════════════════════╗
║                 PUBLIC SERVER                  ║ 
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
			choosingDirectory2(input[5:])
		} else if input == "" {
			continue
		} else if input[:3] != "" {
			color.Red("Unknown command")
			return
		}
	}
}
