<script lang="ts">
	import { onMount } from "svelte";

    type Recipe = {
        name: string 
        id: string
        author: string
        ingredients: string[]
        steps: { description: string }[]
    }

    let recipe:Recipe 

    onMount( async () => { 
        recipe = await fetch("http://localhost:8080/recipes?id=213")
            .then((response) => response.json())
            .then((data) => data as Recipe)
    })

</script>

{#if recipe}
<div>
    <p>Name: {recipe.name}</p>
    <div>
      <h3>Method</h3>
      <ol>
        {#each recipe.steps as step}
            <li>{step.description}</li>
        {/each}
      </ol>
      <h3>Ingredients</h3>
      <ul>
        {#each recipe.ingredients as i}
          <li>{i}</li>
        {/each}
      </ul>
        
    </div>
</div>

{:else}
<p> waiting.... </p>
{/if}
