# JSON<sub>pea</sub>
A tiny JSONP proxy server written in Go

[![Build Status](https://travis-ci.org/greghaskins/jsonpea.svg)](https://travis-ci.org/greghaskins/jsonpea)

## A JSONP proxy for local development

`jsonpea` is an alternative to [anyorigin.com](http://anyorigin.com/) or [whateverorigin.org](http://www.whateverorigin.org/) for local development. It's a little single-binary proxy server to work around same-origin issues while developing websites.

### Installation

```
go get github.com/greghaskins/jsonpea
```

### Usage

```
~$ $GOPATH/bin/jsonpea
Listening on :7070 ...
```

And in another terminal or browser...
```
~$ curl http://localhost:7070/get\?url\=http://httpbin.org/get\&callback\=doSomething
doSomething({
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go 1.1 package http"
  },
  "origin": "...",
  "url": "http://httpbin.org/get"
}
);
```
