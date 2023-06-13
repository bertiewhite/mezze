package recipes

type RecipeStore interface {
	GetRecipe(id string) Recipe
	GetAllRecipes() []Recipe
}
