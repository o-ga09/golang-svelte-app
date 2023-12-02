import { Cognito, StackContext, use } from "sst/constructs";
import { NoteApi } from "./NoteApiStack";
import { Storage } from "./StorageStack";
import * as iam from 'aws-cdk-lib/aws-iam';

export function Auth({ stack, app }: StackContext) {
    const { api } = use(NoteApi)
    const { bucket } = use(Storage)
    
    const auth = new Cognito(stack,"Auth", {
        login: ["email"],
    });

    auth.attachPermissionsForAuthUsers(stack, [
        api,
        new iam.PolicyStatement({
            actions: ["s3:*"],
            effect: iam.Effect.ALLOW,
            resources: [
                bucket.bucketArn + "/private/${cognito-identity.amazonaws.com:sub}/*"
            ],
        }),
    ]);

    stack.addOutputs({
        Region: app.region,
        UserPoolId: auth.userPoolId,
        UserPoolClientId: auth.userPoolClientId,
        IdentityPoolId: auth.cognitoIdentityPoolId,
    });

    return {
        auth,
    };
}