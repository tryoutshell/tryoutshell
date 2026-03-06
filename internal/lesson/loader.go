package lesson

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"

	"github.com/tryoutshell/tryoutshell/types"
)

func DiscoverLessons(roots ...string) []types.DiscoveredLesson {
	if len(roots) == 0 {
		roots = lessonsSearchPaths()
	}

	var lessons []types.DiscoveredLesson
	seen := map[string]bool{}

	for _, root := range roots {
		found := discoverInDir(root)
		for _, l := range found {
			key := l.OrgID + "/" + l.LessonID
			if !seen[key] {
				seen[key] = true
				lessons = append(lessons, l)
			}
		}
	}

	sort.Slice(lessons, func(i, j int) bool {
		if lessons[i].OrgID != lessons[j].OrgID {
			return lessons[i].OrgID < lessons[j].OrgID
		}
		return lessons[i].LessonID < lessons[j].LessonID
	})

	return lessons
}

func discoverInDir(root string) []types.DiscoveredLesson {
	var lessons []types.DiscoveredLesson

	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		return nil
	}

	orgEntries, err := os.ReadDir(root)
	if err != nil {
		return nil
	}

	for _, orgEntry := range orgEntries {
		if !orgEntry.IsDir() || strings.HasPrefix(orgEntry.Name(), "_") || strings.HasPrefix(orgEntry.Name(), ".") {
			continue
		}

		orgID := orgEntry.Name()
		orgDir := filepath.Join(root, orgID)

		orgMeta := loadOrgMeta(orgDir, orgID)

		lessonEntries, err := os.ReadDir(orgDir)
		if err != nil {
			continue
		}

		for _, lessonEntry := range lessonEntries {
			if lessonEntry.Name() == "meta.yaml" || lessonEntry.Name() == "meta.yml" {
				continue
			}

			lessonDir := filepath.Join(orgDir, lessonEntry.Name())

			if lessonEntry.IsDir() {
				dl := discoverLessonDir(orgID, orgMeta, lessonEntry.Name(), lessonDir)
				if dl != nil {
					lessons = append(lessons, *dl)
				}
				continue
			}

			if strings.HasSuffix(lessonEntry.Name(), ".yaml") || strings.HasSuffix(lessonEntry.Name(), ".yml") {
				lessonID := strings.TrimSuffix(strings.TrimSuffix(lessonEntry.Name(), ".yaml"), ".yml")
				dl := discoverLegacyLesson(orgID, orgMeta, lessonID, filepath.Join(orgDir, lessonEntry.Name()))
				if dl != nil {
					lessons = append(lessons, *dl)
				}
			}
		}
	}

	return lessons
}

func discoverLessonDir(orgID string, orgMeta types.OrgMeta, lessonID, dir string) *types.DiscoveredLesson {
	lessonYAML := filepath.Join(dir, "lesson.yaml")
	slidesPath := filepath.Join(dir, "slides.md")
	exercisesPath := filepath.Join(dir, "exercises.sh")

	hasLessonYAML := fileExists(lessonYAML)
	hasSlides := fileExists(slidesPath)

	if !hasLessonYAML && !hasSlides {
		return nil
	}

	dl := &types.DiscoveredLesson{
		OrgID:    orgID,
		OrgMeta:  orgMeta,
		LessonID: lessonID,
		Dir:      dir,
		HasSlides:  hasSlides,
		HasExercises: fileExists(exercisesPath),
	}

	if hasLessonYAML {
		meta, err := loadLessonMeta(lessonYAML)
		if err != nil {
			log.Printf("Warning: failed to parse %s: %v", lessonYAML, err)
			dl.LessonMeta = types.LessonMeta{
				ID:    lessonID,
				Title: lessonID,
			}
		} else {
			dl.LessonMeta = meta
		}
		if dl.LessonMeta.ID == "" {
			dl.LessonMeta.ID = lessonID
		}
		if dl.LessonMeta.Title == "" {
			dl.LessonMeta.Title = lessonID
		}
	} else {
		dl.LessonMeta = types.LessonMeta{
			ID:    lessonID,
			Title: lessonID,
		}
	}

	return dl
}

