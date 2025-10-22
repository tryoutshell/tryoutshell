package lessons_pkg

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func GetLessonContent() LessonFormat {
	filePath := "lessons/chainguard/cosign-sign-verify.yaml"
	//filePath := "lessons/other/hello-world.yaml"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	var lessonData LessonFormat
	err = yaml.Unmarshal(fileContent, &lessonData)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	//fmt.Println(orgStruct.Organizations)
	return lessonData
}
