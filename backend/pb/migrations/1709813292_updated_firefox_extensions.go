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

		collection, err := dao.FindCollectionByNameOrId("d0lxxvunn80r1ye")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("d0lxxvunn80r1ye")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.headers.x_service_token = @collection.internal_services.service_token")

		return dao.SaveCollection(collection)
	})
}
