package common

type RemoteRepository struct {
	ID         string
	Name       string
	Url        string
	UserID     string
	APIToken   string
	LoginExtra string
	PluginID   string
}

func (repo *RemoteRepository) ObjID() string {
	return repo.ID
}

func (repo *RemoteRepository) ObjName() string {
	return repo.Name
}

func (repo *RemoteRepository) ObjType() string {
	return "RemoteRepository"
}
