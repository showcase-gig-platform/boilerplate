package cmd

import (
	"fmt"
	"os"

	"github.com/showcase-gig-platform/boilerplate/pkg/boilerplate"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Gen struct {
	Source      *string
	Destination *string
}

func RegisterGen(app *kingpin.Application) {
	cmd := app.Command("gen", "Generate code from boilerplate")
	gen := Gen{
		Source:      cmd.Flag("src", "source").Short('s').Required().ExistingDir(),
		Destination: cmd.Flag("dst", "destination").Short('d').Required().String(),
	}
	cmd.Action(gen.Action)
}

func (g *Gen) Action(ctx *kingpin.ParseContext) error {
	cfg, err := boilerplate.SearchConfig(*g.Source)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	rules, err := inputReplaceRules(cfg)
	if err != nil {
		return fmt.Errorf("failed to input replace plans: %w", err)
	}
	rules.Print()

	plans, err := boilerplate.NewFileGenerationPlans(*g.Source, rules, cfg.IgnorePrefixes)
	if err != nil {
		return fmt.Errorf("failed to generate plans: %w", err)
	}
	generator := boilerplate.Generator{
		SrcPath: *g.Source,
		DstPath: *g.Destination,
		Rules:   rules,
	}
	for _, plan := range plans {
		if err := generator.Generate(plan); err != nil {
			fmt.Println("failed to generate file:", err)
			os.Exit(1)
		}
	}
	return nil
}

func inputReplaceRules(cfg *boilerplate.Config) (boilerplate.ReplaceRules, error) {
	var rules boilerplate.ReplaceRules

	// Input new project name
	fmt.Printf("New project name (example: %s): ", cfg.Project)
	var newProject string
	if _, err := fmt.Scanln(&newProject); err != nil {
		return nil, fmt.Errorf("failed to read project name: %w", err)
	}

	// Input replace targets
	for _, target := range cfg.Targets {
		fmt.Printf("New %s (example: %s): ", target.Name, target.String)
		var val string
		if _, err := fmt.Scanln(&val); err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", target.Name, err)
		}
		rules = append(rules, boilerplate.ReplaceRule{
			Old: target.String,
			New: val,
		})
	}

	// Append new project name rule
	rules = append(rules, boilerplate.NewReplaceRulesFromCamelCase(cfg.Project, newProject)...)

	return rules, nil
}
