export type TaskStatus = "TODO" | "IN-PROGRESS" | "DONE";

export interface Task {
  readonly id: number;
  description: string;
  status: TaskStatus;
  readonly createdAt: string;
  updatedAt?: string;
}

export type ActionStatus = "SUCCESS" | "FAILURE";

export type TaskMutationResult = {
  id: number;
  status: ActionStatus;
  task?: Task;
};

export interface TaskTrackerDatabase {
  getTasks(filter?: TaskStatus): Promise<Task[]>;
  addTask(description: string): Promise<TaskMutationResult>;
  deleteTask(id: number): Promise<TaskMutationResult>;
  updateTaskDescription(
    id: number,
    description: string
  ): Promise<TaskMutationResult>;
  updateTaskStatus(id: number, status: TaskStatus): Promise<TaskMutationResult>;
}

export type CommandAction =
  | "ADD"
  | "UPDATE"
  | "DELETE"
  | "LIST"
  | "MARK-IN-PROGRESS"
  | "MARK-DONE";

export type Command = {
  action: CommandAction;
  options: (string | number)[];
};
