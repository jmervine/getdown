# getdown

### docker

```
# quickstart
docker run -p 3000:3000 -v /path/to/markdown:/data jmervine/getdown

# custom config example
docker run -p 3000:3000 -v /path/to/markdown:/data \
    -e GETDOWN_STYLE="https://maxcdn.bootstrapcdn.com/bootswatch/3.3.5/flatly/bootstrap.min.css" \
    -e GETDWON_INDEX="readme.md" \
    jmervine/getdown
```

### local

```
$ go get github.com/jmervine/getdown
$ getdown -h

NAME:
    getdown - Markdown file parser and server written in Go and designed to run on Docker.

USAGE:
    getdown [global options] command [command options] [arguments...]

VERSION:
    0.0.1

COMMANDS:
    help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
    --addr, -a 'localhost'
         listener address [$GETDOWN_ADDR]
    --port, -p '3000'
         listener port [$GETDOWN_PORT]
    --basedir, -b '.'
         base markdown directory [$GETDOWN_BASEDIR]
    --index, -i 'README.md'
         root file default [$GETDOWN_INDEX]
    --style 'https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css'
         bootstrap compatable stylesheet url [$GETDOWN_STYLE]
    --help, -h
         show help
    --version, -v
         print the version
```

