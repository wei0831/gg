package utli

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	// ProviderUnkown : Unknown
	ProviderUnkown int = 0
	// ProviderGithub : Github.com
	ProviderGithub int = 1
	// ProviderGitlab : Gitlab.com
	ProviderGitlab int = 2
	// ProviderBitBucket : BitBucket.org
	ProviderBitBucket int = 3
)

func generateSourceURL(remoteURL string, isDir bool, branch, pathURL string) (string, error) {
	provider := checkProvider(remoteURL)
	switch provider {
	case ProviderGithub, ProviderGitlab:
		if isDir {
			return fmt.Sprintf("%s/tree/%s%s", remoteURL, branch, pathURL), nil
		}
		return fmt.Sprintf("%s/blob/%s%s", remoteURL, branch, pathURL), nil
	case ProviderBitBucket:
		return fmt.Sprintf("%s/src/%s%s", remoteURL, branch, pathURL), nil
	}
	return "", errors.New("Remote URL: provider not supported./n")
}

func generateBlameURL(remoteURL string, isDir bool, branch, pathURL string) (string, error) {
	provider := checkProvider(remoteURL)

	switch provider {
	case ProviderGithub, ProviderGitlab:
		if isDir {
			return fmt.Sprintf("%s/tree/%s%s", remoteURL, branch, pathURL), nil
		}
		return fmt.Sprintf("%s/blame/%s%s", remoteURL, branch, pathURL), nil
	case ProviderBitBucket:
		if isDir {
			return fmt.Sprintf("%s/src/%s%s", remoteURL, branch, pathURL), nil
		}
		return fmt.Sprintf("%s/annotate/%s%s", remoteURL, branch, pathURL), nil
	}
	return "", errors.New("Remote URL: provider not supported./n")
}

func checkProvider(remoteURL string) int {
	if isGithub, _ := regexp.MatchString(`^https:\/\/.*github.com\/.*`, remoteURL); isGithub {
		return ProviderGithub
	}
	if isGitlab, _ := regexp.MatchString(`^https:\/\/.*gitlab.com\/.*`, remoteURL); isGitlab {
		return ProviderGitlab
	}
	if isBitBucket, _ := regexp.MatchString(`^https:\/\/.*bitbuck.org\/.*`, remoteURL); isBitBucket {
		return ProviderBitBucket
	}
	return ProviderUnkown
}
