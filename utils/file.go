package utils

import (
	"fmt"
	"mime"
	"path/filepath"

	"github.com/alfin87aa/go-common/constants/char"
	"github.com/alfin87aa/go-common/constants/messages"
)

func GetFileExtension(filename string) (extension string, err error) {
	// Get the file extension
	ext := filepath.Ext(filename)
	if ext == char.Empty {
		return char.Empty, fmt.Errorf(messages.FileExtensionNotFound)
	}

	// Remove the dot
	extension = ext[1:]

	// Get the mime type
	mimeTypes := mime.TypeByExtension("." + extension)
	if mimeTypes == char.Empty {
		return char.Empty, fmt.Errorf(messages.FileExtensionNotFound)
	}

	return extension, nil
}
