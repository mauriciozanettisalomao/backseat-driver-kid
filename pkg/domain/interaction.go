package domain

// Interaction represents a user or system interaction and holds the logic to format questions and responses
type Interaction struct {
	Preamble Preamble
	Prompts  []*Prompt
}

// Preamble represents the initial part of an interaction
type Preamble struct {
	Context      string
	Instructions string
	Examples     string
}

// Prompt represents a question or a statement
type Prompt struct {
	Input  string
	Output string
}
