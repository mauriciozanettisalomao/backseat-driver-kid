package models

// Preamble represents the initial part of an interaction
type Preamble struct {
	Context      string `yaml:"context"`
	Instructions string `yaml:"instructions"`
	Examples     string `yaml:"examples"`
}

// Prompt represents a question or statement
type Prompt struct {
	Input  string `yaml:"input"`
	Output string `yaml:"output"`
}

// Interaction represents a user or system interaction and holds the logic to format questions and responses
type Interaction struct {
	// used to receive the location of the extended knowledge file and storing the content
	ExtendedKnowledgeDir     []string `yaml:"extendedKownledgeDir"`
	ExtendedKnowledgeContent []byte   `yaml:"extendedKnowledgeContent"`
	ExtendedKnowledge        string   `yaml:"extendedKnowledge"`
	// used to receive the location of the prompt file and storing the content
	PromptFile    string `yaml:"promptFile"`
	PromptContent []byte `yaml:"promptContent"`
	// main data to be used in the interaction
	Prompts  []*Prompt `yaml:"prompts"`
	Preamble *Preamble `yaml:"preamble"`
}

// ConfigMapData represents the data section in the ConfigMap
type ConfigMapData struct {
	// main data to be used in the interaction
	Interaction *Interaction `yaml:"interaction"`
}

// ConfigMap represents the yaml structure of the configuration file
type ConfigMap struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Data *ConfigMapData `yaml:"data"`
}
