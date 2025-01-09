package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type UserPublicEvents []Event

func UnmarshalUserActivityEvents(data []byte) (UserPublicEvents, error) {
	var r UserPublicEvents
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserActivityEvents) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Event
type Event struct {
	// Actor
	Actor     ActorClass `json:"actor"`
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	// Actor
	Org     *OrgClass `json:"org,omitempty"`
	Payload Payload   `json:"payload"`
	Public  bool      `json:"public"`
	Repo    Repo      `json:"repo"`
	Type    *string   `json:"type"`
}

// Actor
type ActorClass struct {
	AvatarURL    string  `json:"avatar_url"`
	DisplayLogin *string `json:"display_login,omitempty"`
	GravatarID   *string `json:"gravatar_id"`
	ID           int64   `json:"id"`
	Login        string  `json:"login"`
	URL          string  `json:"url"`
}

// Actor
type OrgClass struct {
	AvatarURL    string  `json:"avatar_url"`
	DisplayLogin *string `json:"display_login,omitempty"`
	GravatarID   *string `json:"gravatar_id"`
	ID           int64   `json:"id"`
	Login        string  `json:"login"`
	URL          string  `json:"url"`
}

type Payload struct {
	Action *string `json:"action,omitempty"`
	// Comments provide a way for people to collaborate on an issue.
	Comment *IssueComment `json:"comment,omitempty"`
	// Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Issue *Issue `json:"issue,omitempty"`
	Pages []Page `json:"pages,omitempty"`
}

// Comments provide a way for people to collaborate on an issue.
type IssueComment struct {
	// How the author is associated with the repository.
	AuthorAssociation AuthorAssociation `json:"author_association"`
	// Contents of the issue comment
	Body      *string   `json:"body,omitempty"`
	BodyHTML  *string   `json:"body_html,omitempty"`
	BodyText  *string   `json:"body_text,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	HTMLURL   string    `json:"html_url"`
	// Unique identifier of the issue comment
	ID                    int64                         `json:"id"`
	IssueURL              string                        `json:"issue_url"`
	NodeID                string                        `json:"node_id"`
	PerformedViaGithubApp *CommentPerformedViaGithubApp `json:"performed_via_github_app,omitempty"`
	Reactions             *CommentReactions             `json:"reactions,omitempty"`
	UpdatedAt             time.Time                     `json:"updated_at"`
	// URL for the issue comment
	URL  string             `json:"url"`
	User *CommentSimpleUser `json:"user"`
}

type CommentPerformedViaGithubApp struct {
	ClientID     *string   `json:"client_id,omitempty"`
	ClientSecret *string   `json:"client_secret,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	Description  *string   `json:"description"`
	// The list of events for the GitHub app
	Events      []string `json:"events"`
	ExternalURL string   `json:"external_url"`
	HTMLURL     string   `json:"html_url"`
	// Unique identifier of the GitHub app
	ID int64 `json:"id"`
	// The number of installations associated with the GitHub app
	InstallationsCount *int64 `json:"installations_count,omitempty"`
	// The name of the GitHub app
	Name   string            `json:"name"`
	NodeID string            `json:"node_id"`
	Owner  *PurpleSimpleUser `json:"owner"`
	Pem    *string           `json:"pem,omitempty"`
	// The set of permissions for the GitHub app
	Permissions map[string]string `json:"permissions"`
	// The slug name of the GitHub app
	Slug          *string   `json:"slug,omitempty"`
	UpdatedAt     time.Time `json:"updated_at"`
	WebhookSecret *string   `json:"webhook_secret,omitempty"`
}

// A GitHub user.
type PurpleSimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type CommentReactions struct {
	The1            int64  `json:"+1"`
	ReactionRollup1 int64  `json:"-1"`
	Confused        int64  `json:"confused"`
	Eyes            int64  `json:"eyes"`
	Heart           int64  `json:"heart"`
	Hooray          int64  `json:"hooray"`
	Laugh           int64  `json:"laugh"`
	Rocket          int64  `json:"rocket"`
	TotalCount      int64  `json:"total_count"`
	URL             string `json:"url"`
}

// A GitHub user.
type CommentSimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

// Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
type Issue struct {
	ActiveLockReason *string           `json:"active_lock_reason,omitempty"`
	Assignee         *FluffySimpleUser `json:"assignee"`
	Assignees        []AssigneeElement `json:"assignees,omitempty"`
	// How the author is associated with the repository.
	AuthorAssociation AuthorAssociation `json:"author_association"`
	// Contents of the issue
	Body        *string              `json:"body,omitempty"`
	BodyHTML    *string              `json:"body_html,omitempty"`
	BodyText    *string              `json:"body_text,omitempty"`
	ClosedAt    *time.Time           `json:"closed_at"`
	ClosedBy    *TentacledSimpleUser `json:"closed_by,omitempty"`
	Comments    int64                `json:"comments"`
	CommentsURL string               `json:"comments_url"`
	CreatedAt   time.Time            `json:"created_at"`
	Draft       *bool                `json:"draft,omitempty"`
	EventsURL   string               `json:"events_url"`
	HTMLURL     string               `json:"html_url"`
	ID          int64                `json:"id"`
	// Labels to associate with this issue; pass one or more label names to replace the set of
	// labels on this issue; send an empty array to clear all labels from the issue; note that
	// the labels are silently dropped for users without push access to the repository
	Labels    []LabelElement `json:"labels"`
	LabelsURL string         `json:"labels_url"`
	Locked    bool           `json:"locked"`
	Milestone *Milestone     `json:"milestone"`
	NodeID    string         `json:"node_id"`
	// Number uniquely identifying the issue within its repository
	Number                int64                       `json:"number"`
	PerformedViaGithubApp *IssuePerformedViaGithubApp `json:"performed_via_github_app,omitempty"`
	PullRequest           *PullRequest                `json:"pull_request,omitempty"`
	Reactions             *IssueReactions             `json:"reactions,omitempty"`
	// A repository on GitHub.
	Repository    *Repository `json:"repository,omitempty"`
	RepositoryURL string      `json:"repository_url"`
	// State of the issue; either 'open' or 'closed'
	State string `json:"state"`
	// The reason for the current state
	StateReason      *StateReason      `json:"state_reason,omitempty"`
	SubIssuesSummary *SubIssuesSummary `json:"sub_issues_summary,omitempty"`
	TimelineURL      *string           `json:"timeline_url,omitempty"`
	// Title of the issue
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	// URL for the issue
	URL  string            `json:"url"`
	User *IndigoSimpleUser `json:"user"`
}

// A GitHub user.
type FluffySimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

// A GitHub user.
type AssigneeElement struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

// A GitHub user.
type TentacledSimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type LabelClass struct {
	Color       *string `json:"color,omitempty"`
	Default     *bool   `json:"default,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// A collection of related issues and pull requests.
type Milestone struct {
	ClosedAt     *time.Time           `json:"closed_at"`
	ClosedIssues int64                `json:"closed_issues"`
	CreatedAt    time.Time            `json:"created_at"`
	Creator      *MilestoneSimpleUser `json:"creator"`
	Description  *string              `json:"description"`
	DueOn        *time.Time           `json:"due_on"`
	HTMLURL      string               `json:"html_url"`
	ID           int64                `json:"id"`
	LabelsURL    string               `json:"labels_url"`
	NodeID       string               `json:"node_id"`
	// The number of the milestone.
	Number     int64 `json:"number"`
	OpenIssues int64 `json:"open_issues"`
	// The state of the milestone.
	State State `json:"state"`
	// The title of the milestone.
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
}

// A GitHub user.
type MilestoneSimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type IssuePerformedViaGithubApp struct {
	ClientID     *string   `json:"client_id,omitempty"`
	ClientSecret *string   `json:"client_secret,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	Description  *string   `json:"description"`
	// The list of events for the GitHub app
	Events      []string `json:"events"`
	ExternalURL string   `json:"external_url"`
	HTMLURL     string   `json:"html_url"`
	// Unique identifier of the GitHub app
	ID int64 `json:"id"`
	// The number of installations associated with the GitHub app
	InstallationsCount *int64 `json:"installations_count,omitempty"`
	// The name of the GitHub app
	Name   string            `json:"name"`
	NodeID string            `json:"node_id"`
	Owner  *StickySimpleUser `json:"owner"`
	Pem    *string           `json:"pem,omitempty"`
	// The set of permissions for the GitHub app
	Permissions map[string]string `json:"permissions"`
	// The slug name of the GitHub app
	Slug          *string   `json:"slug,omitempty"`
	UpdatedAt     time.Time `json:"updated_at"`
	WebhookSecret *string   `json:"webhook_secret,omitempty"`
}

