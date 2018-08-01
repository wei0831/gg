package utli

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func locateGitFolder(dir string) (string, bool) {
	curPath, _ := filepath.Abs(dir)
	curPath = filepath.ToSlash(curPath)
	attempt := path.Join(curPath, gitFolderName)
	_, err := os.Stat(attempt)
	// while file not exists
	for err != nil {
		curPath = path.Dir(curPath)
		attempt = path.Join(curPath, gitFolderName)
		_, err = os.Stat(attempt)
		// If the path is empty, we have reached to the root
		if curPath == "." {
			return "", false
		}
	}
	return attempt, true
}

func convertRemoteURL(url string) string {
	toReplace := []string{"https://", "ssh://git@", "git@", "git://"}
	for _, r := range toReplace {
		if strings.HasPrefix(url, r) {
			url = url[len(r):]
			url = strings.Replace(url, ":", "/", -1)
		}
	}
	if strings.HasSuffix(url, ".git") {
		url = url[0 : len(url)-4]
	}
	return "https://" + url
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
