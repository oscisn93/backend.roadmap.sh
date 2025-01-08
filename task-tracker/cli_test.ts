// @ts-ignore: Deno std lib
import { assertEquals } from "@std/testing/asserts";
import { Task, ActionStatus, TaskMutationResult } from './types.ts';
import { CLI } from './cli.ts';

type TestInstance = {
  input: (string | number)[],
  expected: TaskMutationResult;
}

type TestConfig = {
  tests: Record<string, TestInstance | TestInstance[]>
}

const testConfig: TestConfig = {
  tests: {
    'add': {
      input: ['add', 'some task'],
      expected: {
        id: 0,
        status: "SUCCESS" as ActionStatus,
        taskInfo:{
          description: 'some task',
          status: 'todo'
        }
      },
    },
    'update': {
      input: ['update', 0, 'new description for some task'],
      expected: {
        id: 0,
        status: "SUCCESS" as ActionStatus,
        taskInfo: {
          description: 'new description for some task',
          status: "todo",
        }
      }
    },
    'mark-in-progress': {
      input: ['mark-in-progress', 0],
      expected: {
        id: 0,
        status: "SUCCESS" as ActionStatus,
        taskInfo: {
          description: 'new description for some task',
          status: "in-progress",
        }
      }
    },
    'mark-done': {
      input: ['mark-done', 0],
      expected: {
        id: 0,
        status: "SUCCESS" as ActionStatus,
        taskInfo: {
          description: 'new description for some task',
          status: "done",
        }
      }
    },
    'delete': {
      input: ['delete', 0],
      expected: {
        id: 0,
        status: "SUCCESS" as ActionStatus,
        taskInfo: null
      },
    }, 
    'list': [
      {
        input: ['add', 'taskOne'],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskOne',
            status: "todo"
          }
        },
      },
      {
        input: ['add', 'taskTwo'],
        expected: {
          id: 1,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskTwo',
            status: "todo"
          }
        }
      },
      {
        input: ['add', 'taskThree'],
        expected: {
          id: 2,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskThree',
            status: "todo"
          }
        }
      },
      {
        input: ['add', 'taskFour'],
        expected: {
          id: 3,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskFour',
            status: "todo"
          }
        }
      },
      {
        input: ['add', 'taskFive'],
        expected: {
          id: 4,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskFive',
            status: "todo"
          }
        }
      }
    ],
    'list-filtered': [
      {
        input: ['mark-done', 0],
        expected: {
          id: 0,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskOne',
            status: "done"
          }
        },
      },
      {
        input: ['mark-done', 1],
        expected: {
          id: 1,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskTwo',
            status: "done"
          }
        }
      },
      {
        input: ['mark-in-progress', 2],
        expected: {
          id: 2,
          status: "SUCCESS" as ActionStatus,
          taskInfo: {
            description: 'taskThree',
            status: "in-progress"
          }
        }
      },
    ],
  }
};

function getFileTasks() {
  // @ts-ignore: Deno namespace
  const jsonFile = Deno.readTextFileSync('./tasks.json');
  const tasks = JSON.parse(jsonFile).tasks;
  return tasks;
}

function mutationResultToTask(res: TaskMutationResult, createdAt: string, updatedAt?: string): Task {
  const task = {
    id: res.id,
    ...res.taskInfo!,
    createdAt
  } as Task;
  if (updatedAt) {
    task['updatedAt'] = updatedAt;
  }
  return task;
}

