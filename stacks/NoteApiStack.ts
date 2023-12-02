import { Api,StackContext,use,Config } from "sst/constructs";
import { Storage } from "./StorageStack";

export function NoteApi({ stack}: StackContext) {
  const { Notes, Users } = use(Storage)

  const AUTH_SECRET_KEY = new Config.Secret(stack,"AUTH_SECRET_KEY")
    const api = new Api(stack, "api", {
        defaults: {
          function: {
            bind: [Notes,Users,AUTH_SECRET_KEY],
            runtime: "go",
            
          },
          authorizer: "iam",
        },
        routes: {
          "GET /notes": "packages/functions/handlers/getnotes/main.go",
          "GET /notes/{id}": "packages/functions/handlers/getnotesbyid/main.go",
          "POST /notes": "packages/functions/handlers/createnote/main.go",
          "PUT /notes/{id}": "packages/functions/handlers/updatenote/main.go",
          "DELETE /notes/{id}": "packages/functions/handlers/deletenote/main.go",
        },
      });

      stack.addOutputs({
        ApiEndpoint:api.url,
      });

      return {
        api
      };
}