import { SSTConfig } from "sst";
import { API } from "./stacks/MyStack";
import { NoteApi } from "./stacks/NoteApiStack";
import { Storage } from "./stacks/StorageStack";
import { Auth } from "./stacks/AuthStack";
import { Frontend } from "./stacks/FrontendSatck";

export default {
  config(_input) {
    return {
      name: "golang-svelte-app",
      region: "ap-northeast-1",
    };
  },
  stacks(app) {
    app.stack(Storage);
    app.stack(API);
    app.stack(NoteApi);
    app.stack(Auth);
    app.stack(Frontend);
  }
} satisfies SSTConfig;
