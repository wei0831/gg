package utli

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// GetConfig : git --config --get -z key
func GetConfig(dir, key string) (string, error) {
	return ExecGit("config", "--get", "-z", key)
}

// GetLocal  git --config --get -z --local key
func GetLocal(dir, key string) (string, error) {
	return ExecGit(dir, "config", "--get", "-z", "--local", key)
}

// GetGlobal  git --config --get -z --global key
func GetGlobal(dir, key string) (string, error) {
	return ExecGit(dir, "config", "--get", "-z", "--global", key)
}

// GetProjectRoot  git rev-parse --show-toplevel
func GetProjectRoot(dir string) (string, error) {
	return ExecGit(dir, "rev-parse", "--show-toplevel")
}

// GetRemoteURL  git --config --get -z --local remote.%s.url
func GetRemoteURL(dir, remoteName string) (string, error) {
	return GetLocal(dir, fmt.Sprintf("remote.%s.url", remoteName))
}

// GetBranchs  git branch
func GetBranchs(dir string) (string, error) {
	return ExecGit(dir, "branch")
}

// ExecGit : git -C dir
func ExecGit(dir string, args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	args = append([]string{"-C", dir}, args...)
	execCmd := exec.Command("git", args...)
	execCmd.Stdout = &stdout
	execCmd.Stderr = &stderr
	err := execCmd.Run()
	if err != nil {
		return "ERROR", errors.New(stderr.String())
	}
	return strings.TrimRight(stdout.String(), "\000"), nil
}
