# Task Tracker

**Descriptions**: A CLI to show recent user activity from the GitHub api created using Go.\

**Project URL**: https://roadmap.sh/projects/task-tracker

### Requirements

- users should be able to provide a valid github username and see the username's user's activity in the terminal.
- Utilize the endpoint https://api.github.com/users/<username>/events
- Display the returned data in the terminal
- NOTE: You must generate a github personal access token with public/default permissions. You will then create a .env file and define insert the token like so:

```sh
GITHUB_PUBLIC_API_TOKEN=[github_auth_token]
```

### Constraints

- All errors must be handled gracefully
- No external http libraries allowed

#### Bonus ####

- Consider adding features like event-type filters, structured formatting, or caching to improve performance, or explore other enpoints to fetch additional publicly available user data.

### API

```bash
github-activity <username>
Output:
- pushed 3 commits to <username>/portfolio
- Opened a new issue on mrdoob/three.js
- Starred kamranahmedse/developer-roadmap
```

### Data Model

The GitHub API accepts the following request:

# TODO: Add data model for reuest #

The corresponding Response body looks like the following:

```json
[
  {
    "id": "22249084947",
    "type": "WatchEvent",
    "actor": {
      "id": 583231,
      "login": "octocat",
      "display_login": "octocat",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "avatar_url": "https://avatars.githubusercontent.com/u/583231?v=4"
    },
    "repo": {
      "id": 1296269,
      "name": "octocat/Hello-World",
      "url": "https://api.github.com/repos/octocat/Hello-World"
    },
    "payload": {
      "action": "started"
    },
    "public": true,
    "created_at": "2022-06-09T12:47:28Z"
  },
  {
    "id": "22249084964",
    "type": "PushEvent",
    "actor": {
      "id": 583231,
      "login": "octocat",
      "display_login": "octocat",
      "gravatar_id": "",
      "url": "https://api.github.com/users/octocat",
      "avatar_url": "https://avatars.githubusercontent.com/u/583231?v=4"
    },
    "repo": {
      "id": 1296269,
      "name": "octocat/Hello-World",
      "url": "https://api.github.com/repos/octocat/Hello-World"
    },
    "payload": {
      "push_id": 10115855396,
      "size": 1,
      "distinct_size": 1,
      "ref": "refs/heads/master",
      "head": "7a8f3ac80e2ad2f6842cb86f576d4bfe2c03e300",
      "before": "883efe034920928c47fe18598c01249d1a9fdabd",
      "commits": [
        {
          "sha": "7a8f3ac80e2ad2f6842cb86f576d4bfe2c03e300",
          "author": {
            "email": "octocat@github.com",
            "name": "Monalisa Octocat"
          },
          "message": "commit",
          "distinct": true,
          "url": "https://api.github.com/repos/octocat/Hello-World/commits/7a8f3ac80e2ad2f6842cb86f576d4bfe2c03e300"
        }
      ]
    },
    "public": true,
    "created_at": "2022-06-07T07:50:26Z"
  }
]
```