func discoverLegacyLesson(orgID string, orgMeta types.OrgMeta, lessonID, filePath string) *types.DiscoveredLesson {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil
	}

	var wrapper struct {
		Metadata struct {
			ID          string   `yaml:"id"`
			Title       string   `yaml:"title"`
			Description string   `yaml:"description"`
			Difficulty  string   `yaml:"difficulty"`
			Duration    string   `yaml:"duration"`
			Author      string   `yaml:"author"`
			Version     string   `yaml:"version"`
			Tags        []string `yaml:"tags"`
		} `yaml:"metadata"`
	}

	if err := yaml.Unmarshal(data, &wrapper); err != nil {
		log.Printf("Warning: failed to parse %s: %v", filePath, err)
		return nil
	}

	meta := types.LessonMeta{
		ID:          wrapper.Metadata.ID,
		Title:       wrapper.Metadata.Title,
		Description: wrapper.Metadata.Description,
		Difficulty:  wrapper.Metadata.Difficulty,
		Duration:    wrapper.Metadata.Duration,
		Author:      wrapper.Metadata.Author,
		Version:     wrapper.Metadata.Version,
		Tags:        wrapper.Metadata.Tags,
	}

	if meta.ID == "" {
		meta.ID = lessonID
	}
	if meta.Title == "" {
		meta.Title = lessonID
	}

	return &types.DiscoveredLesson{
		OrgID:      orgID,
		OrgMeta:    orgMeta,
		LessonID:   lessonID,
		LessonMeta: meta,
		Dir:        filepath.Dir(filePath),
		HasLegacy:  true,
	}
}

func loadLessonMeta(path string) (types.LessonMeta, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return types.LessonMeta{}, err
	}

	var meta types.LessonMeta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return types.LessonMeta{}, fmt.Errorf("parsing %s: %w", path, err)
	}

	return meta, nil
}

func loadOrgMeta(orgDir, orgID string) types.OrgMeta {
	metaPath := filepath.Join(orgDir, "meta.yaml")
	data, err := os.ReadFile(metaPath)
	if err != nil {
		return types.OrgMeta{
			ID:   orgID,
			Name: titleCase(orgID),
		}
	}

	var meta types.OrgMeta
	if err := yaml.Unmarshal(data, &meta); err != nil {
		return types.OrgMeta{
			ID:   orgID,
			Name: titleCase(orgID),
		}
	}

	if meta.ID == "" {
		meta.ID = orgID
	}
	if meta.Name == "" {
		meta.Name = titleCase(orgID)
	}

	return meta
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			runes := []rune(w)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func LoadSlides(dir string) (string, error) {
	path := filepath.Join(dir, "slides.md")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading slides.md: %w", err)
	}
	return string(data), nil
}

func GroupByOrg(lessons []types.DiscoveredLesson) map[string][]types.DiscoveredLesson {
	groups := map[string][]types.DiscoveredLesson{}
	for _, l := range lessons {
		groups[l.OrgID] = append(groups[l.OrgID], l)
	}
	return groups
}

func GetOrgList(lessons []types.DiscoveredLesson) []types.OrgMeta {
	seen := map[string]bool{}
	var orgs []types.OrgMeta
	for _, l := range lessons {
		if !seen[l.OrgID] {
			seen[l.OrgID] = true
			orgs = append(orgs, l.OrgMeta)
		}
	}
	return orgs
}

func FindLesson(lessons []types.DiscoveredLesson, orgID, lessonID string) *types.DiscoveredLesson {
	for i := range lessons {
		if lessons[i].OrgID == orgID && lessons[i].LessonID == lessonID {
			return &lessons[i]
		}
	}
	return nil
}

func lessonsSearchPaths() []string {
	paths := []string{"lessons"}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		paths = append(paths, filepath.Join(homeDir, ".local", "share", "tryoutshell", "lessons"))
	}

	return paths
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
