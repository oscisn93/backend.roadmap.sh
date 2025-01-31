import {
  ActionStatus,
  Task,
  TaskMutationResult,
  TaskStatus,
  TimestampProvider,
} from "./types.ts";

export class Database {
  private static instance?: Database;
  private filename: string;
  private tasks!: Task[];
  private counter!: number;
  private timestampProvider: TimestampProvider;
  constructor(filename: string, timestampProvider: TimestampProvider) {
    if (!Database.instance) {
      Database.instance = this;
    }
    this.filename = filename;
    this.timestampProvider = timestampProvider;
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

  printTask(task: Task): void {
    console.log(
      `{ id: ${task.id}, description: ${task.description}, status: ${task.status}, createdAt: ${task.createdAt}${
        task.updatedAt ? `, updatedAt: ${task.updatedAt} ` : " "
      }}`,
    );
  }

  private async createTask(id: number, description: string): Promise<Task> {
    const task = {
      id,
      description,
      createdAt: await this.timestampProvider.getCurrentTimestamp(),
      status: "todo",
    } satisfies Task;
    return task;
  }

  async addTask(description: string): Promise<TaskMutationResult> {
    await this.getSavedTasks();
    const id = await this.getNextID();
    const task = await this.createTask(id, description);
    this.tasks.push(task);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task,
    };
  }

  async getTasks(statusFilter?: TaskStatus) {
    await this.getSavedTasks();
    if (statusFilter) {
      return this.tasks.filter((task) => task.status === statusFilter);
    }
    return this.tasks;
  }

  private async updateTask(
    task: Task,
    description?: string,
    status?: TaskStatus,
  ) {
    if (description) {
      task.description = description;
      task.updatedAt = await this.timestampProvider.getCurrentTimestamp();
    }
    if (status) {
      task.status = status;
      task.updatedAt = await this.timestampProvider.getCurrentTimestamp();
    }
    return task;
  }

  async updateTaskDescription(id: number, description: string) {
    await this.getSavedTasks();
    if (id >= this.counter) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
        task: null,
      };
    }
    const task = this.tasks[id];
    await this.updateTask(task, description, undefined);
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
        task: null,
      };
    }
    const task = this.tasks[id];
    await this.updateTask(task, undefined, status);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task,
    };
  }

  async deleteTask(id: number) {
    await this.getSavedTasks();
    if (id >= this.counter || id < 0) {
      return {
        id,
        status: "FAILURE" as ActionStatus,
        task: null,
      };
    }
    this.tasks = this.tasks.filter((task) => task.id !== id);
    await this.saveTasks();
    return {
      id,
      status: "SUCCESS" as ActionStatus,
      task: null,
    };
  }
}
