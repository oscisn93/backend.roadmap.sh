package client

import (
	"time"
)

type PublicUserEvents []Event

type Event struct {
	Actor     ActorClass `json:"actor"`
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Org       *OrgClass  `json:"org,omitempty"`
	Payload   Payload    `json:"payload"`
	Public    bool       `json:"public"`
	Repo      Repo       `json:"repo"`
	Type      *string    `json:"type"`
}

type ActorClass struct {
	AvatarURL    string  `json:"avatar_url"`
	DisplayLogin *string `json:"display_login,omitempty"`
	GravatarID   *string `json:"gravatar_id"`
	ID           int64   `json:"id"`
	Login        string  `json:"login"`
	URL          string  `json:"url"`
}

type OrgClass struct {
	AvatarURL    string  `json:"avatar_url"`
	DisplayLogin *string `json:"display_login,omitempty"`
	GravatarID   *string `json:"gravatar_id"`
	ID           int64   `json:"id"`
	Login        string  `json:"login"`
	URL          string  `json:"url"`
}

type Payload struct {
	Action  *string       `json:"action,omitempty"`
	Comment *IssueComment `json:"comment,omitempty"`
	Issue   *Issue        `json:"issue,omitempty"`
	Pages   []Page        `json:"pages,omitempty"`
}

type IssueComment struct {
	AuthorAssociation     AuthorAssociation             `json:"author_association"`
	Body                  *string                       `json:"body,omitempty"`
	BodyHTML              *string                       `json:"body_html,omitempty"`
	BodyText              *string                       `json:"body_text,omitempty"`
	CreatedAt             time.Time                     `json:"created_at"`
	HTMLURL               string                        `json:"html_url"`
	ID                    int64                         `json:"id"`
	IssueURL              string                        `json:"issue_url"`
	NodeID                string                        `json:"node_id"`
	PerformedViaGithubApp *CommentPerformedViaGithubApp `json:"performed_via_github_app,omitempty"`
	Reactions             *CommentReactions             `json:"reactions,omitempty"`
	UpdatedAt             time.Time                     `json:"updated_at"`
	URL                   string                        `json:"url"`
	User                  *CommentSimpleUser            `json:"user"`
}

type CommentPerformedViaGithubApp struct {
	ClientID           *string           `json:"client_id,omitempty"`
	ClientSecret       *string           `json:"client_secret,omitempty"`
	CreatedAt          time.Time         `json:"created_at"`
	Description        *string           `json:"description"`
	Events             []string          `json:"events"`
	ExternalURL        string            `json:"external_url"`
	HTMLURL            string            `json:"html_url"`
	ID                 int64             `json:"id"`
	InstallationsCount *int64            `json:"installations_count,omitempty"`
	Name               string            `json:"name"`
	NodeID             string            `json:"node_id"`
	Owner              *PurpleSimpleUser `json:"owner"`
	Pem                *string           `json:"pem,omitempty"`
	Permissions        map[string]string `json:"permissions"`
	Slug               *string           `json:"slug,omitempty"`
	UpdatedAt          time.Time         `json:"updated_at"`
	WebhookSecret      *string           `json:"webhook_secret,omitempty"`
}

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

