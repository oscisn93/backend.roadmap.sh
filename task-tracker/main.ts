// @ts-ignore: Deno std lib
import { parse } from "@std/flags";
import { CLI } from "./cli.ts";

// @ts-ignore: Deno namespace
const args = parse(Deno.args)._;
const cli = new CLI();
const cmd = args.join(" ");
const results = await cli.runCommand(args);
console.log("Results for command", "\"" + cmd + "\"", "is", results)

