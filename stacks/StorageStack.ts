import { Table, Bucket, StackContext } from "sst/constructs";

export function Storage({ stack}: StackContext) {
  const table = new Table(stack, "counter", {
    fields: {
      counter: "string",
    },
    primaryIndex: {partitionKey: "counter"},
  });  
  
  const Notes = new Table(stack, "notes", {
        fields: {
            userId: "string",
            noteId: "string",
        },
        primaryIndex: {partitionKey: "userId", sortKey: "noteId"},
      });
    
      const Users = new Table(stack, "users", {
        fields: {
            userId: "string",
        },
        primaryIndex: {partitionKey: "userId"},
      });

      const bucket = new Bucket(stack, "Uploads", {
        cors: [
          {
            maxAge: "1 day",
            allowedOrigins: ["*"],
            allowedHeaders: ["*"],
            allowedMethods: ["GET","PUT","POST","DELETE","HEAD"],
          },
        ],
      });

      return {
        table,
        Notes,
        Users,
        bucket
      };
}