type Issue struct {
	ActiveLockReason      *string                     `json:"active_lock_reason,omitempty"`
	Assignee              *FluffySimpleUser           `json:"assignee"`
	Assignees             []AssigneeElement           `json:"assignees,omitempty"`
	AuthorAssociation     AuthorAssociation           `json:"author_association"`
	Body                  *string                     `json:"body,omitempty"`
	BodyHTML              *string                     `json:"body_html,omitempty"`
	BodyText              *string                     `json:"body_text,omitempty"`
	ClosedAt              *time.Time                  `json:"closed_at"`
	ClosedBy              *TentacledSimpleUser        `json:"closed_by,omitempty"`
	Comments              int64                       `json:"comments"`
	CommentsURL           string                      `json:"comments_url"`
	CreatedAt             time.Time                   `json:"created_at"`
	Draft                 *bool                       `json:"draft,omitempty"`
	EventsURL             string                      `json:"events_url"`
	HTMLURL               string                      `json:"html_url"`
	ID                    int64                       `json:"id"`
	Labels                []LabelElement              `json:"labels"`
	LabelsURL             string                      `json:"labels_url"`
	Locked                bool                        `json:"locked"`
	Milestone             *Milestone                  `json:"milestone"`
	NodeID                string                      `json:"node_id"`
	Number                int64                       `json:"number"`
	PerformedViaGithubApp *IssuePerformedViaGithubApp `json:"performed_via_github_app,omitempty"`
	PullRequest           *PullRequest                `json:"pull_request,omitempty"`
	Reactions             *IssueReactions             `json:"reactions,omitempty"`
	Repository            *Repository                 `json:"repository,omitempty"`
	RepositoryURL         string                      `json:"repository_url"`
	State                 string                      `json:"state"`
	StateReason           *StateReason                `json:"state_reason,omitempty"`
	SubIssuesSummary      *SubIssuesSummary           `json:"sub_issues_summary,omitempty"`
	TimelineURL           *string                     `json:"timeline_url,omitempty"`
	Title                 string                      `json:"title"`
	UpdatedAt             time.Time                   `json:"updated_at"`
	URL                   string                      `json:"url"`
	User                  *IndigoSimpleUser           `json:"user"`
}

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
	Number       int64                `json:"number"`
	OpenIssues   int64                `json:"open_issues"`
	State        State                `json:"state"`
	Title        string               `json:"title"`
	UpdatedAt    time.Time            `json:"updated_at"`
	URL          string               `json:"url"`
}

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
	ClientID           *string           `json:"client_id,omitempty"`
	ClientSecret       *string           `json:"client_secret,omitempty"`
	CreatedAt          time.Time         `json:"created_at"`
	Description        *string           `json:"description"`
	Events             []string          `json:"events"`
	ExternalURL        string            `json:"external_url"`
	HTMLURL            string            `json:"html_url"`
	ID                 int64             `json:"id"`
	InstallationsCount *int64            `json:"installations_count,omitempty"`
	Name               string            `json:"name"`
	NodeID             string            `json:"node_id"`
	Owner              *StickySimpleUser `json:"owner"`
	Pem                *string           `json:"pem,omitempty"`
	Permissions        map[string]string `json:"permissions"`
	Slug               *string           `json:"slug,omitempty"`
	UpdatedAt          time.Time         `json:"updated_at"`
	WebhookSecret      *string           `json:"webhook_secret,omitempty"`
}

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

