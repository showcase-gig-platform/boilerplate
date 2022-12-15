package boilerplate

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type ReplaceRule struct {
	Old string
	New string
}

func (rr *ReplaceRule) Replace(target string) string {
	return strings.ReplaceAll(target, rr.Old, rr.New)
}

func (rr *ReplaceRule) Print() {
	fmt.Printf("ReplaceRule: %s -> %s\n", rr.Old, rr.New)
}

type ReplaceRules []ReplaceRule

func (rr *ReplaceRules) Replace(target string) string {
	for _, rule := range *rr {
		target = rule.Replace(target)
	}
	return target
}

func (rr *ReplaceRules) Print() {
	for _, rule := range *rr {
		rule.Print()
	}
}

func NewReplaceRulesFromCamelCase(old, new string) ReplaceRules {
	r := ReplaceRules{
		NewReplaceRule(old, new, AllSmall),
		NewReplaceRule(old, new, strcase.ToSnake),
		NewReplaceRule(old, new, strcase.ToScreamingSnake),
		NewReplaceRule(old, new, strcase.ToKebab),
		NewReplaceRule(old, new, strcase.ToScreamingKebab),
		NewReplaceRule(old, new, strcase.ToCamel),
		NewReplaceRule(old, new, strcase.ToLowerCamel),
	}
	return r
}

func NewReplaceRule(old, new string, replace func(string) string) ReplaceRule {
	return ReplaceRule{
		Old: replace(old),
		New: replace(new),
	}
}

func AllSmall(s string) string {
	return strings.ToLower(strcase.ToCamel(s))
}
