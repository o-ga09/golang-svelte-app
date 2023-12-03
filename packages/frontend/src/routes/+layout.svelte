<script lang="ts">
    import  { Amplify, Auth }  from "aws-amplify"
    import '../app.css';
    import { onMount } from "svelte";
    import {config} from "../config";
	import { page } from "$app/stores";
	import { loggedIn } from "$lib/store";
    import { goto } from '$app/navigation'
    import {checkLoginStatus} from "$lib/auth";
    import Loading from "$lib/components/Loading.svelte";
    
    Amplify.configure({
        Auth: {
            mandatorySignIn: true,
            region: config.cognito.REGION,
            userPoolId: config.cognito.USER_POOL_ID,
            identityPoolId: config.cognito.IDENTITY_POOL_ID,
            userPoolWebClientId: config.cognito.APP_CLIENT_ID,
            storage: global.sessionStorage
        },
        Storage: {
            region: config.S3.REGION,
            bucket: config.S3.BUCKET,
            identityPoolId: config.cognito.IDENTITY_POOL_ID,
        },
        API: {
            endpoints: [ 
                {
                    name: "notes",
                    endpoint: config.apiGateway.URL,
                    region: config.apiGateway.REGION,
                },
            ],
        },
    });

    let isLoaded = false;

    // ユーザーを指定のパスにリダイレクトする関数
    function navigateTo(path: string) {
        // goto関数を使用して指定したURLに遷移
        goto(path);
    }
    
    onMount(async () => {
        const isLoggedIn = await checkLoginStatus();
        if (!isLoggedIn) {
            if ($page.url.pathname !== '/signin' && $page.url.pathname !== '/signup') {
                window.location.href = '/signin';
            } 
            loggedIn.set(false);    
        }
        else {
            loggedIn.set(true);
        }
        isLoaded = true;
    })

    function handleLogout() {
        Auth.signOut().then(() => {
            loggedIn.set(false);
            navigateTo('/signin');
        }).catch((err) => {
            console.error (err);
        });
    }
</script>

{#if !isLoaded}
    <Loading />
{:else}
    <div class="flex flex-col h-screen">
        <!-- 全画面の高さを持つflexboxコンテナを作成 -->

        <header class="p-4 bg-gray-800 text-white flex justify-between items-center">
            <!-- ヘッダー部分。背景色、テキスト色、内側の余白、アイテムの配置などのスタイリングが適用されています -->

            <!-- ロゴエリア -->
            <div>
                <!-- <img src="/logo.png" alt="SimpleNote logo" class="h-8" /> -->
                <!-- アプリケーションのロゴを表示 -->
            </div>

            <div class="flex">
                <!-- ナビゲーションボタンのエリア -->

                {#if $loggedIn}
                    <!-- loggedIn変数がtrueの場合、ユーザーがログインしていることを示します -->

                    <!-- ログアウトボタン -->
                    <button on:click={() => handleLogout()}>Log out</button>
                {:else}
                    <!-- ユーザーがログインしていない場合 -->

                    <!-- ログインページへのボタン -->
                    <button class="mr-2" on:click={() => navigateTo('/signin')}>Log In</button>

                    <!-- サインアップページへのボタン -->
                    <button on:click={() => navigateTo('/signup')}>Sign Up</button>
                {/if}
            </div>
        </header>

        <main class="flex-grow p-4">
            <!-- メインコンテンツエリア。flex-growクラスにより、利用可能なスペースを最大限に利用します -->

            <slot></slot>
            <!-- Svelteのスロット機能。このコンポーネントを使用する他のコンポーネントがコンテンツを挿入できる部分 -->
        </main>
    </div>
{/if}