// A GitHub user.
type StickySimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type PullRequest struct {
	DiffURL  *string    `json:"diff_url"`
	HTMLURL  *string    `json:"html_url"`
	MergedAt *time.Time `json:"merged_at,omitempty"`
	PatchURL *string    `json:"patch_url"`
	URL      *string    `json:"url"`
}

type IssueReactions struct {
	The1            int64  `json:"+1"`
	ReactionRollup1 int64  `json:"-1"`
	Confused        int64  `json:"confused"`
	Eyes            int64  `json:"eyes"`
	Heart           int64  `json:"heart"`
	Hooray          int64  `json:"hooray"`
	Laugh           int64  `json:"laugh"`
	Rocket          int64  `json:"rocket"`
	TotalCount      int64  `json:"total_count"`
	URL             string `json:"url"`
}

// A repository on GitHub.
type Repository struct {
	// Whether to allow Auto-merge to be used on pull requests.
	AllowAutoMerge *bool `json:"allow_auto_merge,omitempty"`
	// Whether to allow forking this repo
	AllowForking *bool `json:"allow_forking,omitempty"`
	// Whether to allow merge commits for pull requests.
	AllowMergeCommit *bool `json:"allow_merge_commit,omitempty"`
	// Whether to allow rebase merges for pull requests.
	AllowRebaseMerge *bool `json:"allow_rebase_merge,omitempty"`
	// Whether to allow squash merges for pull requests.
	AllowSquashMerge *bool `json:"allow_squash_merge,omitempty"`
	// Whether or not a pull request head branch that is behind its base branch can always be
	// updated even if it is not required to be up to date before merging.
	AllowUpdateBranch *bool `json:"allow_update_branch,omitempty"`
	// Whether anonymous git access is enabled for this repository
	AnonymousAccessEnabled *bool  `json:"anonymous_access_enabled,omitempty"`
	ArchiveURL             string `json:"archive_url"`
	// Whether the repository is archived.
	Archived         bool       `json:"archived"`
	AssigneesURL     string     `json:"assignees_url"`
	BlobsURL         string     `json:"blobs_url"`
	BranchesURL      string     `json:"branches_url"`
	CloneURL         string     `json:"clone_url"`
	CollaboratorsURL string     `json:"collaborators_url"`
	CommentsURL      string     `json:"comments_url"`
	CommitsURL       string     `json:"commits_url"`
	CompareURL       string     `json:"compare_url"`
	ContentsURL      string     `json:"contents_url"`
	ContributorsURL  string     `json:"contributors_url"`
	CreatedAt        *time.Time `json:"created_at"`
	// The default branch of the repository.
	DefaultBranch string `json:"default_branch"`
	// Whether to delete head branches when pull requests are merged
	DeleteBranchOnMerge *bool   `json:"delete_branch_on_merge,omitempty"`
	DeploymentsURL      string  `json:"deployments_url"`
	Description         *string `json:"description"`
	// Returns whether or not this repository disabled.
	Disabled      bool   `json:"disabled"`
	DownloadsURL  string `json:"downloads_url"`
	EventsURL     string `json:"events_url"`
	Fork          bool   `json:"fork"`
	Forks         int64  `json:"forks"`
	ForksCount    int64  `json:"forks_count"`
	ForksURL      string `json:"forks_url"`
	FullName      string `json:"full_name"`
	GitCommitsURL string `json:"git_commits_url"`
	GitRefsURL    string `json:"git_refs_url"`
	GitTagsURL    string `json:"git_tags_url"`
	GitURL        string `json:"git_url"`
	// Whether discussions are enabled.
	HasDiscussions *bool `json:"has_discussions,omitempty"`
	// Whether downloads are enabled.
	HasDownloads bool `json:"has_downloads"`
	// Whether issues are enabled.
	HasIssues bool `json:"has_issues"`
	HasPages  bool `json:"has_pages"`
	// Whether projects are enabled.
	HasProjects bool `json:"has_projects"`
	// Whether the wiki is enabled.
	HasWiki  bool    `json:"has_wiki"`
	Homepage *string `json:"homepage"`
	HooksURL string  `json:"hooks_url"`
	HTMLURL  string  `json:"html_url"`
	// Unique identifier of the repository
	ID int64 `json:"id"`
	// Whether this repository acts as a template that can be used to generate new repositories.
	IsTemplate      *bool          `json:"is_template,omitempty"`
	IssueCommentURL string         `json:"issue_comment_url"`
	IssueEventsURL  string         `json:"issue_events_url"`
	IssuesURL       string         `json:"issues_url"`
	KeysURL         string         `json:"keys_url"`
	LabelsURL       string         `json:"labels_url"`
	Language        *string        `json:"language"`
	LanguagesURL    string         `json:"languages_url"`
	License         *LicenseSimple `json:"license"`
	MasterBranch    *string        `json:"master_branch,omitempty"`
	// The default value for a merge commit message.
	//
	// - `PR_TITLE` - default to the pull request's title.
	// - `PR_BODY` - default to the pull request's body.
	// - `BLANK` - default to a blank commit message.
	MergeCommitMessage *MergeCommitMessage `json:"merge_commit_message,omitempty"`
	// The default value for a merge commit title.
	//
	// - `PR_TITLE` - default to the pull request's title.
	// - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull
	// request #123 from branch-name).
	MergeCommitTitle *MergeCommitTitle `json:"merge_commit_title,omitempty"`
	MergesURL        string            `json:"merges_url"`
	MilestonesURL    string            `json:"milestones_url"`
	MirrorURL        *string           `json:"mirror_url"`
	// The name of the repository.
	Name             string `json:"name"`
	NodeID           string `json:"node_id"`
	NotificationsURL string `json:"notifications_url"`
	OpenIssues       int64  `json:"open_issues"`
	OpenIssuesCount  int64  `json:"open_issues_count"`
	// A GitHub user.
	Owner       OwnerClass   `json:"owner"`
	Permissions *Permissions `json:"permissions,omitempty"`
	// Whether the repository is private or public.
	Private     bool       `json:"private"`
	PullsURL    string     `json:"pulls_url"`
	PushedAt    *time.Time `json:"pushed_at"`
	ReleasesURL string     `json:"releases_url"`
	// The size of the repository, in kilobytes. Size is calculated hourly. When a repository is
	// initially created, the size is 0.
	Size int64 `json:"size"`
	// The default value for a squash merge commit message:
	//
	// - `PR_BODY` - default to the pull request's body.
	// - `COMMIT_MESSAGES` - default to the branch's commit messages.
	// - `BLANK` - default to a blank commit message.
	SquashMergeCommitMessage *SquashMergeCommitMessage `json:"squash_merge_commit_message,omitempty"`
	// The default value for a squash merge commit title:
	//
	// - `PR_TITLE` - default to the pull request's title.
	// - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull
	// request's title (when more than one commit).
	SquashMergeCommitTitle *SquashMergeCommitTitle `json:"squash_merge_commit_title,omitempty"`
	SSHURL                 string                  `json:"ssh_url"`
	StargazersCount        int64                   `json:"stargazers_count"`
	StargazersURL          string                  `json:"stargazers_url"`
	StarredAt              *string                 `json:"starred_at,omitempty"`
	StatusesURL            string                  `json:"statuses_url"`
	SubscribersURL         string                  `json:"subscribers_url"`
	SubscriptionURL        string                  `json:"subscription_url"`
	SvnURL                 string                  `json:"svn_url"`
	TagsURL                string                  `json:"tags_url"`
	TeamsURL               string                  `json:"teams_url"`
	TempCloneToken         *string                 `json:"temp_clone_token,omitempty"`
	Topics                 []string                `json:"topics,omitempty"`
	TreesURL               string                  `json:"trees_url"`
	UpdatedAt              *time.Time              `json:"updated_at"`
	URL                    string                  `json:"url"`
	// Whether a squash merge commit can use the pull request title as default. **This property
	// is closing down. Please use `squash_merge_commit_title` instead.
	UseSquashPRTitleAsDefault *bool `json:"use_squash_pr_title_as_default,omitempty"`
	// The repository visibility: public, private, or internal.
	Visibility    *string `json:"visibility,omitempty"`
	Watchers      int64   `json:"watchers"`
	WatchersCount int64   `json:"watchers_count"`
	// Whether to require contributors to sign off on web-based commits
	WebCommitSignoffRequired *bool `json:"web_commit_signoff_required,omitempty"`
}

