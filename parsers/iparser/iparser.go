package iparser

type ImageParser struct {
	path *string
}

func New(path *string) *ImageParser {
	return &ImageParser{path: path}
}

func (p *ImageParser) Parse() ([][4]int, error) {
	return nil, nil
}
