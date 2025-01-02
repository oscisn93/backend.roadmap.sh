import {
  Task,
  TaskStatus
} from './types';
import {
  JSON_FILE_PATH
} from './constants';


type ActionStatus = 'SUCCESS' | 'FAILURE';

interface TaskTrackerDatabase {
  getTasks(filter: TaskStatus): Promise<Task[]>,
  getTask(id: number): Promise<Task | null>,
  addTask(description: string): Promise<ActionStatus>,
  deleteTask(id: number): Promise<ActionStatus>,
  updateTaskStatus(id: number, description: string): Promise<ActionStatus>,
  updateTaskStatus(id: number, status: TaskStatus): Promise<ActionStatus>
}

function currentTimestamp(): string {
  const date = new Date();
  const timestamp = date.toISOString();
  return timestamp;
}

function createTask(id: number, description: string) {
  const task = {
    id,
    description,
    createdAt: currentTimestamp(),
    status: "TODO"
  } satisfies Task;
  return task;
}

function printTask(task: Task){
  return `{ id: ${
    this._id
  }, description: ${
    this.description
  }, status: ${
    this.status
  }, createdAt: ${
    this.createdAt
  }${
    !!this.updatedAt ? `, updatedAt: ${this.updatedAt} ` : " "
  }}`;
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

export class TaskDB implements TaskTrackerDatabase {
  private static instance?: TaskDB;
  private tasks: Task[];
  private counter: number;
  private isSparse: boolean;
  constructor() {
    if (!TaskDB.instance) {
      TaskDB.instance = this;
    }
    this.getSavedTasks();
    this.isSparse = false;
    return TaskDB.instance;
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
    } catch(err) {
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
    return 'SUCCESS' as ActionStatus;
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
      return 'FAILURE';
    }
    const task = this.tasks[id];
    updateTask(task, description, undefined);
    return 'SUCCESS'
  }

  async updateTaskStatus(id: number, status: TaskStatus) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
      return 'FAILURE';
    }
    const task = this.tasks[id];
    updateTask(task, undefined, status)
    return 'SUCCESS'
  }

  async deleteTask(id: number) {
    if (!this.tasks) {
      await this.getSavedTasks();
    }
    if (id >= this.counter) {
        return 'FAILURE'
    }
    delete this.tasks[id];
    return 'SUCCESS';
  }
}

