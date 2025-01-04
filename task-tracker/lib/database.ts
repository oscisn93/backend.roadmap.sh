import {
  ActionStatus,
  Task,
  TaskStatus,
  TaskTrackerDatabase,
} from "./types.ts";
import { JSON_FILE_PATH } from "./constants.ts";

function currentTimestamp(): string {
  const date = new Date();
  const timestamp = date.toISOString();
  return timestamp;
}

function createTask(id: number, description: string): Task {
  const task = {
    id,
    description,
    createdAt: currentTimestamp(),
    status: "TODO",
  } satisfies Task;
  return task;
}

export function printTask(task: Task): void {
  console.log(
    `{ id: ${task.id}, description: ${task.description}, status: ${
      task.status
    }, createdAt: ${task.createdAt}${
      task.updatedAt ? `, updatedAt: ${task.updatedAt} ` : " "
    }}`,
  );
}

function updateTask(task: Task, description?: string, status?: TaskStatus) {
  if (description) {
    task.description = description;
    task.updatedAt = currentTimestamp();
  }
  if (status) {
    task.status = status;
    task.updatedAt = currentTimestamp();
  }
  return task;
}

export class Database implements TaskTrackerDatabase {
  private static instance?: Database;
  private tasks!: Task[];
  private counter!: number;
  constructor() {
    if (!Database.instance) {
      Database.instance = this;
    }
    return Database.instance;
  }

  private compressTasksArray(): void {
    this.tasks = this.tasks.filter((task) => !!task);
    return;
  }

  private async getSavedTasks(): Promise<Task[]> {
    this.tasks = [];
    const data = await import(`.${JSON_FILE_PATH}`, {
      with: { type: "json" },
    });
    if (data.default.tasks) {
      for (const task of data.default.tasks as Task[]) {
        this.tasks.push(task);
      }
    }
    const lastIdx = this.tasks.length - 1;
    if (lastIdx === -1) {
      this.counter = 0;
    } else {
      this.counter = this.tasks[lastIdx].id + 1;
    }
    return this.tasks;
  }

  private async saveTasks() {
    try {
      const data = JSON.stringify({ tasks: this.tasks });
      // @ts-ignore: Deno namespace conflicts with tsserver
      await Deno.writeTextFile(JSON_FILE_PATH, data);
    } catch (err) {
      console.error((err as Error).message);
    }
  }

  private getNextID() {
    return this.counter++;
  }

  async addTask(description: string) {
    await this.getSavedTasks();
    const id = this.getNextID();
    const task = createTask(id, description);
    this.tasks.push(task);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task: task,
    };
  }

  async getTasks(statusFilter?: string) {
    await this.getSavedTasks();
    if (statusFilter) {
      return this.tasks.filter(
        (task) => task.status === statusFilter.toUpperCase(),
      );
    }
    return this.tasks;
  }

  async updateTaskDescription(id: number, description: string) {
    await this.getSavedTasks();
    if (id >= this.counter) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
      };
    }
    const task = this.tasks[id];
    updateTask(task, description, undefined);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task,
    };
  }

  async updateTaskStatus(id: number, status: TaskStatus) {
    await this.getSavedTasks();
    if (id >= this.counter) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
      };
    }
    const task = this.tasks[id];
    updateTask(task, undefined, status);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task,
    };
  }

  async deleteTask(id: number) {
    await this.getSavedTasks();
    if (id >= this.counter) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
      };
    }
    this.tasks = this.tasks.filter((task) => task.id !== id);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
    };
  }
}
