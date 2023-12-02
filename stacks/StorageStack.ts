import { Table, Bucket, StackContext } from "sst/constructs";

export function Storage({ stack}: StackContext) {
    const Notes = new Table(stack, "notes", {
        fields: {
            userId: "string",
            noteId: "string",
        },
        primaryIndex: {partitionKey: "userId", sortKey: "noteId"},
      });
    
      const Users = new Table(stack, "Users", {
        fields: {
            userId: "string",
            created_at: "string",
            updated_at: "string",
            deleted_at: "string"
        },
        primaryIndex: {partitionKey: "userId", sortKey: "userId"},
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
        Notes,
        Users,
        bucket
      };
}