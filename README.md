go-commonmark
=======


[![Build Status](https://travis-ci.org/rhinoman/go-commonmark.svg?branch=master)](https://travis-ci.org/rhinoman/go-commonmark)

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

Note:  The C API is stil considered 'pre-release' and is under fairly heavy development.  Thus, changes that break this wrapper can happen from time to time, and I'll be playing catch-up when that happens.  Once a release of CommonMark C API is made, this wrapper should be pegged to a specific release/tag.

For now, this wrapper does indeed work with CommonMark master/6291b23400 

Documentation
-------------

See the Godoc: http://godoc.org/github.com/rhinoman/go-commonmark


Example Usage
-------------
If all you need is to convert CommonMark text to Html, just do this:

```go

import "github.com/rhinoman/go-commonmark"

...

	htmlText := commonmark.Md2Html(mdText)  

```
