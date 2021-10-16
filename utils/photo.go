package utils

import (
	"log"

	"github.com/gosexy/exif"
)

func ExtractPhoto(path string) map[string]string {
	parser := exif.New()

	err := parser.Open(path)
	if err != nil {
		log.Fatalf("Error parsing exif:", err)
	}

	return parser.Tags
}