go-commonmark
=======


[![Build Status](https://travis-ci.org/rhinoman/go-commonmark.svg?branch=master)](https://travis-ci.org/rhinoman/go-commonmark)

Description
-----------

go-commonmark is a Go wrapper for the CommonMark C library


Installation
------------

1.) First, install CommonMark located at: https://github.com/jgm/CommonMark

2.) Make sure the directory you install libcmark to is somewhere your system will find it at runtime (i.e., somewhere in your /etc/ld.so.conf: /usr/lib, etc.).  Or you can be lazy and add its location to your LD_LIBRARY_PATH. 

3.) Then, just:

```
go get github.com/rhinoman/go-commonmark
```

**Note:**  The C library is still considered 'pre-release' and is under fairly heavy development.  Thus, changes that break this wrapper can happen from time to time, and I'll be playing catch-up when that occurs.  Once a release of the CommonMark C library is made, this wrapper should be pegged to a specific release/tag.

For now, this wrapper does indeed work with CommonMark master/48d19922aa :)

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
