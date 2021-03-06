package vcs_event

import (
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/vcs-webhook-svc/webhook"
	"strings"
	"time"
)

const (
	HeadRefPrefix = "refs/heads/"
	TagRefPrefix  = "refs/tags/"
)

func ConvertToProjectModel(from *vcs.VcsPushEvent) *model.Project {
	var (
		out            = model.Project{}
		fromOrg        = from.GetOrganization()
		fromRepository = from.GetRepository()
		fromCommits    = from.GetCommits()
		branch         = &model.Branch{
			 Meta: &model.BranchMeta{},
			Name: RemoveRefPrefix(from.GetRef()),
		}
	)

	out.Name = fromOrg.GetName()

	out.Type = from.GetVcsType().String()

	repo := ConvertEventRepositoryToRepositoryModel(fromRepository)
	repo.RepositoryType = from.GetVcsType().String()
	repo.RepositoryOwner = fromOrg.GetName()
	var branchCommits []*model.Commit

	for _, commit := range fromCommits {
		outCommit := ConvertEventCommitToCommitModel(commit)
		branchCommits = append(branchCommits, outCommit)
	}

	branch.RepositoryOwner = repo.RepositoryOwner

	branch.RepositoryName = repo.RepositoryName

	branch.RepositoryType = repo.RepositoryType
	
	branch.Commits = branchCommits

	repo.Branches = append(repo.Branches, branch)

	out.Repositories = append(out.Repositories, repo)

	return &out

}

func ConvertEventRepositoryToRepositoryModel(from *vcs.Repository) *model.Repository {
	return &model.Repository{
		RepositoryName: from.GetName(),
	}
}

func ConvertEventCommitToCommitModel(from *vcs.PushCommit) *model.Commit {

	timestamp, err := time.Parse(time.RFC3339, from.GetTimestamp())

	var timestampPtr *time.Time

	if err != nil {
		timestampPtr = nil //default
	} else {
		timestamp = timestamp.UTC()
		timestampPtr = &timestamp
	}

	out := model.Commit{
		Meta: &model.CommitMeta{},
		Sha:       from.GetSha(),
		Message:   from.GetMessage(),
		Timestamp: timestampPtr,
		Added:     from.GetAdded(),
		Modified:  from.GetModified(),
		Removed:   from.GetRemoved(),
	}

	if from.GetAuthor() != nil {
		out.AuthorName = from.GetAuthor().GetName()
		out.AuthorEmail = from.GetAuthor().GetEmail()
	}

	return &out
}

func RemoveRefPrefix(withRefPrefix string) string {
	if strings.HasPrefix(withRefPrefix, HeadRefPrefix) {
		return strings.TrimPrefix(withRefPrefix, HeadRefPrefix)
	} else if strings.HasPrefix(withRefPrefix, TagRefPrefix) {
		return strings.TrimPrefix(withRefPrefix, TagRefPrefix)
	}
	return withRefPrefix
}
