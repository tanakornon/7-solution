// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/{version}/beef": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beef"
                ],
                "summary": "Get beef details",
                "parameters": [
                    {
                        "enum": [
                            "v1"
                        ],
                        "type": "string",
                        "description": "App Version",
                        "name": "version",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "501": {
                        "description": "Not Implemented",
                        "schema": {
                            "$ref": "#/definitions/beef.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/{version}/beef/summary": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Beef"
                ],
                "summary": "Get beef summary",
                "parameters": [
                    {
                        "enum": [
                            "v1"
                        ],
                        "type": "string",
                        "description": "App Version",
                        "name": "version",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/beef.GetSummaryResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/beef.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "beef.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "beef.GetSummaryResponse": {
            "type": "object",
            "properties": {
                "beef": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pie Fire Dire",
	Description:      "This is the API for the Pie Fire Dire application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
