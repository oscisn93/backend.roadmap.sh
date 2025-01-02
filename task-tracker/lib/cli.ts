import type { Command, CLI, TaskStatus } from "./types";
import { Database } from "./database";
import { UPDATE_MISSING_ARGS } from "./constants";

class TaskTracker implements CLI {
  private database: Database;
  constructor() {
    this.database = new Database();
  }
  list(filter?: TaskStatus): void {
    throw new Error("Method not implemented.");
  }
  update(id: number, description: string): void {
    throw new Error("Method not implemented.");
  }
  delete(): void {
    throw new Error("Method not implemented.");
  }
  add(): void {
    throw new Error("Method not implemented.");
  }
  markInProgress(): void {
    throw new Error("Method not implemented.");
  }
  markDone(): void {
    throw new Error("Method not implemented.");
  }
  private parseInput(cmd: Command, database: Database) {
    const { action, options } = cmd;
    switch (action) {
      case "ADD":
        if (options.length < 1) {
          console.error(
            "task-cli command 'add' requires one positional argument [description] but none was supplied"
          );
        }
        return database.addTask(options[1]);
      case "DELETE":
        if (options.length < 1) {
          console.error(
            "Must specify a valid taskID in order to delete a task"
          );
        }
        return this.delete(parseInt(options[1]));
      case "UPDATE":
        if (options.length < 2) {
          console.error(UPDATE_MISSING_ARGS);
        }
        let taskID = parseInt(options[0]);
    }
  }
}
