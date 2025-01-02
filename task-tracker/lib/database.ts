import {
  ActionStatus,
  Task,
  TaskStatus,
  TaskTrackerDatabase,
} from './types';
import {
  JSON_FILE_PATH
} from './constants';

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
    status: "TODO"
  } satisfies Task;
  return task;
}

export function printTask(task: Task): void {
  console.log(`{ id: ${this._id
    }, description: ${this.description
    }, status: ${this.status
    }, createdAt: ${this.createdAt
    }${!!this.updatedAt ? `, updatedAt: ${this.updatedAt} ` : " "
    }}`);
}

function updateTask(
  task: Task,
  description?: string,
  status?: TaskStatus
) {
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
  private tasks: Task[];
  private counter: number;
  private isSparse: boolean;
  constructor() {
    if (!Database.instance) {
      Database.instance = this;
    }
    this.getSavedTasks();
    this.isSparse = false;
    return Database.instance;
  }

  private compressTasksArray(): void {
    if (this.isSparse) {
      return;
    }
    this.tasks = this.tasks.filter(task => !!task);
    this.isSparse = false;
    return;
  }

  private async getSavedTasks(): Promise<Task[]> {
    this.tasks = [];
    const data = await import(JSON_FILE_PATH, {
      with: { type: "json" },
    });
    if (data.default.tasks) {
      for (const task of data.default.tasks as Task[]) {
        this.tasks.push(task);
      }
    }
    const lastIdx = this.tasks.length;
    if (lastIdx === 0) {
      this.counter = 0;
    } else {
      this.counter = this.tasks[lastIdx].id;
    }
    return this.tasks;
  }

  private async saveTasks() {
    try {
      this.compressTasksArray()
      const data = JSON.stringify({ tasks: this.tasks });
      // @ts-ignore: Deno namespace conflicts with tsserver
      await Deno.writeTextFile(JSON_FILE_PATH, data);
    } catch (err) {
      console.error((err as Error).message)
    }
  }

  private getNextID() {
    return this.counter++;
  }

  async addTask(description: string) {
    const id = this.getNextID();
    const task = createTask(id, description);
    this.tasks.push(task);
    return {
      id,
      status: 'SUCCESS' as ActionStatus,
      task: task
    };
  }

  async getTasks(filter: TaskStatus = "TODO") {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    return this.tasks;
  }

  async getTask(id: number) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
      return null;
    }
    if (this.tasks[id].id !== id) {
      const task = this.tasks.find((task: Task) => task.id === id);
      if (!task) {
        return null;
      }
      return task;
    }
    return this.tasks[id];
  }

  async updateTaskDescription(id: number, description: string) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
    return {
        id,
        status: 'FAILURE' as ActionStatus,
      }
    }
    const task = this.tasks[id];
    updateTask(task, description, undefined);
    return {
      id,
      status: 'SUCCESS' as ActionStatus,
      task
    }
  }

  async updateTaskStatus(id: number, status: TaskStatus) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
      return {
        id,
        status: 'FAILURE' as ActionStatus,
      };
    }
    const task = this.tasks[id];
    updateTask(task, undefined, status)
    return {
      id,
      status: 'SUCCESS' as ActionStatus,
      task
    }
  }

  async deleteTask(id: number) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
      return {
        id,
        status: 'FAILURE' as ActionStatus
      }
    }
    delete this.tasks[id];
    return {
      id,
      status: 'SUCCESS' as ActionStatus,
    }
  }
}

