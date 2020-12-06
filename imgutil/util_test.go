package imgutil_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	util "github.com/tarokent10/go-utils/imgutil"
)

func TestPng2Jpeg(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.png")
	dest := filepath.Join(wd, "testwork", "nanu.png")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Png2Jpeg(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}
func TestPng2Gif(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.png")
	dest := filepath.Join(wd, "testwork", "nanu.png")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Png2Gif(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}
func TestJpeg2Png(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.jpeg")
	dest := filepath.Join(wd, "testwork", "nanu.jpeg")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Jpeg2Png(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}

func TestJpeg2Gif(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.jpeg")
	dest := filepath.Join(wd, "testwork", "nanu.jpeg")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Jpeg2Gif(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}

func TestGif2Png(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.gif")
	dest := filepath.Join(wd, "testwork", "nanu.gif")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Gif2Png(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}

func TestGif2Jpeg(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("err:%s", err.Error())
	}
	src := filepath.Join(wd, "testinputs", "nanu.gif")
	dest := filepath.Join(wd, "testwork", "nanu.gif")
	if err := copy(src, dest); err != nil {
		t.Errorf("err:%s", err.Error())
	}
	if newfpath, err := util.Gif2Jpeg(dest); err != nil {
		t.Errorf("err:%s", err.Error())
	} else {
		if err := delete(newfpath); err != nil {
			println(err.Error())
		}
	}
}

func copy(srcpath, destpath string) error {
	src, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	return err
}

func delete(t string) error {
	return os.Remove(t)
}
