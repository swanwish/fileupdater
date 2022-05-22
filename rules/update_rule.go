package rules

type UpdateRule struct {
	Exts             []string
	Pattern          string
	GetMatchReplacer func(match []string) *string
}
