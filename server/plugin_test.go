package main

import (
	"bufio"
	"bytes"
	"image"
	"io/ioutil"
	"testing"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin/plugintest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/auyer/steganography.v2"
)

func TestFileWillBeUpload(t *testing.T) {

	t.Run("original image is corrupt", func(t *testing.T) {
		setupAPI := func() *plugintest.API {
			api := &plugintest.API{}
			api.On("LogInfo", mock.Anything).Maybe()
			return api
		}

		api := setupAPI()
		defer api.AssertExpectations(t)
		p := &Plugin{}
		p.API = api

		fi := &model.FileInfo{
			Extension: "JPG",
		}

		r := bytes.NewReader([]byte{})

		var buf bytes.Buffer
		w := bufio.NewWriter(&buf)

		_, reason := p.FileWillBeUploaded(nil, fi, r, w)
		assert.Equal(t, reason, "ERROR: original image is corrupt image: unknown format")
	})

	t.Run("png watermark", func(t *testing.T) {

		setupAPI := func() *plugintest.API {
			api := &plugintest.API{}
			return api
		}

		api := setupAPI()
		defer api.AssertExpectations(t)
		p := &Plugin{}
		p.API = api
		p.configuration = &configuration{WaterMark: "Test Watermark"}

		data, err := ioutil.ReadFile("../assets/test.png")
		assert.Nil(t, err)

		fi := &model.FileInfo{
			Extension: "PNG",
		}

		r := bytes.NewReader(data)

		var buf bytes.Buffer
		w := bufio.NewWriter(&buf)

		_, reason := p.FileWillBeUploaded(nil, fi, r, w)
		assert.Equal(t, reason, "")

		img, _, err := image.Decode(bytes.NewReader(buf.Bytes())) // decoding to golang's image.Image
		assert.Nil(t, err)
		sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

		msg := steganography.Decode(sizeOfMessage, img)
		assert.Equal(t, string(msg), "Test Watermark")
	})

	t.Run("jpg watermark", func(t *testing.T) {

		setupAPI := func() *plugintest.API {
			api := &plugintest.API{}
			return api
		}

		api := setupAPI()
		defer api.AssertExpectations(t)
		p := &Plugin{}
		p.API = api
		p.configuration = &configuration{WaterMark: "Test Watermark"}

		data, err := ioutil.ReadFile("../assets/test.jpg")
		assert.Nil(t, err)

		fi := &model.FileInfo{
			Extension: "JPG",
		}

		r := bytes.NewReader(data)

		var buf bytes.Buffer
		w := bufio.NewWriter(&buf)

		_, reason := p.FileWillBeUploaded(nil, fi, r, w)
		assert.Equal(t, reason, "")

		img, _, err := image.Decode(bytes.NewReader(buf.Bytes())) // decoding to golang's image.Image
		assert.Nil(t, err)
		sizeOfMessage := steganography.GetMessageSizeFromImage(img) // retrieving message size to decode in the next line

		msg := steganography.Decode(sizeOfMessage, img)
		assert.Equal(t, string(msg), "Test Watermark")
	})
}