const cli = new CLI();

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to add tasks and those tasks should be written to the json file", 
  async function() {
    const config  = testConfig.tests.add as TestInstance;
    const args = config.input;
    const results = await cli.runCommand(args) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected);

    const jsonTasks = getFileTasks();
    const tasks = [{
      id: results.id,
      ...expected.taskInfo,
      createdAt: jsonTasks[0].createdAt
    }];

    assertEquals(jsonTasks, tasks);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to update the task description for an existing task']",
  async function() {
    const config  = testConfig.tests.update as TestInstance;
    const args = config.input;
    const results = await cli.runCommand(args) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected)

    const jsonTasks = getFileTasks();
    const tasks = [{
      id: results.id,
      ...expected.taskInfo,
      createdAt: jsonTasks[0].createdAt,
      updatedAt: jsonTasks[0].updatedAt!
    }];

    assertEquals(jsonTasks, tasks);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to mark tasks in-progress",
  async function() {
    const config  = testConfig.tests['mark-in-progress'] as TestInstance;
    const args = config.input;
    const results = await cli.runCommand(args) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected);

    const jsonTasks = getFileTasks();
    const tasks = [{
      id: results.id,
      ...expected.taskInfo,
      createdAt: jsonTasks[0].createdAt,
      updatedAt: jsonTasks[0].updatedAt!
    }];

    assertEquals(jsonTasks, tasks);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to mark tasks as done", 
  async function() {
    const config  = testConfig.tests['mark-done'] as TestInstance;
    const args = config.input;
    const results = await cli.runCommand(args) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected)

    const jsonTasks = getFileTasks();
    const tasks = [{
      id: results.id,
      ...expected.taskInfo,
      createdAt: jsonTasks[0].createdAt,
      updatedAt: jsonTasks[0].updatedAt!
    }];

    assertEquals(jsonTasks, tasks);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to delete existing tasks",
  async function() {
    const config  = testConfig.tests.delete as TestInstance;
    const args = config.input;
    const results = await cli.runCommand(args) as TaskMutationResult;
    const expected = config.expected;

    assertEquals(results, expected)
 
    const jsonTasks = getFileTasks();

    assertEquals(jsonTasks, []);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to list all existing tasks",
  async function() {
    const config = testConfig.tests.list as TestInstance[];
    for (let idx = 0; idx < config.length; idx++) {
      const args = config[idx].input;
      const results = await cli.runCommand(args) as TaskMutationResult;
      const expected = config[idx].expected;

      assertEquals(results, expected as TaskMutationResult)
    }

    const jsonTasks = getFileTasks();
    const tasks: Task[] = config.map((
        instance: TestInstance,
        i: number
      ) => mutationResultToTask(instance.expected, jsonTasks[i].createdAt, jsonTasks[i].updatedAt)
    );

    assertEquals(tasks, jsonTasks);
});

// @ts-ignore: Deno namespace
Deno.test(
  "CLI should allow users to filter listed tasks by status",
  async function() {
    const config = testConfig.tests['list-filtered'] as TestInstance[];
    for (let idx = 0; idx < config.length; idx++) {
      const args = config[idx].input;
      const results = await cli.runCommand(args) as TaskMutationResult;
      const expected = config[idx].expected;

      assertEquals(results, expected as TaskMutationResult)
    }

    const jsonTasks = getFileTasks();
    const configTodos = (testConfig.tests.list as TestInstance[]).slice(3);
    const tasks: Task[] = config.concat(configTodos).map((
        instance: TestInstance,
        i: number
      ) => mutationResultToTask(instance.expected, jsonTasks[i].createdAt, jsonTasks[i].updatedAt)
    );
    const todoTasks = tasks.splice(3, 2);
    const inProgressTasks = [tasks[2]];
    tasks.pop();
    const filteredTestConfig = [
      {
        args: ['list', 'done'],
        expected: jsonTasks.filter((task: Task) => task.status === 'done'),
        actual: tasks
      },
      {
        args: ['list', 'in-progress'],
        expected: [jsonTasks[2]],
        actual: inProgressTasks
      },
      {
        args: ['list', 'todo'],
        expected: jsonTasks.filter((task: Task) => task.status === 'todo'),
        actual: todoTasks
      }
    ];
    
    for (const filterTest of filteredTestConfig) {
      const filteredResult = (await cli.runCommand(filterTest.args)) as Task[];
      assertEquals(filteredResult, filterTest.expected);
      assertEquals(filterTest.actual, filterTest.expected)
    }
    // cleanup
    for (let i = 0; i < 5; i++) {
      await cli.runCommand(['delete', i]);
    }
    
    const cleanTasks = await cli.runCommand(['list']);
    assertEquals(cleanTasks, []);

    const emptyTasks = getFileTasks();
    assertEquals(emptyTasks, [])
  }
)

