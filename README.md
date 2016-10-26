go-reddit
=========

[![Build Status](https://travis-ci.org/cameronstanley/go-reddit.svg?branch=master)](https://travis-ci.org/cameronstanley/go-reddit)
[![GoDoc](https://godoc.org/github.com/cameronstanley/go-reddit?status.svg)](https://godoc.org/github.com/cameronstanley/go-reddit)
[![Go Report Card](https://goreportcard.com/badge/github.com/cameronstanley/go-reddit)](https://goreportcard.com/report/github.com/cameronstanley/go-reddit)

A Golang wrapper for the [Reddit API](https://github.com/reddit/reddit/wiki/API). This package aims to implement every endpoint exposed according to the [documentation](https://www.reddit.com/dev/api) in a user friendly, well tested and documented manner.

## Installation

Install the package with

`go get github.com/cameronstanley/go-reddit`

## Authentication

Many endpoints in the Reddit API require OAuth2 authentication to access. To get started, register an app at https://www.reddit.com/prefs/apps and be sure to note the ID, secret, and redirect URI. 

## Examples

````Go
// Returns a new unauthenticated client for invoking the API
client := reddit.NoAuthClient

// Retrives a listing of default subreddits
client.GetDefaultSubreddits()

// Retrives a listing of hot links for the "news" subreddit
client.GetHotLinks("news")
````
