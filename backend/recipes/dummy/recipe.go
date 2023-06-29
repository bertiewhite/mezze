package dummy

import "mezze/recipes"

func GetDummyRecipe() recipes.Recipe {
	return recipes.Recipe{
    Name:        "A Dummy Recipe With Nonsence Data",
		Id:          "gid://mezze/Recipe/1234",
		Author:      "Bertie White",
		Description: "A delicious and very real recipe",
    Ingredients: []string{"lime","lemon"},
		Steps: []recipes.RecipeStep{
			{
				Priority:    0,
				Description: "You gotta buy the ingredients",
			},
			{
				Priority:    1,
				Description: "Then you gotta cook the ingredients",
			},
			{
				Priority:    2,
				Description: "Then you gotta serve the meal",
			},
		},
	}
}
