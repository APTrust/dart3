package application_test

import (
	"dart/application"
	"dart/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoteRepositoryController(t *testing.T) {
	defer common.ClearDartTable()
	app := application.GetAppInstance()

	testRemoteRepositoryCreate(t, app)
	testRemoteRepositorySave(t, app)
	testRemoteRepositorySaveInvalid(t, app)
	testRemoteRepositoryEdit(t, app)
	testRemoteRepositoryDelete(t, app)
	testRemoteRepositoryList(t, app)
}

func testRemoteRepositoryCreate(t *testing.T, app *application.App) {
	resp := app.RemoteRepositoryCreate()
	assert.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, `<form method="post" action="#" id="RemoteRepository">`)
	assert.Contains(t, resp.Content, `<input type="hidden" name="ID"`)
}

func testRemoteRepositorySave(t *testing.T, app *application.App) {
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("Repo_%d", i)
		_url := fmt.Sprintf("https://repo.example.com/%d", i)
		repo := common.NewRemoteRepository()
		repo.Name = name
		repo.Url = _url
		resp := app.RemoteRepositorySave(repo)
		assert.Nil(t, resp.Data["error"])

		savedrepo, err := common.RemoteRepositoryFind(repo.ID)
		assert.Nil(t, err)
		assert.NotNil(t, savedrepo)
	}
}

func testRemoteRepositorySaveInvalid(t *testing.T, app *application.App) {
	repo := common.NewRemoteRepository()
	resp := app.RemoteRepositorySave(repo)
	assert.Equal(t, "object contains validation errors", resp.Data["error"])
	assert.Contains(t, resp.Content, "Please enter a name")
	assert.Contains(t, resp.Content, "must be a valid URL")
}

func testRemoteRepositoryEdit(t *testing.T, app *application.App) {
	repo := getFirstRemoteRepository(t, app)
	resp := app.RemoteRepositoryEdit(repo.ID)
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, `<form method="post" action="#" id="RemoteRepository">`)
	assert.Contains(t, resp.Content, fmt.Sprintf(`<input type="hidden" name="ID" value="%s"`, repo.ID))
	assert.Contains(t, resp.Content, repo.Name)
	assert.Contains(t, resp.Content, repo.Url)
}

func testRemoteRepositoryDelete(t *testing.T, app *application.App) {
	repo := getFirstRemoteRepository(t, app)
	resp := app.RemoteRepositoryDelete(repo.ID)
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)

	deletedRepo, err := common.RemoteRepositoryFind(repo.ID)
	assert.NotNil(t, err)
	assert.Nil(t, deletedRepo)
}

func testRemoteRepositoryList(t *testing.T, app *application.App) {
	// We added five repos above, then deleted one,
	// so we should have four left.
	resp := app.RemoteRepositoryList()
	require.Empty(t, resp.Data["error"])
	require.NotEmpty(t, resp.Content)
	assert.Contains(t, resp.Content, "Repo_1")
	assert.Contains(t, resp.Content, "Repo_2")
	assert.Contains(t, resp.Content, "Repo_3")
	assert.Contains(t, resp.Content, "Repo_4")
}

func getFirstRemoteRepository(t *testing.T, app *application.App) *common.RemoteRepository {
	repos, err := common.RemoteRepositoryList("obj_name", 1, 0)
	require.Nil(t, err)
	require.NotEmpty(t, repos)
	repo := repos[0]
	require.NotNil(t, repo)
	return repo
}
