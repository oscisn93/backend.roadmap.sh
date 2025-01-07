# Task Tracker #

**Descriptions**: A simple CLI to keep track of our daily tasks created using TypeScript on Deno.
**Project URL**: https://github.com/oscisn93/backend.roadmap.sh/task-tracker

### Requirements ###

- Users can add, update, and delete tasks
- Updates include marking tasks as in progress or done
- Users should be able to list all tasks
- Users should be able to filter tasks by the done or in progress status

### Constraints ###

- Must use positional arguments for user inputs
- Store tasks in a JSON file in the same directory as the source file
- If a file doesn't exist, the program should not fail: it must create one
- Must interact with the JSON file using the native file system module
- Must not use external libraries (only the standard library)
- Must ensure there are no unhandled errors: should fail gracefully

### API ###

The following demonstrates the available commands and their usage

```bash
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully! (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

### Data Model ###

```typescript
interface Task {
    id: number,
    description: string,
    status: 'TODO' | 'IN_PROGRESS' | 'DONE',
    createdAt: string | Date,
    updatedAt: string | Date
}
```

In order to complete this project, I utilized the following guide: [Building Cross-Plarform CLIs with Deno in 5 Minutes](https://deno.com/blog/build-cross-platform-cli)

