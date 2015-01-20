go-commonmark
=======


[![Build Status](https://travis-ci.org/rhinoman/go-commonmark.svg?branch=master)](https://travis-ci.org/rhinoman/go-commonmark)

Description
-----------

go-commonmark is a [Go](http://golang.org) (golang) wrapper for the CommonMark C library


Installation
------------

1.) First, install CommonMark located at: https://github.com/jgm/CommonMark 

2.) Make sure the directory you install libcmark to is somewhere your system will find it at runtime (i.e., somewhere in your /etc/ld.so.conf: /usr/lib, etc.).  Or you can be lazy and add its location to your LD_LIBRARY_PATH. 

3.) Then, just:

```
go get github.com/rhinoman/go-commonmark
```

**Note:**  The C library is still considered 'pre-release' and is under fairly heavy development.  Thus, changes that break this wrapper can happen from time to time, and I'll be playing catch-up when that occurs.  It's probably best to use one of the tags with it's corresponding version branch/tag of CommonMark (0.16 is the most recent as of this writing).


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
