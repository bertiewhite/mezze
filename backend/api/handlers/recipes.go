package handlers

import (
	"encoding/json"
	"mezze/recipes"
	"net/http"
)

type RecipesHandlers struct {
	recipeStore recipes.RecipeStore
}

func NewRecipesHandlers(recipeStore recipes.RecipeStore) *RecipesHandlers {
	return &RecipesHandlers{recipeStore: recipeStore}
}

func (rh *RecipesHandlers) GetHandlers() Handlers {
	return map[string]func(w http.ResponseWriter, r *http.Request){
		"/recipes": rh.getRecipeHanlder(),
	}
}

type AllRecipesResponse struct {
	Recipes []recipes.Recipe `json:"recipes"`
}

func (rh *RecipesHandlers) getRecipeHanlder() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		queryParams := r.URL.Query()

		if len(queryParams) > 1 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("only at most 1 query param, id, is accepted"))
			return
		}

		recipeIds, present := queryParams["id"]
		if !present {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			recipesResponse := &AllRecipesResponse{
				Recipes: rh.recipeStore.GetAllRecipes(),
			}

			json.NewEncoder(w).Encode(recipesResponse)
			return
		}

		var recipeId string
		if len(recipeIds) == 1 {
			recipeId = recipeIds[0]
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Only supports the passing of exactly 1 recipe id currently"))
			return
		}

		recipe := rh.recipeStore.GetRecipe(recipeId)

		json.NewEncoder(w).Encode(recipe)
		w.WriteHeader(http.StatusOK)
		return
	}
}
