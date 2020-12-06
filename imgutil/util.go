package imgutil

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	jpegconv converter = &jpegConverter{}
	pngconv  converter = &pngConverter{}
	gifconv  converter = &gifConverter{}
)

// Jpeg2Png converts filetype
func Jpeg2Png(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".jpeg", ".png")
	i, err := decode(jpegconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(pngconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

// Jpeg2Gif converts filetype
func Jpeg2Gif(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".jpeg", ".gif")
	i, err := decode(jpegconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(gifconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

// Png2Jpeg converts filetype
func Png2Jpeg(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".png", ".jpeg")
	i, err := decode(pngconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(jpegconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

// Png2Gif converts filetype
func Png2Gif(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".png", ".gif")
	i, err := decode(pngconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(gifconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

// Gif2Jpeg converts filetype
func Gif2Jpeg(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".gif", ".jpeg")
	i, err := decode(gifconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(jpegconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

// Gif2Png converts filetype
func Gif2Png(filepath string) (string, error) {
	newpath := replaceExt(filepath, ".gif", ".png")
	i, err := decode(gifconv, filepath)
	if err != nil {
		return "", err
	}
	if err := encode(pngconv, i, newpath); err != nil {
		return "", err
	}
	if err := os.Remove(filepath); err != nil {
		return "", err
	}
	return newpath, nil
}

func decode(c converter, path string) (image.Image, error) {
	f, openErr := os.Open(path)
	if openErr != nil {
		return nil, openErr
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	img, encErr := c.Decode(f)
	if encErr != nil {
		return nil, encErr
	}
	return img, nil
}

func encode(c converter, i image.Image, path string) error {
	f2, createErr := os.Create(path)
	if createErr != nil {
		return createErr
	}
	if err := c.Encode(f2, i); err != nil {
		return err
	}
	return nil
}
func replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) <= 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

type converter interface {
	Encode(w io.Writer, i image.Image) error
	Decode(r io.Reader) (image.Image, error)
}

type jpegConverter struct {
}

func (c jpegConverter) Encode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, nil)
}
func (c jpegConverter) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

type pngConverter struct {
}

func (c pngConverter) Encode(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}
func (c pngConverter) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

type gifConverter struct {
}

func (c gifConverter) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, nil)
}
func (c gifConverter) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}
