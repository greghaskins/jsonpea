# JSON<sub>pea</sub>
A tiny JSONP proxy server written in Go

[![Build Status](https://travis-ci.org/greghaskins/jsonpea.svg)](https://travis-ci.org/greghaskins/jsonpea)

## A JSONP proxy for local development

`jsonpea` is an alternative to [anyorigin.com](http://anyorigin.com/) or [whateverorigin.org](http://www.whateverorigin.org/) for local development. It's a little single-binary proxy server to work around same-origin issues while developing websites. It fetches the content from a given `url` and calls the `callback` function with the result.

### Installation

```
go get github.com/greghaskins/jsonpea
```

### Usage

Start up `jsonpea`:

```
~$ $GOPATH/bin/jsonpea
Listening on :7070 ...
```

Then hit `http://localhost:7070` in another terminal or browser:
```
GET /get?url=http://httpbin.org/get&callback=doSomething
```

```
HTTP/1.1 200 OK

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

#### Query Parameters for `/get`

- `url` (required): The remote URL to load.
- `callback` (optional): A named JavaScript function to be called with the payload passed as a parameter. If no callback is specified, the raw data is returned as-is.
