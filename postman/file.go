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

type OpenApiFileFields struct {
	FileType
	Info FileInfo `json:"info" yaml:"info"`

	Content []byte `json:"-" yaml:"-"`
}

type FileType struct {
	Swagger string `json:"swagger" yaml:"swagger"`
	Openapi string `json:"openapi" yaml:"openapi"`
}

type FileInfo struct {
	Title   string   `json:"title" yaml:"title"`
	Contact struct{} `json:"contact" yaml:"contact"`
	Version string   `json:"version" yaml:"version"`
}

type OpenApiFileResponse struct {
	Format         string
	APISchemaType  string
	CollectionName string
	Content        []byte
}

func (p *Postman) ReadOpenApiFileAndTakeRequiredFields(filePath string) (openApiFileResponse *OpenApiFileResponse, err error) {
	fileFormat, err := getFileFormat(filePath)
	if err != nil {
		return nil, err
	}

	file, err := readFileAndUnmarshal(filePath, fileFormat)
	if err != nil {
		return nil, err
	}

	schemaType := getAPISchemaType(file.FileType)

	openApiFileResponse = &OpenApiFileResponse{
		Format:         fileFormat,
		APISchemaType:  schemaType,
		CollectionName: file.Info.Title,
		Content:        file.Content,
	}

	return openApiFileResponse, nil
}

func getFileFormat(filePath string) (string, error) {
	fileExtension := strings.Split(filePath, ".")
	if len(fileExtension) < 2 {
		return "", fmt.Errorf("invalid file path: %s", filePath)
	}

	fileFormat := fileExtension[len(fileExtension)-1]

	if fileFormat != fileFormatJson && fileFormat != fileFormatYaml && fileFormat != fileFormatYml {
		return "", fmt.Errorf("invalid file format: %q, must be .%s or .%s (.%s)", fileFormat, fileFormatJson, fileFormatYaml, fileFormatYml)
	}

	return fileFormat, nil
}

func readFileAndUnmarshal(filePath string, fileFormat string) (*OpenApiFileFields, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var file OpenApiFileFields

	if fileFormat == fileFormatJson {
		err := json.Unmarshal(content, &file)
		if err != nil {
			return nil, err
		}
	} else {
		err := yaml.Unmarshal(content, &file)
		if err != nil {
			return nil, err
		}
	}

	file.Content = content

	return &file, nil
}

func getAPISchemaType(fileType FileType) string {
	schemaType := fileType.Swagger

	if strings.TrimSpace(schemaType) == "" {
		schemaType = fileType.Openapi
	}

	if strings.TrimSpace(schemaType) == "" {
		return defaultSchemaType
	}

	apiSchemaTypeFirstNumber := strings.Split(schemaType, ".")[0]
	return prefixSchemaType + apiSchemaTypeFirstNumber
}
