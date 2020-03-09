package main

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"gopkg.in/auyer/steganography.v2"
)

// Plugin comment
type Plugin struct {
	plugin.MattermostPlugin

	configurationLock sync.RWMutex

	configuration *configuration
}

//FileWillBeUploaded hook
func (p *Plugin) FileWillBeUploaded(_ *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	switch strings.ToLower(info.Extension) {
	case "jpg", "jpeg", "png":
		data, err := ioutil.ReadAll(file)
		if err != nil {
			errMsg := "Failed to read image: " + err.Error()
			p.API.LogWarn(errMsg)
			return nil, errMsg
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			errMsg := "Original image is corrupt: " + err.Error()
			p.API.LogWarn(errMsg)
			return nil, errMsg
		}

		configuration := p.getConfiguration()
		w := &bytes.Buffer{}

		err = steganography.Encode(w, img, []byte(configuration.WaterMark))
		if err != nil {
			errMsg := "Failed to encode watermark into image: " + err.Error()
			p.API.LogWarn(errMsg)
			return nil, errMsg
		}

		if _, err := output.Write(w.Bytes()); err != nil {
			errMsg := "Failed to write new image: " + err.Error()
			p.API.LogWarn(errMsg)
			return nil, errMsg
		}
	}

	return nil, ""
}
