# Mattermost Plugin Watermark
[![Build Status](https://img.shields.io/circleci/project/github/scottleedavis/mattermost-plugin-watermark/master.svg)](https://circleci.com/gh/scottleedavis/mattermost-plugin-watermark)  [![codecov](https://codecov.io/gh/scottleedavis/mattermost-plugin-watermark/branch/master/graph/badge.svg)](https://codecov.io/gh/scottleedavis/mattermost-plugin-watermark)   [![Releases](https://img.shields.io/github/release/scottleedavis/mattermost-plugin-watermark.svg)](https://github.com/scottleedavis/mattermost-plugin-watermark/releases/latest)

A plugin for [Mattermost](https://mattermost.com) that uses a [steganography library](https://github.com/auyer/steganography) to add a hidden string in an image uploaded to the server.

Supports PNG & JPG.

##### Decode
To view the encoded message/watermark in the file, run `go run decode.go`.  For example:
```bash
$ go run decode.go assets/test.png 
This is an image that has been uploaded to Mattermost
```


##### Build
```
make
```

This will produce a single plugin file (with support for multiple architectures) for upload to your Mattermost server:

```
dist/com.github.scottleedavis.mattermost-plugin-watermark.tar.gz
```
