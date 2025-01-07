import { ActionStatus, TaskMutationResult } from './types.ts';
// @ts-ignore: Deno std lib
import { assertEquals } from "@std/testing/asserts";
import { CLI } from './cli.ts';

// @ts-ignore: Deno namespace
Deno.test("CLI should take the arguments ['add', 'some task'] in and return an object with a numeric id, a success status message, and the task object that was added. The json file should contain that same object.", async () => {
  const args = ['add', 'some task'];
  const cli = new CLI(args);
  const results = await cli.run() as TaskMutationResult;
  assertEquals(results, {
    id: 0,
    status: "SUCCESS" as ActionStatus,
    task: {
      id: 0,
      description: 'some task',
      createdAt: results.task!.createdAt,
      status: "todo"
    },
  })
  // @ts-ignore: Deno namespace
  const jsonFile = Deno.readTextFileSync('./tasks.json');
  const jsonTasks = JSON.parse(jsonFile).tasks;
  assertEquals(results.task!, jsonTasks[0]);
});

// @ts-ignore: Deno namespace
Deno.test("CLI should allow users to update the task description by passing the arguments ['update', 0, 'new description for some task']", async () => {
  const args = ['update', 0, 'new description for some task'];
  const cli = new CLI(args);
  const results = await cli.run() as TaskMutationResult;
  console.log(results)
  assertEquals(results, {
    id: 0,
    status: "SUCCESS" as ActionStatus,
    task: {
      id: 0,
      description: 'new description for some task',
      createdAt: results.task!.createdAt,
      status: "todo",
      updatedAt: results.task!.updatedAt!
    },
  })
  // @ts-ignore: Deno namespace
  const jsonFile = await Deno.readTextFile('./tasks.json');
  const jsonTasks = JSON.parse(jsonFile).tasks;
  assertEquals(results.task!, jsonTasks[0]);
});
