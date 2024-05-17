package generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type Language string

const (
	English    Language = "en"
	Portuguese Language = "pt_br"
)

const DefaultLanguage Language = English

const ProjectRoot = "./"

var DefaultPathToLanguages map[Language]string = map[Language]string{
	English:    "data/en/",
	Portuguese: "data/pt_br/",
}

const GrammarFile = "_grammar.json"

type Config struct {
	Language   Language
	ConfigPath string
}

type Grammar struct {
	Rules map[string]Rule `json:"rules"`
}

type RuleType string

const (
	Terminal    RuleType = "terminal"
	NonTerminal RuleType = "non-terminal"
)

type Rule struct {
	Type      RuleType   `json:"type"`
	Generates [][]string `json:"generates"`
}

const NameSeparator = "-"

func GenerateCoolname(config Config) (string, error) {
	err := ValidateConfig(config)
	if err != nil {
		return "", err
	}

	var basePath string
	if config.ConfigPath != "" {
		basePath = config.ConfigPath
	} else {
		basePath = ProjectRoot + DefaultPathToLanguages[config.Language]
	}

	var grammar Grammar
	if err := LoadJSON(basePath+GrammarFile, &grammar); err != nil {
		return "", err
	}

	var coolname string

	ProcessGrammar(basePath, grammar, grammar.Rules["start"], &coolname)

	return coolname, nil
}

func ProcessGrammar(basePath string, grammar Grammar, currRule Rule, result *string) {
	if currRule.Type == NonTerminal {
		randIndex := rand.Intn(len(currRule.Generates))

		for _, symbol := range currRule.Generates[randIndex] {
			nextRule, ok := grammar.Rules[symbol]

			if ok {
				ProcessGrammar(basePath, grammar, nextRule, result)
			} else if FileExists(basePath + symbol + ".json") {
				var externalRule Rule
				LoadJSON(basePath+symbol+".json", &externalRule)
				ProcessGrammar(basePath, grammar, externalRule, result)
			}
		}
	}

	if currRule.Type == Terminal {
		randIndex := rand.Intn(len(currRule.Generates))

		wordsToAppend := currRule.Generates[randIndex]
		for _, word := range wordsToAppend {
			if *result != "" {
				*result += NameSeparator
			}

			*result += word
		}
	}
}

func LoadJSON[T any](path string, result *T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(result); err != nil {
		return err
	}

	return nil
}

func ValidateConfig(config Config) error {
	if _, ok := DefaultPathToLanguages[config.Language]; !ok {
		return fmt.Errorf("invalid language: %s", config.Language)
	}

	if config.ConfigPath != "" && !FileExists(config.ConfigPath) {
		return fmt.Errorf("config file not found: %s", config.ConfigPath)
	}

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
