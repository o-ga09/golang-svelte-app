import { Api,StackContext,use,Config } from "sst/constructs";
import { Storage } from "./StorageStack";

export function NoteApi({ stack}: StackContext) {
  const { table,Notes, Users } = use(Storage)

  const AUTH_SECRET_KEY = new Config.Secret(stack,"AUTH_SECRET_KEY")
    const noteapi = new Api(stack, "api", {
        defaults: {
          function: {
            bind: [Notes,Users,table,AUTH_SECRET_KEY],
            runtime: "go",
            
          },
          authorizer: "iam",
        },
        routes: {
          "GET /hello": "packages/functions/cmd/healthcheck/healthcheck.go",
          "GET /counter": "packages/functions/cmd/getcount/getCount.go",
          "POST /counter": "packages/functions/cmd/updatecount/updateCount.go",
          "GET /notes": "packages/functions/cmd/getnotes/main.go",
          "GET /notes/{id}": "packages/functions/cmd/getnotesbyid/main.go",
          "POST /notes": "packages/functions/cmd/createnote/main.go",
          "PUT /notes/{id}": "packages/functions/cmd/updatenote/main.go",
          "DELETE /notes/{id}": "packages/functions/cmd/deletenote/main.go",
        },
      });

      stack.addOutputs({
        ApiEndpoint:noteapi.url,
      });

      return {
        noteapi
      };
}