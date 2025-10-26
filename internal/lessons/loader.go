package lessons_pkg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

// GetLessonContent loads a specific lesson file
func GetLessonContent(orgID, lessonID string) (LessonFormat, error) {
	// Construct file path: lessons/<org>/<lesson>.yaml
	filePath := filepath.Join("lessons", orgID, lessonID+".yaml")

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return LessonFormat{}, fmt.Errorf("lesson file not found: %s", filePath)
	}

	// Read file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return LessonFormat{}, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Parse YAML
	var lessonData LessonFormat
	err = yaml.Unmarshal(fileContent, &lessonData)
	if err != nil {
		return LessonFormat{}, fmt.Errorf("error parsing YAML from %s: %w", filePath, err)
	}

	return lessonData, nil
}

// GetLessonContentFromPath loads a lesson from an explicit path (for testing)
func GetLessonContentFromPath(filePath string) (LessonFormat, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return LessonFormat{}, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	var lessonData LessonFormat
	err = yaml.Unmarshal(fileContent, &lessonData)
	if err != nil {
		return LessonFormat{}, fmt.Errorf("error parsing YAML: %w", err)
	}

	return lessonData, nil
}
