package lessons_pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type LessonMetadata struct {
	ID            string   `yaml:"id"`
	Org           string   `yaml:"org"`
	Title         string   `yaml:"title"`
	Description   string   `yaml:"description"`
	Difficulty    string   `yaml:"difficulty"`
	Duration      string   `yaml:"duration"`
	Prerequisites []string `yaml:"prerequisites"`
	Tags          []string `yaml:"tags"`
}

type LessonMetadataFile struct {
	Metadata LessonMetadata `yaml:"metadata"`
}

// GetLessonMetadata loads only the metadata portion of a lesson file
func GetLessonMetadata(orgID, lessonID string) (LessonMetadata, error) {
	filePath := filepath.Join("lessons", orgID, lessonID+".yaml")

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return LessonMetadata{}, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	var lessonFile LessonMetadataFile
	err = yaml.Unmarshal(fileContent, &lessonFile)
	if err != nil {
		return LessonMetadata{}, fmt.Errorf("error parsing YAML: %w", err)
	}

	return lessonFile.Metadata, nil
}

// GetAllLessonMetadata loads metadata for all lessons of an org
func GetAllLessonMetadata(orgID string, lessonIDs []string) []LessonMetadata {
	var metadata []LessonMetadata

	for _, lessonID := range lessonIDs {
		meta, err := GetLessonMetadata(orgID, lessonID)
		if err != nil {
			// If can't load metadata, use lesson ID as title
			metadata = append(metadata, LessonMetadata{
				ID:    lessonID,
				Title: lessonID,
			})
			continue
		}
		metadata = append(metadata, meta)
	}

	return metadata
}
