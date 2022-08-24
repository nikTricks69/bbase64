kubectl currently doesnt allow me to decode secrets and sometimes its a pain in windows , in linux its an easy pipe
Hence scribbled a simple lean implementation in golang

## Get started

 1.checkout the repo and do a go build to build the exe

## Usage

```
bbase64 operation payload
operation - d for decode and e for encode

example 
    bbase64 e FABULOUS --this should base64 encode 'FABULOUS'

```


