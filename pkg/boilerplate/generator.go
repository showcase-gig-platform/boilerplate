package boilerplate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	SrcPath string
	DstPath string
	Rules   ReplaceRules
}

func (g *Generator) Generate(plan FileGenerationPlan) error {
	oldFilePath := g.SrcPath + string(filepath.Separator) + plan.OldPath
	newFilePath := g.DstPath + string(filepath.Separator) + plan.NewPath
	fmt.Printf("Generating... %s -> %s\n", oldFilePath, newFilePath)

	// Create new dir
	newDir := g.DstPath + string(filepath.Separator) + plan.NewDirPath()
	if f, err := os.Stat(newDir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.MkdirAll(newDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	// Read old file and replace by rules
	oldFileStat, err := os.Stat(oldFilePath)
	if err != nil {
		return fmt.Errorf("failed to get stat file: %w", err)
	}
	b, err := os.ReadFile(oldFilePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	newStr := g.Rules.Replace(string(b))

	if err := os.WriteFile(newFilePath, []byte(newStr), oldFileStat.Mode()); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

type FileGenerationPlan struct {
	OldPath string
	NewPath string
}

func (fgp *FileGenerationPlan) NewDirPath() string {
	return filepath.Dir(fgp.NewPath)
}

type FileGenerationPlans []FileGenerationPlan

func NewFileGenerationPlans(srcPath string, rules ReplaceRules, ignorePrefixes []string) (FileGenerationPlans, error) {
	fgp := FileGenerationPlans{}
	err := filepath.WalkDir(srcPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk dir: %w", err)
		}

		path, err = cleanRelPath(srcPath, path)
		if err != nil {
			return fmt.Errorf("failed to clean relative path: %w", err)
		}

		if d.IsDir() {
			return nil
		}
		for _, prefix := range ignorePrefixes {
			if strings.HasPrefix(path, prefix) {
				return nil
			}
		}
		fgp = append(fgp, FileGenerationPlan{
			OldPath: path,
			NewPath: rules.Replace(path),
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}
	return fgp, nil
}

// Cleanup relative path
// example: ./././boilerplate -> boilerplate
func cleanRelPath(basePath string, path string) (string, error) {
	baseAbsPath, err := filepath.Abs(basePath)
	if err != nil {
		return "", fmt.Errorf("failed to get base absolute path: %w", err)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	relPath, err := filepath.Rel(baseAbsPath, absPath)
	if err != nil {
		return "", fmt.Errorf("failed to get relative path: %w", err)
	}
	return relPath, nil
}
