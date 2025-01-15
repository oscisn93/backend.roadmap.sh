--- Table to store GitHub actors (users)
CREATE TABLE
  IF NOT EXISTS actor (
    id INTEGER PRIMARY KEY,
    login TEXT NOT NULL,
    display_login TEXT,
    gravatar_id TEXT,
    url TEXT,
    avatar_url TEXT
  );

-- Table to store GitHub repositories
CREATE TABLE
  IF NOT EXISTS repo (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL
  );

-- Table to store push payload
CREATE TABLE 
  IF NOT EXISTS push (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    push_id TEXT NOT NULL,
    repository_id INTEGER NOT NULL,
    size INTEGER NOT NULL,
    distinct_size INTEGER NOT NULL,
    ref TEXT NOT NULL,
    head TEXT NOT NULL,
    before TEXT NOT NULL
  );

-- Table to store pull request payload
CREATE TABLE 
  IF NOT EXISTS pull_request (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    number INTEGER NOT NULL,
    pull_request_id INTEGER NOT NULL,
    state TEXT NOT NULL,
    title TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    closed_at DATETIME,
    merged_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES actor(id)
  );

-- Table to store issue payload
CREATE TABLE
  IF NOT EXISTS issue (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    issue_id INTEGER NOT NULL,
    number INTEGER NOT NULL,
    title TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    state TEXT NOT NULL,
    body TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    closed_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES actor(id)
  );

-- Table to store fork payload
CREATE TABLE
  IF NOT EXISTS fork (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    forkee_id INTEGER NOT NULL,
    forkee_name TEXT NOT NULL,
    forkee_full_name TEXT NOT NULL,
    forkee_private BOOLEAN NOT NULL,
    FOREIGN KEY (forkee_id) REFERENCES repo(id)
  );

-- Table to store watch payload
CREATE TABLE
  IF NOT EXISTS watch (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL
  );

-- Table to store create payload
CREATE TABLE
  IF NOT EXISTS ref_create (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ref TEXT,
    ref_type TEXT NOT NULL,
    master_branch TEXT,
    description TEXT,
    pusher_type TEXT NOT NULL
  );

-- Table to store delete payload
CREATE TABLE
  IF NOT EXISTS ref_delete (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ref TEXT,
    ref_type TEXT NOT NULL
  );

-- Table to store release payload
CREATE TABLE
  IF NOT EXISTS release (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    release_id INTEGER NOT NULL,
    tag_name TEXT NOT NULL,
    target_commitish TEXT NOT NULL,
    name TEXT,
    body TEXT,
    draft BOOLEAN NOT NULL,
    prerelease BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL,
    published_at DATETIME NOT NULL
  );

-- Table to store issue comment payload
CREATE TABLE
  IF NOT EXISTS issue_comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    issue_id INTEGER NOT NULL,
    comment_id INTEGER NOT NULL,
    comment_user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (issue_id) REFERENCES issue(id),
    FOREIGN KEY (comment_user_id) REFERENCES actor(id)
  );

-- Table to store pull request review payload
CREATE TABLE
  IF NOT EXISTS pull_request_review (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    review_id INTEGER NOT NULL,
    pull_request_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT,
    submitted_at DATETIME NOT NULL,
    FOREIGN KEY (pull_request_id) REFERENCES pull_request(id),
    FOREIGN KEY (user_id) REFERENCES actor(id)
  );

-- Table to store pull request review comment payload
CREATE TABLE
  IF NOT EXISTS pr_review_comment (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    comment_id INTEGER NOT NULL,
    pull_request_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (pull_request_id) REFERENCES pull_request(id),
    FOREIGN KEY (user_id) REFERENCES actor(id)
  );

-- Table to store gollum payload
CREATE TABLE
  IF NOT EXISTS gollum (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    page_name TEXT NOT NULL,
    title TEXT NOT NULL,
    summary TEXT,
    action TEXT NOT NULL,
    sha TEXT NOT NULL,
    html_url TEXT NOT NULL
  );

-- Table to store member payload
CREATE TABLE
  IF NOT EXISTS member (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    action TEXT NOT NULL,
    member_id INTEGER NOT NULL,
    FOREIGN KEY (member_id) REFERENCES actor(id)
  );

-- Table to store team add payload
CREATE TABLE
  IF NOT EXISTS team_add (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_id INTEGER NOT NULL,
    team_name TEXT NOT NULL,
    team_slug TEXT NOT NULL,
    FOREIGN KEY (team_id) REFERENCES actor(id)
  );

-- Table to store commit details
CREATE TABLE
  IF NOT EXISTS commit (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    payload_id INTEGER NOT NULL,
    sha TEXT NOT NULL,
    author_email TEXT NOT NULL,
    author_name TEXT NOT NULL,
    message TEXT NOT NULL,
    url TEXT NOT NULL,
    FOREIGN KEY (payload_id) REFERENCES push_event(id)
  );

-- Table to store cached GitHub user events
CREATE TABLE
  IF NOT EXISTS event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL,
    actor_id INTEGER NOT NULL,
    repo_id INTEGER NOT NULL,
    payload_id INTEGER NOT NULL,
    public BOOLEAN NOT NULL,
    created_at DATETIME NOT NULL,
    cached_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (actor_id) REFERENCES actor(id),
    FOREIGN KEY (repo_id) REFERENCES repo(id)
  );

-- Indexes
CREATE INDEX 
  IF NOT EXISTS idx_actor_id_event_type 
    ON event(actor_id, type);
CREATE INDEX 
  IF NOT EXISTS idx_event_created_at 
    ON event(created_at);
CREATE INDEX 
  IF NOT EXISTS idx_actor_login 
    ON actor(login);
CREATE INDEX 
  IF NOT EXISTS idx_repo_name 
    ON repo(name);
CREATE INDEX 
  IF NOT EXISTS idx_commit_payload_id
    ON commit(payload_id);