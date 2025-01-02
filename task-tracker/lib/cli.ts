import type { Command, CLI } from "./types";
import { TaskDB } from "./database";

class TaskTracker implements CLI {
  private database: TaskDB;
  constructor() {
    this.database = new TaskDB();
  }
  list(): void {
    throw new Error("Method not implemented.");
  }
  update(): void {
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
  private parseInput(cmd: Command, database: TaskDB) {
    const { action, options } = cmd;
    switch (action) {
      case "ADD":
        if (options.length < 1) {
          console.error("task-cli command 'add' requires one positional argument [description] but none was supplied");
        }
        return database.addTask(options[1])
      case "DELETE":
        if (options.length < 1) {
          console.error("Must specify a valid taskID in order to delete a task");
        }
        return database.deleteTask(parseInt(options[1]));
      case "UPDATE":
        if (options.length < 2) {
          console.error("task-cli commands 'update' requires two positional arguments, however that constraint was not met. The first argument must be a taskID and te secode")
        }
        switch (options) {

        }
    }
  }
}