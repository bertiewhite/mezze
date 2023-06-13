package dummy

import "mezze/recipes"

func GetDummyRecipe() recipes.Recipe {
	return recipes.Recipe{
		Id:          "gid://mezze/Recipe/1234",
		Author:      "Bertie White",
		Description: "A delicious and very real recipe",
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
