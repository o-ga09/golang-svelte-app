import {
    PUBLIC_API_ENDPOINT,
    PUBLIC_BUCKET,
    PUBLIC_IDENTITY_POOL_ID,
    PUBLIC_REGION,
    PUBLIC_USER_POOL_CLIENT_ID,
    PUBLIC_USER_POOL_ID,
 } from "$env/static/public"

 export const config = {
    S3: {
        REGION: PUBLIC_REGION, 
        BUCKET: PUBLIC_BUCKET,
    },
    apiGateway: {
        REGION: PUBLIC_REGION,
        URL: PUBLIC_API_ENDPOINT,
    },
    cognito: {
        REGION: PUBLIC_REGION,
        USER_POOL_ID: PUBLIC_USER_POOL_ID,
        APP_CLIENT_ID: PUBLIC_USER_POOL_CLIENT_ID,
        IDENTITY_POOL_ID: PUBLIC_IDENTITY_POOL_ID,
    },
 };