package parsers

import "fmt"

type UnknownFileExtension struct {
	extension string
}

func (e *UnknownFileExtension) Error() string {
	return fmt.Sprintf("There are %s extension, but [.txt, .png] expected", e.extension)
}
