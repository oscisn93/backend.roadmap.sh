export const JSON_FILE_PATH = "./tasks.json";

export const ADD_MISSIING_ARG =
  "task-cli command 'add' requires one positional argument [description] but none was supplied";

export const UPDATE_MISSING_ARGS =
  "task-cli command 'update' requires two positional arguments, however that constraint was not met, Requires: [id: number, description: string].";

export const DELETE_MISSING_ARG =
  "task-cli command 'delete' requires one postional argument [id] but none was supplied.";

export const MARK_IN_PROGRESS_MISSING_ARG =
  "task-cli command 'mark-in-progress' requires one positional argument [id] but none was supplied.";

export const MARK_DONE_MISSING_ARG =
  "task-cli command 'mark-done' requires one positional argument [id] but none was supplied.";
