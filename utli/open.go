package utli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	gitFolderName string = ".git"
)

// Open launches the source/blame remote URL for the target path with the given branch.
func Open(p, remote, branch string, blame bool) {
	// Transform input path to abs path
	curPath, err := filepath.Abs(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	curPath = filepath.ToSlash(curPath)

	// Check if path is valid
	f, err := os.Stat(curPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Extract dir name
	var dir string
	pathIsDir := f.IsDir()
	if pathIsDir {
		dir = curPath
	} else {
		dir, _ = filepath.Split(curPath)
	}

	// Get project root
	projectRoot, err := GetProjectRoot(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	projectRoot = strings.TrimRight(projectRoot, "\n")

	// Get remote URL and convert to https url
	remoteURL, err := GetRemoteURL(dir, remote)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remoteURL = convertRemoteURL(remoteURL)
	// Completed path - Project root path = url path to the target dir/file
	urlPath := curPath[len(projectRoot):]

	// Generate blame url or source url
	if blame {
		finalURL, err := generateBlameURL(remoteURL, pathIsDir, branch, urlPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Opening... %s\n", finalURL)
		openBrowser(finalURL)
	} else {
		finalURL, err := generateSourceURL(remoteURL, pathIsDir, branch, urlPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Opening... %s\n", finalURL)
		openBrowser(finalURL)
	}
}
