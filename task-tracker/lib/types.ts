export type TaskStatus = "TODO" | "IN_PROGRESS" | "DONE";

export interface Task {
  readonly id: number;
  description: string;
  status: TaskStatus;
  readonly createdAt: string;
  updatedAt?: string;
}

export type CommandAction = "ADD" | "UPDATE" | "DELETE" | "LIST";

export type CommandOptions = "TASK_ID" | "DESCRIPTION" | "STATUS";

export type Command = {
  action: CommandAction;
  options: CommandOptions[];
};

export interface CLI {
  list(): void,
  update(): void,
  delete(): void,
  add(): void,
  markInProgress(): void,
  markDone(): void
}