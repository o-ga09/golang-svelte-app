import { StackContext, Api, SvelteKitSite, Table } from "sst/constructs";

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
      "GET /hello": "packages/functions/main.go",
    },
  });

  const site = new SvelteKitSite(stack, "SvelteSite", {
    path: "packages/frontend",
    environment: {
      PUBLIC_API_ENDPOINT: api.url,
    },
  });

  stack.addOutputs({
    SiteUrl: site.url,
    ApiEndpoint: api.url,
  });
}
