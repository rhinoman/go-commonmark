go-commonmark
=======


Description
-----------

go-commonmark is a Go wrapper for the CommonMark C library


Installation
------------

1.) First, install CommonMark located at: https://github.com/jgm/CommonMark

2.) Make sure the directory containg libcmark (e.g., /usr/local/lib) is in your LD_LIBRARY_PATH

3.) Then, just:

```
go get github.com/rhinoman/go-commonmark
```

Documentation
-------------

See the Godoc: http://godoc.org/github.com/rhinoman/go-commonmark


Example Usage
-------------

```go

import "github.com/rhinoman/go-commonmark"

...

	htmlText := commonmark.Md2Html(mdText)  

```
