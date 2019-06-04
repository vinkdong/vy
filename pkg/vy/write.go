package vy

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/pkg/errors"
	"github.com/vinkdong/gox/log"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"strings"
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

// TODO: 标签对应的值 文档中重复会报错
/*
;非常简单粗暴的方法
;通过json找到需要替换的字符串，直接字符串替换
;可以保留注释
;功能很多欠缺
*/

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

	if _, ok := js.CheckGet(tags[0]); !ok {
		newData := append(data, []byte(fmt.Sprintf("\n%s: %s", w.Tag, w.Value))...)
		if err := ioutil.WriteFile(w.Path, []byte(newData), w.Mode); err != nil {
			return err
		}
		return nil
	}

	jsPath := js.GetPath(tags...)

	var value string

	if jsPath.Interface() == nil {
		return errors.New("cant find request path")
	} else {
		value, err = jsPath.String()
	}

	if err != nil {
		log.Error(err)
		return err
	}
	orgData := string(data[:])

	if strings.Count(orgData, value) == 1 {
		newData := strings.Replace(orgData, value, w.Value, 1)
		if err := ioutil.WriteFile(w.Path, []byte(newData), w.Mode); err != nil {
			return err
		}
	} else {
		log.Info("haven't support now for your document...")
	}
	return nil
}
