package postman

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

const (
	fileFormatJson = "json"
	fileFormatYaml = "yaml"
	fileFormatYml  = "yml"

	defaultSchemaType = "openapi3"
	prefixSchemaType  = "openapi"
)

type ApiSchemaType struct {
	Swagger string `json:"swagger" yaml:"swagger"`
	Openapi string `json:"openapi" yaml:"openapi"`
}

func (p *Postman) ReadFileAndFindApiSchemaType(filePath string) (content []byte, format string, schemaType string, err error) {
	fileExtension := strings.Split(filePath, ".")
	if len(fileExtension) < 2 {
		return nil, "", "", fmt.Errorf("invalid file: %s", filePath)
	}

	fileFormat := fileExtension[len(fileExtension)-1]
	if fileFormat != fileFormatJson && fileFormat != fileFormatYaml && fileFormat != fileFormatYml {
		return nil, "", "", fmt.Errorf("invalid file format: %q, must be .%s or .%s (.%s)", fileFormatJson, fileFormatYaml, fileFormatYml, fileFormat)
	}

	content, err = os.ReadFile(filePath)
	if err != nil {
		return nil, "", "", err
	}

	var apiSchemaType ApiSchemaType

	if fileFormat == fileFormatJson {
		err = json.Unmarshal(content, &apiSchemaType)
		if err != nil {
			return nil, "", "", err
		}
	} else {
		err := yaml.Unmarshal(content, &apiSchemaType)
		if err != nil {
			return nil, "", "", err
		}
	}

	schemaType = apiSchemaType.Swagger

	if schemaType == "" {
		schemaType = apiSchemaType.Openapi
	}

	if schemaType == "" {
		schemaType = defaultSchemaType
	} else {
		apiSchemaTypeFirstNumber := strings.Split(schemaType, ".")[0]
		schemaType = prefixSchemaType + apiSchemaTypeFirstNumber
	}

	return content, fileFormat, schemaType, nil
}
