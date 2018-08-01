# ghg
## Git Helper in Go  
This project is inspired by [cnishina/git-web-launcher](https://github.com/cnishina/git-web-launcher).

![cmdOpen](doc/screenshots/cmdOpen.gif?raw=true)

## Requirements
* [Git](https://git-scm.com/download) 
* [Go](https://golang.org/dl/)

## Install
Install via go command:
```bash
$ go get -u github.com/wei0831/ghg
```

## Usage
```bash
ghg [commands] [flags]
```

### Help
```bash
$ ghg -h
```

### Open
Open current git repo folder's remote Source URL in the browser.
```bash
$ ghg open
```
or 
```bash
$ ghg open .
```

### Blame
Open current git repo file's remote Blame URL in the browser.
```bash
$ ghg blame ./[FileName]
```

## Uninstall
Uninstall via go command:
```bash
$ go clean -i github.com/wei0831/ghg
```

## Supported Git Providers
#### For commands ```Open``` and ```Blame```
* Github.com
* Gitlab.com
* BitBucket.org
