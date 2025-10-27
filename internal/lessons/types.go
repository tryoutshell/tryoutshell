package lessons_pkg

type LessonFormat struct {
	Metadata     MetadataType     `yaml:"metadata"`
	Introduction IntroductionType `yaml:"introduction"`
	Steps        []StepType       `yaml:"steps"`
	Conclusion   ConclusionType   `yaml:"conclusion"`
}

type MetadataType struct {
	ID            string         `yaml:"id"`
	Org           string         `yaml:"org"`
	Title         string         `yaml:"title"`
	Description   string         `yaml:"description"`
	Difficulty    string         `yaml:"difficulty"`
	Duration      string         `yaml:"duration"`
	Prerequisites []string       `yaml:"prerequisites"`
	Tags          []string       `yaml:"tags"`
	Author        string         `yaml:"author"`
	Version       string         `yaml:"version"`
	Resources     []ResourceType `yaml:"resources"`
}
type ResourceType struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
	Type  string `yaml:"type"` // "docs", "video", "tutorial", "github"
}

type IntroductionType struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

type StepType struct {
	Type                   string           `yaml:"type"`
	ID                     string           `yaml:"id"`
	Title                  string           `yaml:"title"`
	Content                string           `yaml:"content"`
	Prompt                 string           `yaml:"prompt"`
	Instruction            string           `yaml:"instruction"`
	Example                string           `yaml:"example"`
	PreContent             string           `yaml:"pre_content"`
	PostContent            string           `yaml:"post_content"`
	Validation             ValidationType   `yaml:"validation"`
	AlternativeValidations []ValidationType `yaml:"alternative_validations"`
	AcceptedCommands       []string         `yaml:"accepted_commands"`
	SuccessMsg             string           `yaml:"success_msg"`
	FailMsg                string           `yaml:"fail_msg"`
	Hints                  []HintType       `yaml:"hints"`
	AllowSkip              bool             `yaml:"allow_skip"`
	Timeout                int              `yaml:"timeout"`
	Highlights             []HighlightType  `yaml:"highlights"`
	CodeBlocks             []CodeBlockType  `yaml:"code_blocks"`
	Callouts               []CalloutType    `yaml:"callouts"`
	Diagram                string           `yaml:"diagram"`
	WaitForContinue        bool             `yaml:"wait_for_continue"`

	// For quiz step
	Questions []QuestionType `yaml:"questions"`

	// For challenge step
	Description  string           `yaml:"description"`
	Verification VerificationType `yaml:"verification"`

	// For interview_prep step
	InterviewQuestions []string `yaml:"-"` // Parsed separately
	RecordAnswers      bool     `yaml:"record_answers"`
	ExportFormat       string   `yaml:"export_format"`
}

type ValidationType struct {
	Type            string   `yaml:"type"`
	Pattern         string   `yaml:"pattern"`
	Contains        string   `yaml:"contains"`
	CaseInsensitive bool     `yaml:"case_insensitive"`
	Expected        int      `yaml:"expected"`
	Files           []string `yaml:"files"`
	Path            string   `yaml:"path"`
	Patterns        []string `yaml:"patterns"`
	AnyMatch        bool     `yaml:"any_match"`
	AllMatch        bool     `yaml:"all_match"`
}

type HintType struct {
	Level int    `yaml:"level"`
	Text  string `yaml:"text"`
}

type HighlightType struct {
	Text  string `yaml:"text"`
	Style string `yaml:"style"`
}

type CodeBlockType struct {
	Label    string `yaml:"label"`
	Code     string `yaml:"code"`
	Language string `yaml:"language"`
}

type CalloutType struct {
	Type string `yaml:"type"`
	Text string `yaml:"text"`
}

type QuestionType struct {
	ID          string   `yaml:"id"`
	Question    string   `yaml:"question"`
	Type        string   `yaml:"type"`
	Options     []string `yaml:"options"`
	Answer      int      `yaml:"answer"`
	Explanation string   `yaml:"explanation"`
}

//type QuestionType struct {
//	Type          string   `json:"type"`
//	Title         string   `json:"title"`
//	Description   string   `json:"description"`
//	Questions     []string `json:"questions"`
//	RecordAnswers bool     `json:"record_answers"`
//	ExportFormat  string   `json:"export_format"`
//}

type VerificationType struct {
	Type   string      `yaml:"type"`
	Checks []CheckType `yaml:"checks"`
}

type CheckType struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path"`
	Pattern string `yaml:"pattern"`
	Command string `yaml:"command"`
}

type ConclusionType struct {
	Title   string      `yaml:"title"`
	Content string      `yaml:"content"`
	Badges  []BadgeType `yaml:"badges"`
}

type BadgeType struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`
}

// Custom unmarshaler to handle interview questions as strings
func (s *StepType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type rawStep StepType
	raw := rawStep{}

	if err := unmarshal(&raw); err != nil {
		return err
	}

	*s = StepType(raw)

	// Handle interview_prep questions (can be simple strings)
	if s.Type == "interview_prep" {
		var questions interface{}
		type tempStep struct {
			Questions interface{} `yaml:"questions"`
		}
		temp := tempStep{}
		if err := unmarshal(&temp); err == nil {
			questions = temp.Questions

			// Check if questions are strings
			if qSlice, ok := questions.([]interface{}); ok {
				for _, q := range qSlice {
					if qStr, ok := q.(string); ok {
						s.InterviewQuestions = append(s.InterviewQuestions, qStr)
						s.Questions = append(s.Questions, QuestionType{
							Question: qStr,
						})
					}
				}
			}
		}
	}

	return nil
}
