// @ts-ignore: Deno std lib provides type definitions
import { assertEquals } from "@std/testing/asserts";
import { Task, ActionStatus, TaskMutationResult } from "./types.ts";
import { CLI } from "./cli.ts";

type TestInstance = {
  input: (string | number)[];
  expected: TaskMutationResult;
};

type TestConfig = {
  tests: Record<string, TestInstance[]>;
};

const testConfig: TestConfig = {
  tests: {
    add: [
      {
        input: ["add", "some task"],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "some task",
            status: "todo",
            createdAt: 'in_due_time'
          },
        },
      },
    ],
    update: [
      {
        input: ["update", 0, "new description for some task"],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "new description for some task",
            status: "todo",
            createdAt: 'in_due_time',
            updatedAt: 'in_due_time'
          },
        },
      },
    ],
    "mark-in-progress": [
      {
        input: ["mark-in-progress", 0],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "new description for some task",
            status: "in-progress",
            createdAt: 'in_due_time',
            updatedAt: 'in_due_time'
          },
        },
      },
    ],
    "mark-done": [
      {
        input: ["mark-done", 0],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "new description for some task",
            status: "done",
            createdAt: "in_due_time",
            updatedAt: 'in_due_time'
          },
        },
      },
    ],
    delete: [
      {
        input: ["delete", 0],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: null,
        },
      },
    ],
    list: [
      {
        input: ["add", "taskOne"],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "taskOne",
            status: "todo",
            createdAt: "in_due_time"
          },
        },
      },
      {
        input: ["add", "taskTwo"],
        expected: {
          id: 1,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 1,
            description: "taskTwo",
            status: "todo",
            createdAt: "in_due_time"
          },
        },
      },
      {
        input: ["add", "taskThree"],
        expected: {
          id: 2,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 2,
            description: "taskThree",
            status: "todo",
            createdAt: "in_due_time"
          },
        },
      },
      {
        input: ["add", "taskFour"],
        expected: {
          id: 3,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 3,
            description: "taskFour",
            status: "todo",
            createdAt: "in_due_time"
          },
        },
      },
      {
        input: ["add", "taskFive"],
        expected: {
          id: 4,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 4,
            description: "taskFive",
            status: "todo",
            createdAt: "in_due_time"
          },
        },
      },
    ],
    "list-filtered": [
      {
        input: ["mark-done", 0],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 0,
            description: "taskOne",
            status: "done",
            createdAt: "in_due_time",
            updatedAt: "in_due_time"
          },
        },
      },
      {
        input: ["mark-done", 1],
        expected: {
          id: 1,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 1,
            description: "taskTwo",
            status: "done",
            createdAt: "in_due_time",
            updatedAt: "in_due_time"
          },
        },
      },
      {
        input: ["mark-in-progress", 2],
        expected: {
          id: 2,
          status: "SUCCESS" as ActionStatus,
          task: {
            id: 2,
            description: "taskThree",
            status: "in-progress",
            createdAt: "in_due_time",
            updatedAt: "in_due_time"
          },
        },
      },
    ],
  },
};

function getFileTasks() {
  // @ts-ignore: Deno namespace
  const jsonFile = Deno.readTextFileSync("./tasks.json");
  const tasks = JSON.parse(jsonFile).tasks;
  return tasks;
}

const testTimestampProvider = {
  getCurrentTimestamp: async () => {
    return "in_due_time";
  },
};

const cli = new CLI(testTimestampProvider);

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to add tasks and those tasks should be written to the json file",
  async function () {
    const config = testConfig.tests.add[0] as TestInstance;
    const args = config.input;
    const results = (await cli.runCommand(args)) as TaskMutationResult;
    const expected = config.expected;
    assertEquals(results, expected);
    const jsonTasks = getFileTasks();
    assertEquals(jsonTasks, [expected.task]);
  },
);

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to update the task description for an existing task']",
  async function () {
    const config = testConfig.tests.update[0] as TestInstance;
    const args = config.input;
    const results = (await cli.runCommand(args)) as TaskMutationResult;
    const expected = config.expected;
    assertEquals(results, expected);
    const jsonTasks = getFileTasks();
    assertEquals(jsonTasks, [expected.task]);
  },
);

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to mark tasks in-progress",
  async function () {
    const config = testConfig.tests["mark-in-progress"][0] as TestInstance;
    const args = config.input;
    const results = (await cli.runCommand(args)) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected);

    const jsonTasks = getFileTasks();
    
    assertEquals(jsonTasks, [expected.task]);
  },
);

// @ts-ignore: Deno namespace
Deno.test("CLI should allow users to mark tasks as done", async function () {
  const config = testConfig.tests["mark-done"][0] as TestInstance;
  const args = config.input;
  const results = (await cli.runCommand(args)) as TaskMutationResult;
  const expected = config.expected;

  assertEquals(results, expected);

  const jsonTasks = getFileTasks();

  assertEquals(jsonTasks, [expected.task]);
});

// @ts-ignore: Deno namespace
Deno.test("CLI should allow users to delete existing tasks", async function () {
  const config = testConfig.tests.delete[0] as TestInstance;
  const args = config.input;
  const results = (await cli.runCommand(args)) as TaskMutationResult;
  const expected = config.expected;

  assertEquals(results, expected);

  const jsonTasks = getFileTasks();

  assertEquals(jsonTasks, []);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to list all existing tasks",
  async function () {
    const config = testConfig.tests.list as TestInstance[];
    for (let idx = 0; idx < config.length; idx++) {
      const args = config[idx].input;
      const results = (await cli.runCommand(args)) as TaskMutationResult;
      const expected = config[idx].expected;

      assertEquals(results, expected as TaskMutationResult);
    }

    const jsonTasks = getFileTasks();
    const tasks: Task[] = config.map((instance: TestInstance) => instance.expected.task!);
    assertEquals(tasks, jsonTasks);
  },
);

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to filter listed tasks by status",
  async function () {
    const config = testConfig.tests["list-filtered"] as TestInstance[];
    for (let idx = 0; idx < config.length; idx++) {
      const args = config[idx].input;
      const results = (await cli.runCommand(args)) as TaskMutationResult;
      const expected = config[idx].expected;

      assertEquals(results, expected as TaskMutationResult);
    }

    const jsonTasks = getFileTasks();
    const configTodos = (testConfig.tests.list as TestInstance[]).slice(3);

    const tasks: Task[] = config
      .concat(configTodos)
      .map((instance: TestInstance) => instance.expected.task!);

    const todoTasks = tasks.splice(3, 2);
    const inProgressTasks = [tasks[2]];
    tasks.pop();

    const filteredTestConfig = [
      {
        args: ["list", "done"],
        expected: jsonTasks.filter((task: Task) => task.status === "done"),
        actual: tasks,
      },
      {
        args: ["list", "in-progress"],
        expected: [jsonTasks[2]],
        actual: inProgressTasks,
      },
      {
        args: ["list", "todo"],
        expected: jsonTasks.filter((task: Task) => task.status === "todo"),
        actual: todoTasks,
      },
    ];

    for (const filterTest of filteredTestConfig) {
      const filteredResult = (await cli.runCommand(filterTest.args)) as Task[];
      assertEquals(filteredResult, filterTest.expected);
      assertEquals(filterTest.actual, filterTest.expected);
    }
    // cleanup
    for (let i = 0; i < 5; i++) {
      await cli.runCommand(["delete", i]);
    }

    const cleanTasks = await cli.runCommand(["list"]);
    assertEquals(cleanTasks, []);

    const emptyTasks = getFileTasks();
    assertEquals(emptyTasks, []);
  },
);
