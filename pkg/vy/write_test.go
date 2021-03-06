package vy

import (
	"fmt"
	"testing"
)

func TestReadWriteError(t *testing.T) {
	w := Write{
		Tag: "image.tag",
	}

	data := []byte(`
image:
  tag: image_tag_values
c: image_tag_values
`)
	err := w.readWrite(data)
	if err != nil {
		fmt.Println(err)
		t.Error("empty data should not have error")
	}
}

func TestAddTagWrite(t *testing.T) {
	w := Write{
		Tag: "icon.yy",
	}

	data := []byte(`
image:
  tag: image_tag_values
`)
	err := w.readWrite(data)
	if err == nil{
		t.Fatal("should get error")
	}
}
