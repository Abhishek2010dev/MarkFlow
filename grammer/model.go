package grammer

type LanguageToolErrors struct {
	Matches []struct {
		Message      string `json:"message"`
		ShortMessage string `json:"shortMessage"`
		Replacements []struct {
			Value string `json:"value"`
		} `json:"replacements"`
		Offset   int    `json:"offset"`
		Length   int    `json:"length"`
		Sentence string `json:"sentence"`
		Rule     struct {
			ID          string `json:"id"`
			Description string `json:"description"`
		} `json:"rule"`
	} `json:"matches"`
}
