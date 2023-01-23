package common

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// RemoteRepository contains config settings describing how to
// connect to a remote repository, such as APTrust. Presumably,
// this is a repository into which you are ingesting data,
// and the repository has a REST API.
//
// The repo config allows you to connect to the repo so you can
// see the state of bags you uploaded. The logic for performing
// those requests and parsing the responses has to be implemented
// elsewhere. In DART 2.x, this was done with plugins, and APTrust
// was the only existing plugin. In DART 3.x, the way to add new
// repo implementations is to be determined. One suggestion is to
// generate clients with Swagger/OpenAPI.
type RemoteRepository struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Url        string            `json:"url"`
	UserID     string            `json:"userId"`
	APIToken   string            `json:"apiToken"`
	LoginExtra string            `json:"loginExtra"`
	PluginID   string            `json:"pluginId"`
	Errors     map[string]string `json:"errors"`
}

func NewRemoteRepository() *RemoteRepository {
	return &RemoteRepository{
		ID: uuid.NewString(),
	}
}

// RemoteRepositoryFind returns the RemoteRepository with the specified UUID,
// or sql.ErrNoRows if no matching record exists.
func RemoteRepositoryFind(uuid string) (*RemoteRepository, error) {
	result, err := ObjFind(uuid)
	if err != nil {
		return nil, err
	}
	return result.RemoteRepository, err
}

// RemoteRepositoryList returns a list of RemoteRepositorys with the specified
// order, offset and limit.
func RemoteRepositoryList(orderBy string, limit, offset int) ([]*RemoteRepository, error) {
	result, err := ObjList(TypeRemoteRepository, orderBy, limit, offset)
	if err != nil {
		return nil, err
	}
	return result.RemoteRepositories, err
}

// ObjID returns this remote repo's UUID.
func (repo *RemoteRepository) ObjID() string {
	return repo.ID
}

// ObjName returns the name of this remote repo.
func (repo *RemoteRepository) ObjName() string {
	return repo.Name
}

// ObjType returns this object's type.
func (repo *RemoteRepository) ObjType() string {
	return TypeRemoteRepository
}

func (repo *RemoteRepository) String() string {
	return fmt.Sprintf("RemoteRepository: '%s'", repo.Name)
}

// Save saves this repo, if it determines the repo is valid.
// It returns common.ErrObjecValidation if the repo is invalid.
// Check repo.Errors if you get a validation error.
func (repo *RemoteRepository) Save() error {
	if !repo.Validate() {
		return ErrObjecValidation
	}
	return ObjSave(repo)
}

// Delete deletes this repo config.
func (repo *RemoteRepository) Delete() error {
	return ObjDelete(repo.ID)
}

// Validate returns true if this RemoteRepository config contains
// valid settings, false if not. Check the Errors map if this returns
// false.
func (repo *RemoteRepository) Validate() bool {
	repo.Errors = make(map[string]string)
	if !LooksLikeHypertextURL(repo.Url) {
		repo.Errors["Url"] = "Repository URL must a valid URL beginning with http:// or https://."
		return false
	}
	return true
}

func (repo *RemoteRepository) ToForm() *Form {
	form := NewForm(TypeRemoteRepository, repo.ID)

	form.AddField("ID", "ID", repo.ID, true)

	nameField := form.AddField("Name", "Name", repo.Name, true)
	nameField.Error = repo.Errors["Name"]

	urlField := form.AddField("Url", "URL", repo.Url, true)
	urlField.Error = repo.Errors["Url"]

	// TODO: Improve creation of choices for select list.
	// TODO: Restore auto-generated ids, as in DART 2
	pluginIdField := form.AddField("PluginID", "Plugin ID", repo.PluginID, true)
	pluginIdField.Error = repo.Errors["PluginID"]
	pluginIdField.Choices = []Choice{
		{"", "", repo.PluginID == ""},
		{"APTrustClient", PluginAPTrustClient, repo.PluginID == PluginAPTrustClient},
	}
	pluginIdField.ID = "RemoteRepostoryForm_PluginID"

	userIdField := form.AddField("UserID", "User", repo.UserID, true)
	userIdField.Error = repo.Errors["UserID"]

	apiTokenField := form.AddField("APIToken", "API Token", repo.APIToken, true)
	apiTokenField.Error = repo.Errors["APIToken"]

	loginExtraField := form.AddField("LoginExtra", "Login Extra", repo.LoginExtra, true)
	loginExtraField.Error = repo.Errors["LoginExtra"]

	form.CancelFunction = "RemoteRepositoryList"
	form.SubmitFunction = "RemoteRepositorySave"
	form.DeleteFunction = "RemoteRepositoryDelete"
	return form
}

func remoteRepositoryList(rows *sql.Rows) ([]*RemoteRepository, error) {
	list := make([]*RemoteRepository, 0)
	for rows.Next() {
		var jsonStr string
		err := rows.Scan(&jsonStr)
		if err != nil {
			return nil, err
		}
		item := &RemoteRepository{}
		err = json.Unmarshal([]byte(jsonStr), item)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
