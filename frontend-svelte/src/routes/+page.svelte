<script lang="ts">
	import { onMount } from "svelte";


    type Recipe = {
        id: string
    }

    let recipe:Recipe 

    onMount( async () => {
        recipe = await fetch("http://localhost:8080/recipes?id=213")
            .then((response) => response.json())
            .then((data) => data as Recipe)
    })

</script>

{#if recipe}
<p>{recipe.id}</p>
{:else}
<p> waiting.... </p>
{/if}
