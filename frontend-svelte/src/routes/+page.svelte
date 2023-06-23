<script lang="ts">
	import { onMount } from "svelte";


    type Recipe = {
        id: string
        author: string
        method: { description: string }[]
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
    <p>Recipe ID: {recipe.id}</p>
    <ol>
        {#each recipe.method as step}
            <li>{step.description}</li>
        {/each}
    </ol>
</div>

{:else}
<p> waiting.... </p>
{/if}
