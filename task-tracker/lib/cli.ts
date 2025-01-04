import type { Command, CommandAction } from "./types.ts";
import { Database } from "./database.ts";
import {
  ADD_MISSIING_ARG,
  DELETE_MISSING_ARG,
  MARK_DONE_MISSING_ARG,
  MARK_IN_PROGRESS_MISSING_ARG,
  UPDATE_MISSING_ARGS,
} from "./constants.ts";

export class CLI {
  private database: Database;
  private command: Command;
  constructor(args: (string | number)[]) {
    this.database = new Database();
    this.command = this.createCommand(args);
  }
  private createCommand(args: (string | number)[]) {
    const cmdName = args[0] as string;
    const action = cmdName.toUpperCase() as CommandAction;
    args.shift();
    return {
      action: action,
      options: args,
    };
  }

  run() {
    const {
      action,
      options
    } = this.command;
    switch (action) {
      case "ADD":
        if (options.length < 1) {
          console.error(ADD_MISSIING_ARG);
        }
        return this.database.addTask(options[0] as string);
      case "DELETE":
        if (options.length < 1) {
          console.error(DELETE_MISSING_ARG);
        }
        return this.database.deleteTask(options[0] as number);
      case "UPDATE":
        if (options.length < 2) {
          console.error(UPDATE_MISSING_ARGS);
        }
        return this.database.updateTaskDescription(
          options[0] as number,
          options[1] as string,
        );
      case "MARK-IN-PROGRESS":
        if (options.length < 1) {
          console.error(MARK_IN_PROGRESS_MISSING_ARG);
        }
        return this.database.updateTaskStatus(
          options[0] as number,
          "IN-PROGRESS",
        );
      case "MARK-DONE":
        if (options.length < 1) {
          console.error(MARK_DONE_MISSING_ARG);
        }
        return this.database.updateTaskStatus(options[0] as number, "DONE");
      case "LIST":
        if (options.length < 1) {
          return this.database.getTasks();
        }
        return this.database.getTasks(options[0] as string);
    }
  }
}
