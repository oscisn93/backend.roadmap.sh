// @ts-ignore: Deno std lib
import { parse } from "@std/flags";
import { CLI } from "./lib/cli.ts";


// @ts-ignore: Deno namespace
const parsedArgs = parse(Deno.args);
const cli = new CLI(parsedArgs._);
const results = await cli.run();
console.log(results);

