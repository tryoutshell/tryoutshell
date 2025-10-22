package lessons_pkg

type metadataType struct {
	Id            string   `json:"id"`  // Unique identifier (kebab-case)
	Org           string   `json:"org"` // Organization name
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Difficulty    string   `json:"difficulty"` // beginner, intermediate, advanced
	Duration      string   `json:"duration"`
	Prerequisites []string `json:"prerequisites"`
	Tags          []string `json:"tags"`
	Author        string   `json:"author"`
	Version       string   `json:"version"`
}

// Set expectations and provide context before diving into steps.
type introductionType struct {
	Title   string `json:"title"`   // Introduction header
	Content string `json:"content"` // Overview text (supports Markdown)
}

// Each step is a YAML object in the steps list.
type stepType struct {
	Type                   string                      `json:"type"`
	Id                     string                      `json:"id,omitempty"`
	Title                  string                      `json:"title,omitempty"`
	Description            string                      `json:"description,omitempty"`
	Content                string                      `json:"content,omitempty"`
	Highlights             []highlightType             `json:"highlights,omitempty"`
	Diagram                string                      `json:"diagram,omitempty"`
	WaitForContinue        bool                        `json:"wait_for_continue,omitempty"`
	CodeBlocks             []codeBlockType             `json:"code_blocks,omitempty"`
	Prompt                 string                      `json:"prompt,omitempty"`
	Instruction            string                      `json:"instruction,omitempty"`
	Example                string                      `json:"example,omitempty"`
	Validation             validationType              `json:"validation,omitempty"`
	AlternativeValidations []alternativeValidationType `json:"alternative_validations,omitempty"`
	SuccessMsg             string                      `json:"success_msg,omitempty"`
	FailMsg                string                      `json:"fail_msg,omitempty"`
	Hints                  []hintType                  `json:"hints,omitempty"`
	AllowSkip              bool                        `json:"allow_skip,omitempty"`
	Timeout                int                         `json:"timeout,omitempty"`
	Callouts               []calloutType               `json:"callouts,omitempty"`
	PreContent             string                      `json:"pre_content,omitempty"`
	PostContent            string                      `json:"post_content,omitempty"`
	AcceptedCommands       []string                    `json:"accepted_commands,omitempty"`
	Questions              []interface{}               `json:"questions,omitempty"`
	Verification           verificationType            `json:"verification,omitempty"`
	RecordAnswers          bool                        `json:"record_answers,omitempty"`
	ExportFormat           string                      `json:"export_format,omitempty"`
}

type highlightType struct {
	Text  string `json:"text"`
	Style string `json:"style"`
}
type codeBlockType struct {
	Label    string `json:"label"`
	Code     string `json:"code"`
	Language string `json:"language"`
}
type validationType struct {
	Type            string   `json:"type"`
	Pattern         string   `json:"pattern,omitempty"`
	CaseInsensitive bool     `json:"case_insensitive,omitempty"`
	Files           []string `json:"files,omitempty"`
	Patterns        []string `json:"patterns,omitempty"`
	AnyMatch        bool     `json:"any_match,omitempty"`
	AllMatch        bool     `json:"all_match,omitempty"`
}
type alternativeValidationType struct {
	Type     string `json:"type"`
	Expected int    `json:"expected,omitempty"`
	Contains string `json:"contains,omitempty"`
}
type hintType struct {
	Level int    `json:"level"`
	Text  string `json:"text"`
}
type calloutType struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type verificationType struct {
	Type   string      `json:"type"`
	Checks []checkType `json:"checks"`
}
type checkType struct {
	Type    string `json:"type"`
	Path    string `json:"path"`
	Pattern string `json:"pattern,omitempty"`
}
type conclusionType struct {
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Badges  []badgeType `json:"badges"`
}
type badgeType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// LessonFormat
// Every lesson YAML file has 4 main sections:
type LessonFormat struct {
	Metadata     metadataType     `json:"metadata"`
	Introduction introductionType `json:"introduction"`
	Steps        []stepType       `json:"steps"`
	Conclusion   conclusionType   `json:"conclusion"`
}
