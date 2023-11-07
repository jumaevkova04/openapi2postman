package main

import (
	"flag"
	"fmt"
	"github.com/jumaevkova04/openapi2postman/postman"
	"log"
)

var (
	apiKey       = flag.String("api-key", "", "Your Postman API key")
	workspaceID  = flag.String("workspace-id", "", "The Postman workspace ID")
	collectionID = flag.String("collection-id", "", "The Postman collection ID")
	file         = flag.String("file", "", "OpenAPI JSON file")
)

func main() {
	flag.Parse()

	p := postman.Postman{
		WorkspaceID:  *workspaceID,
		ApiKey:       *apiKey,
		CollectionID: *collectionID,
	}

	// -------------------------------------------------------------------------

	filePath := *file

	fileFields, err := p.ReadOpenApiFileAndTakeRequiredFields(filePath)
	checkErrorAndPrintResponse("READ_FILE", filePath, err)

	checkErrorAndPrintResponse("FILE_FORMAT", fileFields.Format, nil)
	checkErrorAndPrintResponse("API_SCHEMA_TYPE", fileFields.APISchemaType, nil)
	checkErrorAndPrintResponse("COLLECTION_NAME", fileFields.CollectionName, nil)

	// -------------------------------------------------------------------------

	apiRequest := postman.CreateAPIRequest{Api: postman.Api{
		Name:        fileFields.CollectionName,
		Summary:     fmt.Sprintf("%s Schema", fileFields.CollectionName),
		Description: fmt.Sprintf("This is a %s", fileFields.CollectionName),
	}}

	api, err := p.CreateAPI(apiRequest)
	checkErrorAndPrintResponse("CREATE_API", api, err)

	// -------------------------------------------------------------------------

	defer func(p *postman.Postman, apiID string) {
		deletedApi, err := p.DeleteAPI(apiID)
		checkErrorAndPrintResponse("DELETE_API", deletedApi, err)
	}(&p, api.Api.ID)

	// -------------------------------------------------------------------------

	apiVersions, err := p.GetAPIVersions(api.Api.ID)
	checkErrorAndPrintResponse("GET_API_VERSIONS", apiVersions, err)

	if len(apiVersions.Versions) < 1 {
		err := fmt.Errorf("empty versions")
		checkErrorAndPrintResponse("GET_API_VERSIONS", apiVersions, err)
	}

	// -------------------------------------------------------------------------

	apiSchemaRequest := postman.CreateAPISchemaRequest{Schema: postman.Schema{
		Language: fileFields.Format,
		Schema:   string(fileFields.Content),
		Type:     fileFields.APISchemaType,
	}}

	apiSchema, err := p.CreateAPISchema(api.Api.ID, apiVersions.Versions[0].ID, apiSchemaRequest)
	checkErrorAndPrintResponse("CREATE_API_SCHEMA", apiSchema, err)

	// -------------------------------------------------------------------------

	apiCollectionRequest := postman.CreateAPICollectionFromSchemaRequest{
		Name: fileFields.CollectionName,
		Relations: []postman.Relations{
			{
				Type: "documentation",
			},
		},
	}

	apiCollection, err := p.CreateAPICollectionFromSchema(api.Api.ID, apiVersions.Versions[0].ID, apiSchema.Schema.ID, apiCollectionRequest)
	checkErrorAndPrintResponse("CREATE_API_COLLECTION", apiCollection, err)

	// -------------------------------------------------------------------------

	newCollection, err := p.GetCollection(apiCollection.Collection.ID)
	checkErrorAndPrintResponse("GET_API_COLLECTION", "OK", err)

	// -------------------------------------------------------------------------

	deletedCollection, err := p.DeleteCollection(apiCollection.Collection.ID)
	checkErrorAndPrintResponse("DELETE_API_COLLECTION", deletedCollection, err)

	// -------------------------------------------------------------------------

	replacedCollection, err := p.ReplaceCollectionsData(*collectionID, newCollection)
	checkErrorAndPrintResponse("REPLACE_API_COLLECTION", replacedCollection, err)
}

func checkErrorAndPrintResponse(info string, response interface{}, err error) {
	if err != nil {
		log.Fatalln("ERR:", err)
		return
	}
	log.Printf("INFO: %s: %+v\n\n", info, response)
}
