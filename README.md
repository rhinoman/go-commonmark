go-mark
=======


Description
-----------

go-mark is a Go wrapper for the CommonMark C library


Installation
------------

1.) First, install CommonMark located at: https://github.com/jgm/CommonMark
2.) Make sure the library containg libcmark is in your LD_LIBRARY_PATH
3.) Then, just:
```
go get github.com/rhinoman/go-mark
```

Documentation
-------------

See the Godoc: http://godoc.org/github.com/rhinoman/go-mark


Example Usage
-------------

```go

import commonmark

...

	htmlText := commonmark.Md2Html(mdText)  

```
