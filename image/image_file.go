package image

import (
)

type ImageFile struct {
	absolutePath string;
	hash string;
}

func NewImageFile(path string) ImageFile {
	o := new(ImageFile);
	o.absolutePath = path;
	o.hash = CalculateFileHash(path);
	return o;
}



