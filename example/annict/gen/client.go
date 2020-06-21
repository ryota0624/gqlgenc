// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

type CreateRecordMutationPayload struct {
	CreateRecord *struct{ ClientMutationID *string }
}

type UpdateStatusMutationPayload struct {
	UpdateStatus *struct{ ClientMutationID *string }
}

type UpdateWorkStatusPayload struct {
	UpdateStatus *struct{ ClientMutationID *string }
}

type GetProfile struct {
	Viewer *struct {
		AvatarURL       *string
		RecordsCount    int64
		WannaWatchCount int64
		WatchingCount   int64
		WatchedCount    int64
	}
}

type ListWorks struct {
	Viewer *struct {
		Works *struct {
			Edges []*struct {
				Cursor string
				Node   *struct {
					Title             string
					AnnictID          int64
					SeasonYear        *int64
					SeasonName        *SeasonName
					EpisodesCount     int64
					ID                string
					OfficialSiteURL   *string
					WikipediaURL      *string
					ViewerStatusState *StatusState
				}
			}
		}
	}
}

type ListRecords struct {
	Viewer *struct {
		Records *struct {
			Edges []*struct {
				Node *struct {
					Work struct {
						Title         string
						EpisodesCount int64
					}
					Episode   struct{ SortNumber int64 }
					CreatedAt string
				}
			}
		}
	}
}

type ListNextEpisodes struct {
	Viewer *struct {
		Records *struct {
			Edges []*struct {
				Node *struct {
					Episode struct {
						NextEpisode *struct {
							ID          string
							Number      *int64
							NumberText  *string
							Title       *string
							NextEpisode *struct{ ID string }
						}
						Work struct {
							ID    string
							Title string
						}
					}
				}
			}
		}
	}
}

type GetWork struct {
	SearchWorks *struct {
		Edges []*struct {
			Node *struct {
				ID       string
				Title    string
				Episodes *struct{ Nodes []*struct{ ID string } }
			}
		}
	}
}

const CreateRecordMutationQuery = `mutation CreateRecordMutation ($episodeId: ID!) {
	createRecord(input: {episodeId:$episodeId}) {
		clientMutationId
	}
}
`

func (c *Client) CreateRecordMutation(ctx context.Context, episodeID string, httpRequestOptions ...client.HTTPRequestOption) (*CreateRecordMutationPayload, error) {
	vars := map[string]interface{}{
		"episodeID": episodeID,
	}

	var res CreateRecordMutationPayload
	if err := c.Client.Post(ctx, CreateRecordMutationQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const UpdateStatusMutationQuery = `mutation UpdateStatusMutation ($state: StatusState!, $workId: ID!) {
	updateStatus(input: {state:$state,workId:$workId}) {
		clientMutationId
	}
}
`

func (c *Client) UpdateStatusMutation(ctx context.Context, state StatusState, workID string, httpRequestOptions ...client.HTTPRequestOption) (*UpdateStatusMutationPayload, error) {
	vars := map[string]interface{}{
		"state":  state,
		"workID": workID,
	}

	var res UpdateStatusMutationPayload
	if err := c.Client.Post(ctx, UpdateStatusMutationQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const UpdateWorkStatusQuery = `mutation UpdateWorkStatus ($workId: ID!) {
	updateStatus(input: {state:WATCHING,workId:$workId}) {
		clientMutationId
	}
}
`

func (c *Client) UpdateWorkStatus(ctx context.Context, workID string, httpRequestOptions ...client.HTTPRequestOption) (*UpdateWorkStatusPayload, error) {
	vars := map[string]interface{}{
		"workID": workID,
	}

	var res UpdateWorkStatusPayload
	if err := c.Client.Post(ctx, UpdateWorkStatusQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetProfileQuery = `query GetProfile {
	viewer {
		avatarUrl
		recordsCount
		wannaWatchCount
		watchingCount
		watchedCount
	}
}
`

func (c *Client) GetProfile(ctx context.Context, httpRequestOptions ...client.HTTPRequestOption) (*GetProfile, error) {
	vars := map[string]interface{}{}

	var res GetProfile
	if err := c.Client.Post(ctx, GetProfileQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListWorksQuery = `query ListWorks ($state: StatusState, $after: String, $n: Int!) {
	viewer {
		works(state: $state, after: $after, first: $n, orderBy: {direction:DESC,field:SEASON}) {
			edges {
				cursor
				node {
					title
					annictId
					seasonYear
					seasonName
					episodesCount
					id
					officialSiteUrl
					wikipediaUrl
					viewerStatusState
				}
			}
		}
	}
}
`

func (c *Client) ListWorks(ctx context.Context, state StatusState, after string, n int64, httpRequestOptions ...client.HTTPRequestOption) (*ListWorks, error) {
	vars := map[string]interface{}{
		"state": state,
		"after": after,
		"n":     n,
	}

	var res ListWorks
	if err := c.Client.Post(ctx, ListWorksQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListRecordsQuery = `query listRecords {
	viewer {
		records {
			edges {
				node {
					work {
						title
						episodesCount
					}
					episode {
						sortNumber
					}
					createdAt
				}
			}
		}
	}
}
`

func (c *Client) ListRecords(ctx context.Context, httpRequestOptions ...client.HTTPRequestOption) (*ListRecords, error) {
	vars := map[string]interface{}{}

	var res ListRecords
	if err := c.Client.Post(ctx, ListRecordsQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListNextEpisodesQuery = `query ListNextEpisodes {
	viewer {
		records {
			edges {
				node {
					episode {
						nextEpisode {
							id
							number
							numberText
							title
							nextEpisode {
								id
							}
						}
						work {
							id
							title
						}
					}
				}
			}
		}
	}
}
`

func (c *Client) ListNextEpisodes(ctx context.Context, httpRequestOptions ...client.HTTPRequestOption) (*ListNextEpisodes, error) {
	vars := map[string]interface{}{}

	var res ListNextEpisodes
	if err := c.Client.Post(ctx, ListNextEpisodesQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const GetWorkQuery = `query GetWork ($ids: [Int!]) {
	searchWorks(annictIds: $ids) {
		edges {
			node {
				id
				title
				episodes(first: 1, orderBy: {direction:ASC,field:SORT_NUMBER}) {
					nodes {
						id
					}
				}
			}
		}
	}
}
`

func (c *Client) GetWork(ctx context.Context, ids int64, httpRequestOptions ...client.HTTPRequestOption) (*GetWork, error) {
	vars := map[string]interface{}{
		"ids": ids,
	}

	var res GetWork
	if err := c.Client.Post(ctx, GetWorkQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
