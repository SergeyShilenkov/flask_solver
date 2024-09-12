package parsers

import (
	"flask_solver/parsers/iparser"
	"flask_solver/parsers/tparser"
	"path/filepath"
)

type Parser interface {
	Parse() ([][4]int, error)
}

func NewParser(path *string) (*Parser, error) {
	var p Parser

	switch ext := filepath.Ext(*path); ext {
	case ".txt":
		p = tparser.New(path)
	case ".png":
		p = iparser.New(path)
	default:
		return nil, &UnknownFileExtension{ext}
	}

	return &p, nil
}
