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

an extra param `format` sets the output format, which takes one of `html`, `htmlraw`, `text` or `json`, default is `html`.

diffing settings can also be controlled by passing in additional query params, setting things like `matchThreshold` & `matchDistance`, in lieu of proper documentation, please see the `NewDifferFromRequest` func in [`differ.go`](differ.go).