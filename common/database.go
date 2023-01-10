package common

import "time"

type DBObject struct {
	ID        string
	Type      string
	Name      string
	Json      string
	UpdatedAt time.Time
}

type Artifact struct {
	ID        string
	BagName   string
	ItemType  string // File or WorkResult
	FileName  string // name of manifest or tag file
	FileType  string // manifest or tag file
	RawData   string // file content or work result json
	UpdatedAt time.Time
}

func InitSchema() error {
	schema := `create table if not exists dart (
		uuid text primary key not null,
		obj_type text not null,
		obj_name text not null,
		obj_json text not null,
		updated_at datetime not null
	);
	create unique index if not exists ix_unique_object_name on dart(obj_type, obj_name);
	create table if not exists artifacts (
		uuid text primary key not null,
		bag_name text not null,
		item_type text not null,
		file_name text,
		file_type text,
		raw_data text not null,
		updated_at datetime not null
	);
	create index if not exists ix_artifact_bag_name on artifacts(bag_name);
	`
	_, err := Dart.DB.Exec(schema)
	return err
}

func ObjSave(obj DartObject) error {

	return nil
}

func ObjGet(uuid string) *QueryResult {

	return nil
}

func ObjList(typeName, orderBy string, offset, limit int) *QueryResult {

	return nil
}

func ObjDelete(uuid string) *QueryResult {

	return nil
}

func ArtifactSave(a *Artifact) error {

	return nil
}

func ArtifactGet(uuid string) *QueryResult {

	return nil
}

func ArtifactList(bagName, orderBy string, offset, limit int) *QueryResult {

	return nil
}

func ArtifactDelete(uuid string) *QueryResult {

	return nil
}
