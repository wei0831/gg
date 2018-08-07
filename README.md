# ghg
## Git Helper in Go  
This project is inspired by [cnishina/git-web-launcher](https://github.com/cnishina/git-web-launcher).

![cmdOpenFile](doc/screenshots/cmdOpenFile.gif?raw=true)

## Requirements
* [Git](https://git-scm.com/download) 
* [Go](https://golang.org/dl/)

## Install
Install via go command:
```bash
$ go get -u github.com/wei0831/ghg
```

## Usage

### Open
```
Usage:
  ghg open [path](default=".") [flags]

Flags:
  -b, --branch string   branch (default "master")
  -h, --help            help for open
  -r, --remote string   remote (default "origin")
```

Open current file/directory's remote Source URL in the browser.
```bash
$ ghg open
```
```bash
$ ghg open README.md
```

### Blame
```
Usage:
  ghg blame [path](default=".") [flags]

Flags:
  -b, --branch string   branch (default "master")
  -h, --help            help for blame
  -r, --remote string   remote (default "origin")
```
Open current git repo file's remote Blame URL in the browser.
```bash
$ ghg blame ./[FileName]
```

### Help
```
Usage:
  ghg [command]

Available Commands:
  blame       Launches the blame remote URL for the target path with the given branch.
  help        Help about any command
  open        Launches the source remote URL for the target path with the given branch.

Flags:
  -h, --help   help for ghg

Use "ghg [command] --help" for more information about a command.
```

```bash
$ ghg -h
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