type Repository struct {
	AllowAutoMerge            *bool                     `json:"allow_auto_merge,omitempty"`
	AllowForking              *bool                     `json:"allow_forking,omitempty"`
	AllowMergeCommit          *bool                     `json:"allow_merge_commit,omitempty"`
	AllowRebaseMerge          *bool                     `json:"allow_rebase_merge,omitempty"`
	AllowSquashMerge          *bool                     `json:"allow_squash_merge,omitempty"`
	AllowUpdateBranch         *bool                     `json:"allow_update_branch,omitempty"`
	AnonymousAccessEnabled    *bool                     `json:"anonymous_access_enabled,omitempty"`
	ArchiveURL                string                    `json:"archive_url"`
	Archived                  bool                      `json:"archived"`
	AssigneesURL              string                    `json:"assignees_url"`
	BlobsURL                  string                    `json:"blobs_url"`
	BranchesURL               string                    `json:"branches_url"`
	CloneURL                  string                    `json:"clone_url"`
	CollaboratorsURL          string                    `json:"collaborators_url"`
	CommentsURL               string                    `json:"comments_url"`
	CommitsURL                string                    `json:"commits_url"`
	CompareURL                string                    `json:"compare_url"`
	ContentsURL               string                    `json:"contents_url"`
	ContributorsURL           string                    `json:"contributors_url"`
	CreatedAt                 *time.Time                `json:"created_at"`
	DefaultBranch             string                    `json:"default_branch"`
	DeleteBranchOnMerge       *bool                     `json:"delete_branch_on_merge,omitempty"`
	DeploymentsURL            string                    `json:"deployments_url"`
	Description               *string                   `json:"description"`
	Disabled                  bool                      `json:"disabled"`
	DownloadsURL              string                    `json:"downloads_url"`
	EventsURL                 string                    `json:"events_url"`
	Fork                      bool                      `json:"fork"`
	Forks                     int64                     `json:"forks"`
	ForksCount                int64                     `json:"forks_count"`
	ForksURL                  string                    `json:"forks_url"`
	FullName                  string                    `json:"full_name"`
	GitCommitsURL             string                    `json:"git_commits_url"`
	GitRefsURL                string                    `json:"git_refs_url"`
	GitTagsURL                string                    `json:"git_tags_url"`
	GitURL                    string                    `json:"git_url"`
	HasDiscussions            *bool                     `json:"has_discussions,omitempty"`
	HasDownloads              bool                      `json:"has_downloads"`
	HasIssues                 bool                      `json:"has_issues"`
	HasPages                  bool                      `json:"has_pages"`
	HasProjects               bool                      `json:"has_projects"`
	HasWiki                   bool                      `json:"has_wiki"`
	Homepage                  *string                   `json:"homepage"`
	HooksURL                  string                    `json:"hooks_url"`
	HTMLURL                   string                    `json:"html_url"`
	ID                        int64                     `json:"id"`
	IsTemplate                *bool                     `json:"is_template,omitempty"`
	IssueCommentURL           string                    `json:"issue_comment_url"`
	IssueEventsURL            string                    `json:"issue_events_url"`
	IssuesURL                 string                    `json:"issues_url"`
	KeysURL                   string                    `json:"keys_url"`
	LabelsURL                 string                    `json:"labels_url"`
	Language                  *string                   `json:"language"`
	LanguagesURL              string                    `json:"languages_url"`
	License                   *LicenseSimple            `json:"license"`
	MasterBranch              *string                   `json:"master_branch,omitempty"`
	MergeCommitMessage        *MergeCommitMessage       `json:"merge_commit_message,omitempty"`
	MergeCommitTitle          *MergeCommitTitle         `json:"merge_commit_title,omitempty"`
	MergesURL                 string                    `json:"merges_url"`
	MilestonesURL             string                    `json:"milestones_url"`
	MirrorURL                 *string                   `json:"mirror_url"`
	Name                      string                    `json:"name"`
	NodeID                    string                    `json:"node_id"`
	NotificationsURL          string                    `json:"notifications_url"`
	OpenIssues                int64                     `json:"open_issues"`
	OpenIssuesCount           int64                     `json:"open_issues_count"`
	Owner                     OwnerClass                `json:"owner"`
	Permissions               *Permissions              `json:"permissions,omitempty"`
	Private                   bool                      `json:"private"`
	PullsURL                  string                    `json:"pulls_url"`
	PushedAt                  *time.Time                `json:"pushed_at"`
	ReleasesURL               string                    `json:"releases_url"`
	Size                      int64                     `json:"size"`
	SquashMergeCommitMessage  *SquashMergeCommitMessage `json:"squash_merge_commit_message,omitempty"`
	SquashMergeCommitTitle    *SquashMergeCommitTitle   `json:"squash_merge_commit_title,omitempty"`
	SSHURL                    string                    `json:"ssh_url"`
	StargazersCount           int64                     `json:"stargazers_count"`
	StargazersURL             string                    `json:"stargazers_url"`
	StarredAt                 *string                   `json:"starred_at,omitempty"`
	StatusesURL               string                    `json:"statuses_url"`
	SubscribersURL            string                    `json:"subscribers_url"`
	SubscriptionURL           string                    `json:"subscription_url"`
	SvnURL                    string                    `json:"svn_url"`
	TagsURL                   string                    `json:"tags_url"`
	TeamsURL                  string                    `json:"teams_url"`
	TempCloneToken            *string                   `json:"temp_clone_token,omitempty"`
	Topics                    []string                  `json:"topics,omitempty"`
	TreesURL                  string                    `json:"trees_url"`
	UpdatedAt                 *time.Time                `json:"updated_at"`
	URL                       string                    `json:"url"`
	UseSquashPRTitleAsDefault *bool                     `json:"use_squash_pr_title_as_default,omitempty"`
	Visibility                *string                   `json:"visibility,omitempty"`
	Watchers                  int64                     `json:"watchers"`
	WatchersCount             int64                     `json:"watchers_count"`
	WebCommitSignoffRequired  *bool                     `json:"web_commit_signoff_required,omitempty"`
}

type LicenseSimple struct {
	HTMLURL *string `json:"html_url,omitempty"`
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	NodeID  string  `json:"node_id"`
	SpdxID  *string `json:"spdx_id"`
	URL     *string `json:"url"`
}

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

type State string

const (
	Closed State = "closed"
	Open   State = "open"
)

type MergeCommitMessage string

const (
	MergeCommitMessageBLANK   MergeCommitMessage = "BLANK"
	MergeCommitMessagePRBODY  MergeCommitMessage = "PR_BODY"
	MergeCommitMessagePRTITLE MergeCommitMessage = "PR_TITLE"
)

type MergeCommitTitle string

const (
	MergeCommitTitlePRTITLE MergeCommitTitle = "PR_TITLE"
	MergeMessage            MergeCommitTitle = "MERGE_MESSAGE"
)

type SquashMergeCommitMessage string

const (
	CommitMessages                 SquashMergeCommitMessage = "COMMIT_MESSAGES"
	SquashMergeCommitMessageBLANK  SquashMergeCommitMessage = "BLANK"
	SquashMergeCommitMessagePRBODY SquashMergeCommitMessage = "PR_BODY"
)

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
