package rules

type UpdateRule struct {
	Suffixes         []string
	Pattern          string
	GetMatchReplacer func(match []string) *string
}
