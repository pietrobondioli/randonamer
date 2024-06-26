package grammar

import (
	"math/rand"
	"randonamer/internal/util"
)

type Grammar struct {
	Rules         map[string]Rule `json:"rules" yaml:"rules"`
	NameSeparator string          `json:"name_separator" yaml:"name_separator"`
}

type RuleType string

const (
	Terminal    RuleType = "terminal"
	NonTerminal RuleType = "non-terminal"
)

type Rule struct {
	Type      RuleType   `json:"type" yaml:"type"`
	Generates [][]string `json:"generates" yaml:"generates"`
}

func ProcessGrammar(dataPath string, grammar Grammar, currRule Rule, result *string) {
	if currRule.Type == NonTerminal {
		randIndex := rand.Intn(len(currRule.Generates))

		for _, symbol := range currRule.Generates[randIndex] {
			nextRule, ok := grammar.Rules[symbol]

			if ok {
				ProcessGrammar(dataPath, grammar, nextRule, result)
			} else if util.FileExists(dataPath + symbol + ".json") {
				var externalRule Rule
				util.AgnosticUnmarshall(dataPath, symbol, &externalRule)
				ProcessGrammar(dataPath, grammar, externalRule, result)
			}
		}
	}

	if currRule.Type == Terminal {
		randIndex := rand.Intn(len(currRule.Generates))

		wordsToAppend := currRule.Generates[randIndex]
		for _, word := range wordsToAppend {
			if *result != "" {
				*result += grammar.NameSeparator
			}

			*result += word
		}
	}
}
