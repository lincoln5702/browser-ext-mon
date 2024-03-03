package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("bxllu29bavy3izv")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.headers.x_service_token = @collection.internal_services.service_token")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("bxllu29bavy3izv")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.headers.x_api_token = @collection.internal_services.api_token")

		return dao.SaveCollection(collection)
	})
}
