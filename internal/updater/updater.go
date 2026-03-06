package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const registryURL = "https://raw.githubusercontent.com/tryoutshell/registry/main/registry.json"

type LessonEntry struct {
	Org     string `json:"org"`
	Lesson  string `json:"lesson"`
	Version string `json:"version"`
	URL     string `json:"url"`
}

type Registry struct {
	Lessons []LessonEntry `json:"lessons"`
}

type UpdateInfo struct {
	Org        string
	Lesson     string
	OldVersion string
	NewVersion string
	URL        string
}

func CheckForUpdates() ([]UpdateInfo, error) {
	remote, err := fetchRegistry()
	if err != nil {
		return nil, fmt.Errorf("fetching registry: %w", err)
	}

	local := loadLocalVersions()
	var updates []UpdateInfo

	for _, entry := range remote.Lessons {
		key := entry.Org + "/" + entry.Lesson
		localVersion, exists := local[key]

		if !exists || localVersion != entry.Version {
			updates = append(updates, UpdateInfo{
				Org:        entry.Org,
				Lesson:     entry.Lesson,
				OldVersion: localVersion,
				NewVersion: entry.Version,
				URL:        entry.URL,
			})
		}
	}

	return updates, nil
}

func DownloadUpdates(updates []UpdateInfo) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	lessonsDir := filepath.Join(homeDir, ".local", "share", "tryoutshell", "lessons")

	for _, u := range updates {
		destDir := filepath.Join(lessonsDir, u.Org, u.Lesson)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return fmt.Errorf("creating directory for %s/%s: %w", u.Org, u.Lesson, err)
		}

		if u.URL == "" {
			continue
		}

		resp, err := http.Get(u.URL)
		if err != nil {
			return fmt.Errorf("downloading %s/%s: %w", u.Org, u.Lesson, err)
		}

		data, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return fmt.Errorf("reading %s/%s: %w", u.Org, u.Lesson, err)
		}

		dest := filepath.Join(destDir, "lesson.yaml")
		if err := os.WriteFile(dest, data, 0644); err != nil {
			return fmt.Errorf("writing %s/%s: %w", u.Org, u.Lesson, err)
		}
	}

	if err := saveLocalVersions(updates); err != nil {
		return fmt.Errorf("saving version info: %w", err)
	}

	return nil
}

func fetchRegistry() (*Registry, error) {
	resp, err := http.Get(registryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("registry returned status %d", resp.StatusCode)
	}

	var reg Registry
	if err := json.NewDecoder(resp.Body).Decode(&reg); err != nil {
		return nil, err
	}
	return &reg, nil
}

func localVersionsPath() string {
	if dir, err := os.UserHomeDir(); err == nil {
		return filepath.Join(dir, ".config", "tryoutshell", "versions.json")
	}
	return ".tryoutshell-versions.json"
}

func loadLocalVersions() map[string]string {
	data, err := os.ReadFile(localVersionsPath())
	if err != nil {
		return map[string]string{}
	}

	var versions map[string]string
	if err := json.Unmarshal(data, &versions); err != nil {
		return map[string]string{}
	}
	return versions
}

func saveLocalVersions(updates []UpdateInfo) error {
	versions := loadLocalVersions()
	for _, u := range updates {
		versions[u.Org+"/"+u.Lesson] = u.NewVersion
	}

	dir := filepath.Dir(localVersionsPath())
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(localVersionsPath(), data, 0644)
}
