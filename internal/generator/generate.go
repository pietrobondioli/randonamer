package generator

import (
	"path/filepath"
	"randonamer/internal/config"
	"randonamer/internal/grammar"
	"randonamer/internal/util"
)

func GenerateCoolname(cfg config.Config) (string, error) {
	var g grammar.Grammar
	realDataPath := filepath.Join(cfg.DataPath, cfg.Language)
	grammarPath := filepath.Join(realDataPath, cfg.GrammarFile)
	if err := util.AgnosticUnmarshall(grammarPath, &g); err != nil {
		return "", err
	}

	var coolname string
	startPoint := g.Rules[cfg.StartPoint]

	grammar.ProcessGrammar(realDataPath, g, startPoint, &coolname)

	return coolname, nil
}
