-- name: GetActor :one
SELECT
  *
FROM
  actor
WHERE
  id = ?
LIMIT
  1;

-- name: ListActors :many
SELECT
  *
FROM
  actor
ORDER by
  login;

-- name: AddActor :one
INSERT INTO
  actor (
    login,
    display_login,
    gravatar_id,
    url,
    avatar_url
  )
VALUES
  (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateActor :one
UPDATE actor
SET
  login = ?,
  display_login = ?,
  gravatar_id = ?,
  url = ?,
  avatar_url = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteActor :exec
DELETE FROM actor
WHERE
  id = ?;

-- name: GetUserRequest :one
SELECT
  *
FROM
  user_request
WHERE
  username = ?
LIMIT 1;

-- name: CreateUserRequest :one
INSERT INTO
  user_request (
    usename,
    actor_id,
    etag,
    rate_limit
  ) VALUES (
    ?, ?, ?, ?
  ) RETURNING *;

-- name: UpdateUserRequest :one
UPDATE user_request
  SET etag = ?,
      rate_limit = ?
  WHERE id = ?
RETURNING *;

-- name: GetRepo :one
SELECT
  *
FROM
  repo
WHERE
  id = ?
LIMIT
  1;

-- name: ListRepos :many
SELECT
  *
FROM
  repo
ORDER by
  name;

-- name: AddRepo :one
INSERT INTO
  repo (name, url)
VALUES
  (?, ?) RETURNING *;

-- name: UpdateRepo :one
UPDATE repo
SET
  name = ?,
  url = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteRepo :exec
DELETE FROM repo
WHERE
  id = ?;

-- name: GetPush :one
SELECT
  *
FROM
  push
WHERE
  id = ?
LIMIT
  1;

-- name: ListPushes :many
SELECT
  *
FROM
  push
ORDER BY
  repository_id;

-- name: AddPush :one
INSERT INTO
  push (
    push_id,
    repository_id,
    size,
    distinct_size,
    ref,
    head,
    before
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: DeletePush :exec
DELETE FROM push
WHERE
  id = ?;

-- name: GetPullRequest :one
SELECT
  *
FROM
  pull_request
WHERE
  id = ?
LIMIT
  1;

-- name: GetPullRequests :many
SELECT
  *
FROM
  pull_request
ORDER BY
  user_id,
  updated_at;

-- name: GetPRsByUserID :many
SELECT
  *
FROM
  pull_request
WHERE
  user_id = ?
ORDER BY
  updated_at;

-- name: AddPullRequest :one
INSERT INTO
  pull_request (
    action,
    number,
    pull_request_id,
    state,
    title,
    user_id,
    body,
    created_at,
    updated_at,
    closed_at,
    merged_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdatePullRequest :one
UPDATE pull_request
SET
  action = ?,
  number = ?,
  pull_request_id = ?,
  state = ?,
  title = ?,
  user_id = ?,
  body = ?,
  updated_at = ?,
  closed_at = ?,
  merged_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeletePullRequest :exec
DELETE FROM pull_request
WHERE
  id = ?;

-- name: GetIssue :one
SELECT
  *
FROM
  issue
WHERE
  id = ?
LIMIT
  1;

-- name: GetIssues :many
SELECT
  *
FROM
  issue
ORDER BY
  user_id,
  updated_at;

-- name: GetIssuesByUserID :many
SELECT
  *
FROM
  issue
WHERE
  user_id = ?
ORDER BY
  updated_at;

-- name: AddIssue :one
INSERT INTO
  issue (
    action,
    issue_id,
    number,
    state,
    title,
    user_id,
    body,
    created_at,
    updated_at,
    closed_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateIssue :one
UPDATE issue
SET
  action = ?,
  number = ?,
  issue_id = ?,
  state = ?,
  title = ?,
  user_id = ?,
  body = ?,
  updated_at = ?,
  closed_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteIssue :exec
DELETE FROM issue
WHERE
  id = ?;

-- name: GetFork :one
SELECT
  *
FROM
  fork
WHERE
  id = ?
LIMIT
  1;

-- name: GetForks :many
SELECT
  *
FROM
  fork
ORDER BY
  forkee_name;

-- name: AddFork :one
INSERT INTO
  fork (
    forkee_id,
    forkee_name,
    forkee_full_name,
    forkee_private
  )
VALUES
  (?, ?, ?, ?) RETURNING *;

-- name: UpdateFork :one
UPDATE fork
SET
  forkee_id = ?,
  forkee_name = ?,
  forkee_full_name = ?,
  forkee_private = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteFork :exec
DELETE FROM fork
WHERE
  id = ?;

-- name: GetWatch :one
SELECT
  *
FROM
  watch
WHERE
  id = ?
LIMIT
  1;

-- name: GetWatches :many
SELECT
  *
FROM
  watch;

-- name: AddWatch :one
INSERT INTO
  watch (action)
VALUES
  (?) RETURNING *;

-- name: UpdateWatch :one
UPDATE watch
SET
  action = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteWatch :exec
DELETE FROM watch
WHERE
  id = ?;

-- name: GetRefCreate :one
SELECT
  *
FROM
  ref_create
WHERE
  id = ?
LIMIT
  1;

-- name: GetRefCreates :many
SELECT
  *
FROM
  ref_create
ORDER BY
  pusher_type;

-- name: AddRefCreate :one
INSERT INTO
  ref_create (
    ref,
    ref_type,
    master_branch,
    description,
    pusher_type
  )
VALUES
  (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateRefCreate :one
UPDATE ref_create
SET
  ref = ?,
  ref_type = ?,
  master_branch = ?,
  description = ?,
  pusher_type = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteRefCreate :exec
DELETE FROM ref_create
WHERE
  id = ?;

-- name: GetRefDelete :one
SELECT
  *
FROM
  ref_delete
WHERE
  id = ?
LIMIT
  1;

-- name: GetRefDeletes :many
SELECT
  *
FROM
  ref_delete
ORDER BY
  ref_type;

-- name: AddRefDelete :one
INSERT INTO
  ref_delete (ref, ref_type)
VALUES
  (?, ?) RETURNING *;

-- name: UpdateRefDelete :one
UPDATE ref_delete
SET
  ref = ?,
  ref_type = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteRefDelete :exec
DELETE FROM ref_delete
WHERE
  id = ?;

-- name: GetRelease :one
SELECT
  *
FROM
  release
WHERE
  id = ?
LIMIT
  1;

-- name: GetReleases :many
SELECT
  *
FROM
  release
ORDER BY
  name;

-- name: AddRelease :one
INSERT INTO
  release (
    action,
    release_id,
    tag_name,
    target_commitish,
    name,
    body,
    prerelease,
    created_at,
    published_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateRelease :one
UPDATE release
SET
  action = ?,
  release_id = ?,
  tag_name = ?,
  target_commitish = ?,
  name = ?,
  body = ?,
  prerelease = ?,
  published_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteRelease :exec
DELETE FROM release
WHERE
  id = ?;

-- name: GetIssueComment :one
SELECT
  *
FROM
  issue_comment
WHERE
  id = ?
LIMIT
  1;

-- name: GetIssueComments :many
SELECT
  *
FROM
  issue_comment
ORDER BY
  created_at DESC;

-- name: AddIssueComment :one
INSERT INTO
  issue_comment (
    action,
    issue_id,
    comment_id,
    comment_user_id,
    body,
    created_at,
    updated_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateIssueComment :one
UPDATE issue_comment
SET
  action = ?,
  issue_id = ?,
  comment_id = ?,
  comment_user_id = ?,
  body = ?,
  created_at = ?,
  updated_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteIssueComment :exec
DELETE FROM issue_comment
WHERE
  id = ?;

-- name: GetPullRequestReview :one
SELECT
  *
FROM
  pull_request_review
WHERE
  id = ?
LIMIT
  1;

-- name: GetPullRequestReviews :many
SELECT
  *
FROM
  pull_request_review
ORDER BY
  submitted_at DESC;

-- name: AddPullRequestReview :one
INSERT INTO
  pull_request_review (
    action,
    review_id,
    pull_request_id,
    user_id,
    body,
    submitted_at
  )
VALUES
  (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdatePullRequestReview :one
UPDATE pull_request_review
SET
  action = ?,
  review_id = ?,
  pull_request_id = ?,
  user_id = ?,
  body = ?,
  submitted_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeletePullRequestReview :exec
DELETE FROM pull_request_review
WHERE
  id = ?;

-- Queries for pr_review_comment table
-- name: GetPRReviewComment :one
SELECT
  *
FROM
  pr_review_comment
WHERE
  id = ?
LIMIT
  1;

-- name: GetPRReviewComments :many
SELECT
  *
FROM
  pr_review_comment
ORDER BY
  created_at DESC;

-- name: AddPRReviewComment :one
INSERT INTO
  pr_review_comment (
    action,
    comment_id,
    pull_request_id,
    user_id,
    body,
    created_at,
    updated_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdatePRReviewComment :one
UPDATE pr_review_comment
SET
  action = ?,
  comment_id = ?,
  pull_request_id = ?,
  user_id = ?,
  body = ?,
  created_at = ?,
  updated_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeletePRReviewComment :exec
DELETE FROM pr_review_comment
WHERE
  id = ?;

-- name: GetGollumPage :one
SELECT
  *
FROM
  gollum
WHERE
  id = ?
LIMIT
  1;

-- name: GetGollumPages :many
SELECT
  *
FROM
  gollum
ORDER BY
  page_name ASC;

-- name: AddGollumPage :one
INSERT INTO
  gollum (page_name, title, summary, action, sha, html_url)
VALUES
  (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateGollumPage :one
UPDATE gollum
SET
  page_name = ?,
  title = ?,
  summary = ?,
  action = ?,
  sha = ?,
  html_url = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteGollumPage :exec
DELETE FROM gollum
WHERE
  id = ?;

-- name: GetMember :one
SELECT
  *
FROM
  member
WHERE
  id = ?
LIMIT
  1;

-- name: GetMembers :many
SELECT
  *
FROM
  member
ORDER BY
  action ASC;

-- name: AddMember :one
INSERT INTO
  member (action, member_id)
VALUES
  (?, ?) RETURNING *;

-- name: UpdateMember :one
UPDATE member
SET
  action = ?,
  member_id = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteMember :exec
DELETE FROM member
WHERE
  id = ?;

-- name: GetTeamAdd :one
SELECT
  *
FROM
  team_add
WHERE
  id = ?
LIMIT
  1;

-- name: GetTeamAdds :many
SELECT
  *
FROM
  team_add
ORDER BY
  team_name ASC;

-- name: AddTeamAdd :one
INSERT INTO
  team_add (team_id, team_name, team_slug)
VALUES
  (?, ?, ?) RETURNING *;

-- name: UpdateTeamAdd :one
UPDATE team_add
SET
  team_id = ?,
  team_name = ?,
  team_slug = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteTeamAdd :exec
DELETE FROM team_add
WHERE
  id = ?;

-- name: GetCommit :one
SELECT
  *
FROM
  commit
WHERE
  id = ?
LIMIT
  1;

-- name: GetCommits :many
SELECT
  *
FROM
  commit
ORDER BY
  id ASC;

-- name: AddCommit :one
INSERT INTO
  commit (
    payload_id,
    sha,
    author_email,
    author_name,
    message,
    url
  )
VALUES
  (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateCommit :one
UPDATE commit
SET
  payload_id = ?,
  sha = ?,
  author_email = ?,
  author_name = ?,
  message = ?,
  url = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteCommit :exec
DELETE FROM commit
WHERE
  id = ?;

-- name: GetEvent :one
SELECT
  *
FROM
  event
WHERE
  id = ?
LIMIT
  1;

-- name: GetEvents :many
SELECT
  *
FROM
  event
WHERE 
  actor_id = ?
ORDER BY
  created_at DESC;

-- name: AddEvent :one
INSERT INTO
  event (
    event_id,
    type,
    actor_id,
    repo_id,
    payload_id,
    public,
    created_at
  )
VALUES
  (?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateEvent :one
UPDATE event
SET
  event_id = ?,
  type = ?,
  actor_id = ?,
  repo_id = ?,
  payload_id = ?,
  public = ?,
  created_at = ?
WHERE
  id = ? RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM event
WHERE
  id = ?;
