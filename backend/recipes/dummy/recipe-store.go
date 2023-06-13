package dummy

import "mezze/recipes"

type DummyStore struct{}

func (ds *DummyStore) GetRecipe(id string) recipes.Recipe {
	return GetDummyRecipe()
}

func (ds *DummyStore) GetAllRecipes() []recipes.Recipe {
	return []recipes.Recipe{GetDummyRecipe()}
}
