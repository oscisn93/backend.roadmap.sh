// @ts-ignore: Deno std lib
import { parseArgs } from "jsr:@std/cli/parse-args";
import { CLI } from "./cli.ts";

const timestampProvider = {
  getCurrentTimestamp: async () => {
    const date = new Date();
    const timestamp = date.toISOString();
    return timestamp;
  },
};

// @ts-ignore: Deno namespace
const args = parseArgs(Deno.args)._;
const cli = new CLI(timestampProvider);
const cmd = args.join(" ");
const results = await cli.runCommand(args);
console.log("Results for command", '"' + cmd + '"', "is", results);
