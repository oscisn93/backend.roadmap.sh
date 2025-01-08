import { ActionStatus, Task, TaskStatus, TaskMutationResult } from "./types.ts";

async function currentTimestamp(): Promise<string> {
  const date = new Date();
  const timestamp = date.toISOString();
  return timestamp;
}

async function createTask(id: number, description: string): Promise<Task> {
  const task = {
    id,
    description,
    createdAt: await currentTimestamp(),
    status: "todo",
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

async function updateTask(
  task: Task,
  description?: string,
  status?: TaskStatus,
) {
  if (description) {
    task.description = description;
    task.updatedAt = await currentTimestamp();
  }
  if (status) {
    task.status = status;
    task.updatedAt = await currentTimestamp();
  }
  return task;
}

export class Database {
  private static instance?: Database;
  private filename: string;
  private tasks!: Task[];
  private counter!: number;
  constructor(filename: string) {
    if (!Database.instance) {
      Database.instance = this;
    }
    this.filename = filename;
    return Database.instance;
  }

  private async getSavedTasks(): Promise<Task[]> {
    this.tasks = [];
    // @ts-ignore: Deno namespace conflicts with tsserver
    const jsonFileText = Deno.readTextFileSync(this.filename);
    const data = JSON.parse(jsonFileText);
    if (data && data.tasks) {
      for (const task of data.tasks as Task[]) {
        this.tasks.push(task);
      }
    }
    if (this.tasks.length === 0) {
      this.counter = this.tasks.length;
    } else {
      this.counter = this.tasks[this.tasks.length - 1].id + 1;
    }
    return this.tasks;
  }

  private async saveTasks() {
    try {
      const data = JSON.stringify({ tasks: this.tasks });
      // @ts-ignore: Deno namespace conflicts with tsserver
      await Deno.writeTextFile(this.filename, data);
    } catch (err) {
      console.error((err as Error).message);
    }
  }

  private async getNextID(): Promise<number> {
    return this.counter++;
  }

  async addTask(description: string): Promise<TaskMutationResult> {
    await this.getSavedTasks();
    const id = await this.getNextID();
    const task = await createTask(id, description);
    this.tasks.push(task);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      taskInfo: {
        description: task.description,
        status: task.status,
      },
    };
  }

  async getTasks(statusFilter?: TaskStatus) {
    await this.getSavedTasks();
    console.log(statusFilter)
    if (statusFilter) {
      return this.tasks.filter(
        (task) => task.status === statusFilter
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
        taskInfo: null,
      };
    }
    const task = this.tasks[id];
    await updateTask(task, description, undefined);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      taskInfo: {
        description: task.description,
        status: task.status,
      },
    };
  }

  async updateTaskStatus(id: number, status: TaskStatus) {
    await this.getSavedTasks();
    if (id >= this.counter) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
        taskInfo: null,
      };
    }
    const task = this.tasks[id];
    await updateTask(task, undefined, status);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      taskInfo: {
        description: task.description,
        status: task.status,
      },
    };
  }

  async deleteTask(id: number) {
    await this.getSavedTasks();
    if (id >= this.counter || id < 0) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
        taskInfo: null,
      };
    }
    this.tasks = this.tasks.filter((task) => task.id !== id);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      taskInfo: null,
    };
  }
}
