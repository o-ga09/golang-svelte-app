<script lang="ts">
    import type { PageData } from './$types';
    // import { updateCount } from '$lib';
    // svelteライブラリからonMount関数をインポート
    import { onMount } from 'svelte';
    // $lib/authからcheckLoginStatus関数をインポート（ログイン状態をチェックする）
    import { checkLoginStatus } from '$lib/auth';

    export let data: PageData;

    const onClick = async () => {
        const cnt = {Count: 1};
        // const cnt = await updateCount();
        data.Count = cnt.Count;
    };

    // onMountは、コンポーネントがマウントされたときに実行されるコールバック関数を受け取る
    onMount(async () => {
        // ログイン状態をチェック
        const isLoggedIn = await checkLoginStatus();
        // ログインしていない場合、リダイレクトしない
        if (!isLoggedIn) {
            return;
        }
        // '/notes'ページへリダイレクト
        window.location.href = '/notes';
    });
</script>

<!-- <div class="App">
    {#if data.Count}<p>You Clicked me {data.Count} times.</p>{/if}
    <button on:click={onClick}>Click Me!</button>
</div> -->

<!-- <style>
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
</style> -->
