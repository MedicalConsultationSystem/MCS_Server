// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dept/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "科室"
                ],
                "summary": "增加一个科室",
                "parameters": [
                    {
                        "description": "科室名称",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddDept"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"科室添加成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dept/listAll": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "科室"
                ],
                "summary": "获取所有科室",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"科室信息获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/drug/listAll": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "药物"
                ],
                "summary": "获取所有药物信息",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"药物信息获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/organization/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "机构"
                ],
                "summary": "增加一个机构",
                "parameters": [
                    {
                        "description": "机构名称",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddOrg"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"机构添加成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/organization/listAll": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "机构"
                ],
                "summary": "获取所有机构",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"机构信息获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.AddDept": {
            "type": "object",
            "properties": {
                "dept_name": {
                    "type": "string"
                }
            }
        },
        "request.AddOrg": {
            "type": "object",
            "properties": {
                "org_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample Server pets",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
