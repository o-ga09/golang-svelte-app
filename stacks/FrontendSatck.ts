import { StackContext, SvelteKitSite, use } from "sst/constructs";
import { NoteApi } from "./NoteApiStack";
import { Storage } from "./StorageStack";
import { Auth } from "./AuthStack";

export function Frontend({ stack,app }:StackContext) {
  const { noteapi } = use(NoteApi)
  const { auth } = use(Auth)
  const { bucket } = use(Storage)

    const site = new SvelteKitSite(stack, "SvelteSite", {
        path: "packages/frontend",
        environment: {
          PUBLIC_API_ENDPOINT: noteapi.url,
          PUBLIC_REGION: app.region,
          PUBLIC_BUCKET: bucket.bucketName,
          PUBLIC_USER_POOL_ID: auth.userPoolId,
          PUBLIC_USER_POOL_CLIENT_ID: auth.userPoolClientId,
          PUBLIC_IDENTITY_POOL_ID: auth.cognitoIdentityPoolId || "",
        },
      });
    
      stack.addOutputs({
        SiteUrl: site.url,
        ApiEndpoint: noteapi.url,
      });
}