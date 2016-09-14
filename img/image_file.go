package img

type ImageInfo struct {
	absolutePath string
	hash         string
}

func NewImageInfo(path string) *ImageInfo {
	o := new(ImageInfo)
	o.absolutePath = path
	o.hash = CalculateFileHash(path)
	return o
}
