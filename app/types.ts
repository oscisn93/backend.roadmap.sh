export type TaskStatus = "todo" | "in-progress" | "done";

export interface Task {
  readonly id: number;
  description: string;
  status: TaskStatus;
  readonly createdAt: string;
  updatedAt?: string;
}

export type ActionStatus = "SUCCESS" | "FAILURE";

export type Nullable<T> = T | null;

export type TaskMutationResult = {
  id: number;
  status: ActionStatus;
  task: Nullable<Task>;
};

export type CommandAction =
  | "add"
  | "update"
  | "delete"
  | "list"
  | "mark-in-progress"
  | "mark-done";

export type Command = {
  action: CommandAction;
  options: (string | number)[];
};

export interface TimestampProvider {
  getCurrentTimestamp: () => Promise<string>;
}
