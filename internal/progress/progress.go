package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LessonProgress struct {
	Completed   bool      `json:"completed"`
	QuizScore   int       `json:"quiz_score"`
	QuizTotal   int       `json:"quiz_total"`
	LastAccess  time.Time `json:"last_access"`
	TimeSpentMs int64     `json:"time_spent_ms"`
	SlideIndex  int       `json:"slide_index"`
}

type Store struct {
	Lessons map[string]LessonProgress `json:"lessons"`
	path    string
}

func NewStore() *Store {
	s := &Store{
		Lessons: make(map[string]LessonProgress),
		path:    progressFilePath(),
	}
	s.load()
	return s
}

func (s *Store) key(orgID, lessonID string) string {
	return orgID + "/" + lessonID
}

func (s *Store) MarkComplete(orgID, lessonID string) {
	k := s.key(orgID, lessonID)
	p := s.Lessons[k]
	p.Completed = true
	p.LastAccess = time.Now()
	s.Lessons[k] = p
	s.save()
}

func (s *Store) SaveQuizScore(orgID, lessonID string, score, total int) {
	k := s.key(orgID, lessonID)
	p := s.Lessons[k]
	p.QuizScore = score
	p.QuizTotal = total
	p.LastAccess = time.Now()
	s.Lessons[k] = p
	s.save()
}

func (s *Store) RecordAccess(orgID, lessonID string, durationMs int64) {
	k := s.key(orgID, lessonID)
	p := s.Lessons[k]
	p.LastAccess = time.Now()
	p.TimeSpentMs += durationMs
	s.Lessons[k] = p
	s.save()
}

func (s *Store) SaveSlideIndex(orgID, lessonID string, idx int) {
	k := s.key(orgID, lessonID)
	p := s.Lessons[k]
	p.SlideIndex = idx
	p.LastAccess = time.Now()
	s.Lessons[k] = p
	s.save()
}

func (s *Store) GetProgress(orgID, lessonID string) LessonProgress {
	return s.Lessons[s.key(orgID, lessonID)]
}

func (s *Store) GetAllProgress() map[string]LessonProgress {
	return s.Lessons
}

func (s *Store) ResetProgress() {
	s.Lessons = make(map[string]LessonProgress)
	s.save()
}

func (s *Store) StatusIcon(orgID, lessonID string) string {
	p := s.GetProgress(orgID, lessonID)
	if p.Completed {
		return "✓"
	}
	if !p.LastAccess.IsZero() {
		return "⟳"
	}
	return "◌"
}

func (s *Store) QuizLabel(orgID, lessonID string) string {
	p := s.GetProgress(orgID, lessonID)
	if p.QuizTotal > 0 {
		return fmt.Sprintf("[%d/%d]", p.QuizScore, p.QuizTotal)
	}
	return ""
}

func (s *Store) load() {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return
	}
	_ = json.Unmarshal(data, s)
}

func (s *Store) save() {
	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(s.path, data, 0644)
}

func progressFilePath() string {
	if dir, err := os.UserHomeDir(); err == nil {
		return filepath.Join(dir, ".config", "tryoutshell", "progress.json")
	}
	return filepath.Join(".", ".tryoutshell-progress.json")
}
