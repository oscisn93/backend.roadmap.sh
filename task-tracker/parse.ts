enum TaskStatus {
  TODO = 'TODO',
  IN_PROGRESS = 'IN_PROGRESS',
  DONE = 'DONE'
};

interface TaskModel {
    id: number,
    description: string,
    status: TaskStatus,
    createdAt: string | Date,
    updatedAt: string | Date
}

type CommandNames =
  | "add"
  | "update"
  | "delete"
  | "list"
  | "mark-in-progress"
  | "mark-done";

type Command = {
  action: CommandNames;
  options: string[];
};

class Task implements TaskModel {
  private readonly _id: number;
  private _description: string;
  private readonly _createdAt: string;
  private _updatedAt: string;
  constructor(id: number, description: string) {
    this._description = description;
    this._id = id;
  }
  // public getters
  get id() {
    return this._id;
  }
  get description() {
    return this._description;
  }
  get status() {
    return TaskStatus.TODO;
  }
  get createdAt() {
    return this._createdAt;
  }
  get updatedAt() {
    return this._updatedAt;
  }
  udpateDescription(description: string) {
    this._description = description;
  }
}

function *idGenerator() {
  let count = 0;
  while (true) {
    yield count;
    count++;
  }
}

const incrementingId = idGenerator();

export default function parse(cmd: Command) {
  const { action, options } = cmd;
  switch (action) {
    case 'add':
      return function (options: string[]) {
        if (options.length > 1) {
          console.warn('Unused arguments: task-cli add only accepts one argument [description]. Any additional arguments were ignored.')
        }
        const id = incrementingId.next().value!;
        const task = new Task(id, options[0]);
      }
  }
}
