package resolvers

import (
	"context"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/cache"
	"github.com/MichaelMure/git-bug/entity"
	"github.com/MichaelMure/git-bug/graphql/connections"
	"github.com/MichaelMure/git-bug/graphql/graph"
	"github.com/MichaelMure/git-bug/graphql/models"
)

var _ graph.RepositoryResolver = &repoResolver{}

type repoResolver struct{}

func (repoResolver) Name(_ context.Context, obj *models.Repository) (*string, error) {
	name := obj.Repo.Name()
	return &name, nil
}

func (repoResolver) AllBugs(_ context.Context, obj *models.Repository, after *string, before *string, first *int, last *int, queryStr *string) (*models.BugConnection, error) {
	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	var query *cache.Query
	if queryStr != nil {
		query2, err := cache.ParseQuery(*queryStr)
		if err != nil {
			return nil, err
		}
		query = query2
	} else {
		query = cache.NewQuery()
	}

	// Simply pass a []string with the ids to the pagination algorithm
	source := obj.Repo.QueryBugs(query)

	// The edger create a custom edge holding just the id
	edger := func(id entity.Id, offset int) connections.Edge {
		return connections.LazyBugEdge{
			Id:     id,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	// The conMaker will finally load and compile bugs from git to replace the selected edges
	conMaker := func(lazyBugEdges []*connections.LazyBugEdge, lazyNode []entity.Id, info *models.PageInfo, totalCount int) (*models.BugConnection, error) {
		edges := make([]*models.BugEdge, len(lazyBugEdges))
		nodes := make([]models.BugWrapper, len(lazyBugEdges))

		for i, lazyBugEdge := range lazyBugEdges {
			excerpt, err := obj.Repo.ResolveBugExcerpt(lazyBugEdge.Id)
			if err != nil {
				return nil, err
			}

			b := models.NewLazyBug(obj.Repo, excerpt)

			edges[i] = &models.BugEdge{
				Cursor: lazyBugEdge.Cursor,
				Node:   b,
			}
			nodes[i] = b
		}

		return &models.BugConnection{
			Edges:      edges,
			Nodes:      nodes,
			PageInfo:   info,
			TotalCount: totalCount,
		}, nil
	}

	return connections.LazyBugCon(source, edger, conMaker, input)
}

func (repoResolver) Bug(_ context.Context, obj *models.Repository, prefix string) (models.BugWrapper, error) {
	excerpt, err := obj.Repo.ResolveBugExcerptPrefix(prefix)
	if err != nil {
		return nil, err
	}

	return models.NewLazyBug(obj.Repo, excerpt), nil
}

func (repoResolver) AllIdentities(_ context.Context, obj *models.Repository, after *string, before *string, first *int, last *int) (*models.IdentityConnection, error) {
	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	// Simply pass a []string with the ids to the pagination algorithm
	source := obj.Repo.AllIdentityIds()

	// The edger create a custom edge holding just the id
	edger := func(id entity.Id, offset int) connections.Edge {
		return connections.LazyIdentityEdge{
			Id:     id,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	// The conMaker will finally load and compile identities from git to replace the selected edges
	conMaker := func(lazyIdentityEdges []*connections.LazyIdentityEdge, lazyNode []entity.Id, info *models.PageInfo, totalCount int) (*models.IdentityConnection, error) {
		edges := make([]*models.IdentityEdge, len(lazyIdentityEdges))
		nodes := make([]models.IdentityWrapper, len(lazyIdentityEdges))

		for k, lazyIdentityEdge := range lazyIdentityEdges {
			excerpt, err := obj.Repo.ResolveIdentityExcerpt(lazyIdentityEdge.Id)
			if err != nil {
				return nil, err
			}

			i := models.NewLazyIdentity(obj.Repo, excerpt)

			edges[k] = &models.IdentityEdge{
				Cursor: lazyIdentityEdge.Cursor,
				Node:   i,
			}
			nodes[k] = i
		}

		return &models.IdentityConnection{
			Edges:      edges,
			Nodes:      nodes,
			PageInfo:   info,
			TotalCount: totalCount,
		}, nil
	}

	return connections.LazyIdentityCon(source, edger, conMaker, input)
}

func (repoResolver) Identity(_ context.Context, obj *models.Repository, prefix string) (models.IdentityWrapper, error) {
	excerpt, err := obj.Repo.ResolveIdentityExcerptPrefix(prefix)
	if err != nil {
		return nil, err
	}

	return models.NewLazyIdentity(obj.Repo, excerpt), nil
}

func (repoResolver) UserIdentity(_ context.Context, obj *models.Repository) (models.IdentityWrapper, error) {
	excerpt, err := obj.Repo.GetUserIdentityExcerpt()
	if err != nil {
		return nil, err
	}

	return models.NewLazyIdentity(obj.Repo, excerpt), nil
}

func (repoResolver) ValidLabels(_ context.Context, obj *models.Repository, after *string, before *string, first *int, last *int) (*models.LabelConnection, error) {
	input := models.ConnectionInput{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}

	edger := func(label bug.Label, offset int) connections.Edge {
		return models.LabelEdge{
			Node:   label,
			Cursor: connections.OffsetToCursor(offset),
		}
	}

	conMaker := func(edges []*models.LabelEdge, nodes []bug.Label, info *models.PageInfo, totalCount int) (*models.LabelConnection, error) {
		return &models.LabelConnection{
			Edges:      edges,
			Nodes:      nodes,
			PageInfo:   info,
			TotalCount: totalCount,
		}, nil
	}

	return connections.LabelCon(obj.Repo.ValidLabels(), edger, conMaker, input)
}
