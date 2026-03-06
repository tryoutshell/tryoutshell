package types

type LessonMeta struct {
	ID          string         `yaml:"id" json:"id"`
	Title       string         `yaml:"title" json:"title"`
	Description string         `yaml:"description" json:"description"`
	Author      string         `yaml:"author" json:"author"`
	Tags        []string       `yaml:"tags" json:"tags"`
	Difficulty  string         `yaml:"difficulty" json:"difficulty"`
	Duration    string         `yaml:"duration" json:"duration"`
	Version     string         `yaml:"version" json:"version"`
	Quiz        []QuizQuestion `yaml:"quiz" json:"quiz"`
}

type QuizQuestion struct {
	Question string   `yaml:"question" json:"question"`
	Options  []string `yaml:"options" json:"options"`
	Answer   int      `yaml:"answer" json:"answer"`
	Explain  string   `yaml:"explain" json:"explain"`
}

type OrgMeta struct {
	ID          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Logo        string `yaml:"logo" json:"logo"`
}

type DiscoveredLesson struct {
	OrgID      string
	OrgMeta    OrgMeta
	LessonID   string
	LessonMeta LessonMeta
	Dir        string
	HasSlides  bool
	HasLegacy  bool
	HasExercises bool
}
