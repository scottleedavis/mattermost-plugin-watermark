package main

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"gopkg.in/auyer/steganography.v2"
)

// Plugin comment
type Plugin struct {
	plugin.MattermostPlugin

	configurationLock sync.RWMutex

	configuration *configuration
}

//FileWillBeUploaded hook
func (p *Plugin) FileWillBeUploaded(c *plugin.Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {

	switch strings.ToUpper(info.Extension) {
	case "JPG", "JPEG", "PNG":
		data, err := ioutil.ReadAll(file)
		if err != nil {
			p.API.LogError(err.Error())
			return nil, err.Error()
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			errMsg := "ERROR: original image is corrupt " + err.Error()
			p.API.LogInfo(errMsg)
			return nil, errMsg
		}

		configuration := p.getConfiguration()
		w := &bytes.Buffer{}
		err = steganography.Encode(w, img, []byte(configuration.WaterMark))
		if err != nil {
			p.API.LogError(err.Error())
			return nil, err.Error()
		}
		if _, err := output.Write(w.Bytes()); err != nil {
			p.API.LogError(err.Error())
		}

	}
	return nil, ""
}
