<script lang="ts">
    import type { PageData } from './$types';
    import { updateCount } from '$lib';

    export let data: PageData;

    const onClick = async () => {
        const cnt = await updateCount();
        data.Count = cnt.Count;
    };
</script>

<div class="App">
    {#if data.Count}<p>You Clicked me {data.Count} times.</p>{/if}
    <button on:click={onClick}>Click Me!</button>
</div>

<style>
    .App {
        text-align: center;
    }
    p {
        margin-top: 0;
        font-size: 20px;
    }
    button {
        font-size: 48px;
    }
</style>
