<script>
    import type { PageData } from './types';
    import { getCount, updateCount } from '$lib';

    export let data: PageData;

    const onClick = async () => {
        const cnt = await updateCount();
        data.count = cnt.count;
    };
</script>

<div class="App">
    {#if data.count}<p>You Clicked me {data.count} times.</p>{/if}
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

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
