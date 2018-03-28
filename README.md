
# go-random-downloader
Download Html using "Random Page"




[![Build Status](https://travis-ci.org/binaryplease/go-random-downloader.svg?branch=master)](https://travis-ci.org/binaryplease/go-random-downloader)
[![GoDoc](https://godoc.org/github.com/binaryplease/go-random-downloader?status.svg)](https://godoc.org/github.com/binaryplease/go-random-downloader)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/go-random-downloader)](https://goreportcard.com/report/github.com/binaryplease/go-random-downloader)
[![codecov](https://codecov.io/gh/binaryplease/go-random-downloader/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/go-random-downloader)

# Usage
```
NAME:
   random-downloader - Download website sources using the Random Page function

USAGE:
   random-downlader url [-n number] [-d directory] [-r]

DESCRIPTION:
   A crawler to download websites using the random page link

AUTHOR:
   Pablo Ovelleiro <pablo1@mailbox.org>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -n value, --number value     Number of pages to download (default: 1)
   --url value, -u value        The url to use
   --retries value, -r value    Retries to get new page (default: 3)
   --directory value, -d value  The download destination directry (default: "./downloaded-content")
   --help, -h                   show help
```
