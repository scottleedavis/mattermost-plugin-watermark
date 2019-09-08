# Mattermost Plugin Watermark
![CircleCI branch](https://img.shields.io/circleci/project/github/scottleedavis/mattermost-plugin-watermark/master.svg)

A plugin for [Mattermost](https://mattermost.com) that uses a [steganography library](https://github.com/auyer/steganography) to add a hidden string in an image that identifies an image as having been uploaded to the server.

Currently supports PNG files.  
Planned JPG support too.

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
