# openapi2postman action

GitHub action to push openapi (OpenAPI 3.0, 3.1 and Swagger 2.0) file to Postman directly from your workflow

## Usage

Add the openapi2postman action and set the required inputs

* `api-key`: your Postman API key
* `workspace-id`: your Postman workspace id
* `collection-id`: your Postman collection id
* `file`: your openapi json file

### File formats available

`.json`   
`.yaml`   
`.yml`

### Update existing collection

Update existing Postman collection

```yaml
    - name: "Update Postman Collection"
      uses: jumaevkova04/openapi2postman@main
      with:
        api-key: ${{ secrets.POSTMAN_API_KEY }}
        workspace-id: ${{ secrets.POSTMAN_WORKSPACE_ID }}
        collection-id: ${{ secrets.POSTMAN_COLLECTION_ID }}
        file: ./docs/swagger.json
```

### Example workflow file

Update Postman collections on `push`

```yaml
name: "Update Postman collection"

on:
  push:
    branches: [ "main" ]

jobs:
  sync-documentation-with-postman:
    runs-on: ubuntu-latest
    steps:
      - name: "Checkout repository"
        uses: actions/checkout@v3

      - name: "Update Postman Collection"
        uses: jumaevkova04/openapi2postman@main
        with:
          api-key: ${{ secrets.POSTMAN_API_KEY }}
          workspace-id: ${{ secrets.POSTMAN_WORKSPACE_ID }}
          collection-id: ${{ secrets.POSTMAN_COLLECTION_ID }}
          file: ./docs/swagger.json
```