// License Simple
type LicenseSimple struct {
	HTMLURL *string `json:"html_url,omitempty"`
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	NodeID  string  `json:"node_id"`
	SpdxID  *string `json:"spdx_id"`
	URL     *string `json:"url"`
}

// A GitHub user.
type OwnerClass struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type Permissions struct {
	Admin    bool  `json:"admin"`
	Maintain *bool `json:"maintain,omitempty"`
	Pull     bool  `json:"pull"`
	Push     bool  `json:"push"`
	Triage   *bool `json:"triage,omitempty"`
}

type SubIssuesSummary struct {
	Completed        int64 `json:"completed"`
	PercentCompleted int64 `json:"percent_completed"`
	Total            int64 `json:"total"`
}

// A GitHub user.
type IndigoSimpleUser struct {
	AvatarURL         string  `json:"avatar_url"`
	Email             *string `json:"email,omitempty"`
	EventsURL         string  `json:"events_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	GravatarID        *string `json:"gravatar_id"`
	HTMLURL           string  `json:"html_url"`
	ID                int64   `json:"id"`
	Login             string  `json:"login"`
	Name              *string `json:"name,omitempty"`
	NodeID            string  `json:"node_id"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	ReposURL          string  `json:"repos_url"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         *string `json:"starred_at,omitempty"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	Type              string  `json:"type"`
	URL               string  `json:"url"`
	UserViewType      *string `json:"user_view_type,omitempty"`
}

