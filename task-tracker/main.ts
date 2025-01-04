import { parse } from "@std/flags";
import { CommandAction } from "./lib/types.ts";
import { CLI } from "./lib/cli.ts";

function createCommand(args: (string | number)[]){
   const cmdName = args[0] as string;
   const action = cmdName.toUpperCase() as CommandAction;
   args.shift();
   return {
      action: action,
      options: args
   }
}

async function main() {
   const parsedArgs = parse(Deno.args);
   const cmd = createCommand(parsedArgs._);
   const cli = new CLI()
   const results = await cli.run(cmd);
   console.log('output:', results)
}

main();