package recipes

type Recipe struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Author      string       `json:"author"`
	Description string       `json:"description"`
	Ingredients []string     `json:"ingredients"`
	Steps       []RecipeStep `json:"steps"`
}

type RecipeStep struct {
	Priority    int    `json:"-"`
	Description string `json:"description"`
}