type Page struct {
	Action   *string `json:"action,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	PageName *string `json:"page_name,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	Summary  *string `json:"summary,omitempty"`
	Title    *string `json:"title,omitempty"`
}

type Repo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// How the author is associated with the repository.
type AuthorAssociation string

const (
	Collaborator         AuthorAssociation = "COLLABORATOR"
	Contributor          AuthorAssociation = "CONTRIBUTOR"
	FirstTimeContributor AuthorAssociation = "FIRST_TIME_CONTRIBUTOR"
	FirstTimer           AuthorAssociation = "FIRST_TIMER"
	Mannequin            AuthorAssociation = "MANNEQUIN"
	Member               AuthorAssociation = "MEMBER"
	None                 AuthorAssociation = "NONE"
	Owner                AuthorAssociation = "OWNER"
)

// The state of the milestone.
type State string

const (
	Closed State = "closed"
	Open   State = "open"
)

// The default value for a merge commit message.
//
// - `PR_TITLE` - default to the pull request's title.
// - `PR_BODY` - default to the pull request's body.
// - `BLANK` - default to a blank commit message.
type MergeCommitMessage string

const (
	MergeCommitMessageBLANK   MergeCommitMessage = "BLANK"
	MergeCommitMessagePRBODY  MergeCommitMessage = "PR_BODY"
	MergeCommitMessagePRTITLE MergeCommitMessage = "PR_TITLE"
)

// The default value for a merge commit title.
//
// - `PR_TITLE` - default to the pull request's title.
// - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull
// request #123 from branch-name).
type MergeCommitTitle string

const (
	MergeCommitTitlePRTITLE MergeCommitTitle = "PR_TITLE"
	MergeMessage            MergeCommitTitle = "MERGE_MESSAGE"
)

// The default value for a squash merge commit message:
//
// - `PR_BODY` - default to the pull request's body.
// - `COMMIT_MESSAGES` - default to the branch's commit messages.
// - `BLANK` - default to a blank commit message.
type SquashMergeCommitMessage string

const (
	CommitMessages                 SquashMergeCommitMessage = "COMMIT_MESSAGES"
	SquashMergeCommitMessageBLANK  SquashMergeCommitMessage = "BLANK"
	SquashMergeCommitMessagePRBODY SquashMergeCommitMessage = "PR_BODY"
)

// The default value for a squash merge commit title:
//
// - `PR_TITLE` - default to the pull request's title.
// - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull
// request's title (when more than one commit).
type SquashMergeCommitTitle string

const (
	CommitOrPRTitle               SquashMergeCommitTitle = "COMMIT_OR_PR_TITLE"
	SquashMergeCommitTitlePRTITLE SquashMergeCommitTitle = "PR_TITLE"
)

type StateReason string

const (
	Completed  StateReason = "completed"
	NotPlanned StateReason = "not_planned"
	Reopened   StateReason = "reopened"
)

type LabelElement struct {
	LabelClass *LabelClass
	String     *string
}

func (x *LabelElement) UnmarshalJSON(data []byte) error {
	x.LabelClass = nil
	var c LabelClass
	object, err := unmarshalUnion(data, nil, nil, nil, &x.String, false, nil, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.LabelClass = &c
	}
	return nil
}

func (x *LabelElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, false, nil, x.LabelClass != nil, x.LabelClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
