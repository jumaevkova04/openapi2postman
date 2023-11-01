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

	apiRequest := postman.CreateAPIRequest{Api: postman.Api{
		Name:        "Test API",
		Summary:     "Test API Schema",
		Description: "This is a test API",
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

	path := *file

	fileContent, fileFormat, apiSchemaType, err := p.ReadFileAndFindApiSchemaType(path)
	checkErrorAndPrintResponse("READ_FILE", path, err)

	checkErrorAndPrintResponse("FILE_FORMAT", fileFormat, nil)
	checkErrorAndPrintResponse("API_SCHEMA_TYPE", apiSchemaType, nil)

	apiSchemaRequest := postman.CreateAPISchemaRequest{Schema: postman.Schema{
		Language: fileFormat,
		Schema:   string(fileContent),
		Type:     apiSchemaType,
	}}

	apiSchema, err := p.CreateAPISchema(api.Api.ID, apiVersions.Versions[0].ID, apiSchemaRequest)
	checkErrorAndPrintResponse("CREATE_API_SCHEMA", apiSchema, err)

	// -------------------------------------------------------------------------

	apiCollectionRequest := postman.CreateAPICollectionFromSchemaRequest{
		Name: "Test Collection",
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
