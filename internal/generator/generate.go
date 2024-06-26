package generator

import (
	"path/filepath"
	"randonamer/internal/config"
	"randonamer/internal/grammar"
	"randonamer/internal/util"
)

func GenerateCoolname(cfg config.Config) (string, error) {
	var g grammar.Grammar
	grammarPath := filepath.Join(cfg.DataPath, cfg.Language)
	if err := util.AgnosticUnmarshall(grammarPath, cfg.GrammarFile, &g); err != nil {
		return "", err
	}

	var coolname string

	grammar.ProcessGrammar(cfg.DataPath, g, g.Rules["start"], &coolname)

	return coolname, nil
}
