import type { CommandAction, TimestampProvider, TaskStatus } from "./types.ts";
import { Database } from "./database.ts";

export const ADD_MISSIING_ARG =
  "task-cli command 'add' requires one positional argument [description] but none was supplied";
export const UPDATE_MISSING_ARGS =
  "task-cli command 'update' requires two positional arguments, however that constraint was not met, Requires: [id: number, description: string].";
export const DELETE_MISSING_ARG =
  "task-cli command 'delete' requires one postional argument [id] but none was supplied.";
export const MARK_IN_PROGRESS_MISSING_ARG =
  "task-cli command 'mark-in-progress' requires one positional argument [id] but none was supplied.";
export const MARK_DONE_MISSING_ARG =
  "task-cli command 'mark-done' requires one positional argument [id] but none was supplied.";

export class CLI {
  private database: Database;
  constructor(timestampProvider: TimestampProvider) {
    const filename = "./tasks.json";
    this.database = new Database(filename, timestampProvider);
  }

  private createCommand(args: (string | number)[]) {
    const action = args[0] as CommandAction;
    args.shift();
    return {
      action: action,
      options: args,
    };
  }

  runCommand(args: (string | number)[]) {
    const { action, options } = this.createCommand(args);
    switch (action) {
      case "add":
        if (options.length < 1) {
          console.error(ADD_MISSIING_ARG);
        }
        return this.database.addTask(options[0] as string);
      case "delete":
        if (options.length < 1) {
          console.error(DELETE_MISSING_ARG);
        }
        return this.database.deleteTask(options[0] as number);
      case "update":
        if (options.length < 2) {
          console.error(UPDATE_MISSING_ARGS);
        }
        return this.database.updateTaskDescription(
          options[0] as number,
          options[1] as string,
        );
      case "mark-in-progress":
        if (options.length < 1) {
          console.error(MARK_IN_PROGRESS_MISSING_ARG);
        }
        return this.database.updateTaskStatus(
          options[0] as number,
          "in-progress",
        );
      case "mark-done":
        if (options.length < 1) {
          console.error(MARK_DONE_MISSING_ARG);
        }
        return this.database.updateTaskStatus(options[0] as number, "done");
      case "list":
        if (options.length < 1) {
          return this.database.getTasks();
        }
        return this.database.getTasks(options[0] as TaskStatus);
    }
  }
}
