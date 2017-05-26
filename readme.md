# Differ

first, [install go](http://golang.org), then `cd` to the folder, build & run. default port is 8080.

```
$ cd differ
$ go build
$ ./differ
```

It's only got one route: `/`, hit it with `a` & `b` query params, where `a` & `b` are urls to diff, eg:

```
http://localhost:8000/diff?a=http://www.apple.com&b=http://www.apple.com/a
```

an extra param `format`, which taks either `html`, `htmlraw`, `text` or `json`, default is html.