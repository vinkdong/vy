package vy

import (
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"github.com/bitly/go-simplejson"
	"strings"
	"os"
	"github.com/vinkdong/gox/log"
)

type Write struct {
	Tag   string
	Value string
	Path  string
	Mode  os.FileMode
}

func (w *Write) Write() error {
	data, err := ioutil.ReadFile(w.Path)
	if err != nil {
		return err
	}
	return w.readWrite(data)
}

func (w *Write) readWrite(data []byte) error {
	jsonData, err := yaml.ToJSON(data)
	if err != nil {
		return err
	}
	js, err := simplejson.NewJson(jsonData)
	if err != nil {
		return err

	}
	tags := strings.Split(w.Tag, ".")
	value, err := js.GetPath(tags...).String()
	if err != nil {
		return err
	}
	orgData := string(data[:])

	if strings.Count(orgData, value) == 1 {
		newData := strings.Replace(orgData, value, w.Value, 1)
		if err := ioutil.WriteFile(w.Path, []byte(newData), w.Mode); err != nil {
			return err
		}
	}else {
		log.Info("haven't support now for your document...")
	}
	return nil
}
