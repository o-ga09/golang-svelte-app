import { StackContext, Api, Table, Bucket } from "sst/constructs";

export function API({ stack }: StackContext) {
  const table = new Table(stack, "counter", {
    fields: {
      counter: "string",
    },
    primaryIndex: {partitionKey: "counter"},
  });

  const api = new Api(stack, "api", {
    defaults: {
      function: {
        bind: [table],
        runtime: "go",
      },
    },
    routes: {
      "GET /hello": "packages/functions/cmd/healthcheck/healthcheck.go",
      "GET /counter": "packages/functions/cmd/getcount/getCount.go",
      "POST /counter": "packages/functions/cmd/updatecount/updateCount.go",
    },
  });
}
