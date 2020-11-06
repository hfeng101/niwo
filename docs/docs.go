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
        "/catalog/economicsRecordList": {
            "get": {
                "description": "Get record list of economics record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetEconomicsRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/entertainmentRecordList": {
            "get": {
                "description": "Get record list of entertainment record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetEntertainmentRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/list": {
            "get": {
                "description": "Get record list by keyword search.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/listByKey": {
            "get": {
                "description": "Get record list by keyword search.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetRecordListByKeywordReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/militaryRecordList": {
            "get": {
                "description": "Get record list of military record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetMilitaryRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/personageRecordList": {
            "get": {
                "description": "Get record list of personage record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetPersonageRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/catalog/sportRecordList": {
            "get": {
                "description": "Get record list of sport record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetSportRecordListReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/content/theme": {
            "get": {
                "description": "Get content of a record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "Param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.GetContentReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    }
                }
            },
            "put": {
                "description": "Update content of a record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "Param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateContentReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    }
                }
            },
            "post": {
                "description": "Add content of a record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "Param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.AddContentReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_content_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserRegistrationReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserRegistrationReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    }
                }
            }
        },
        "/user/registration": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserRegistrationReq"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    },
                    "50001": {
                        "schema": {
                            "$ref": "#/definitions/github.com_hfeng101_niwo_controller_user_v1.ResponseContent"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github.com_hfeng101_niwo_controller_catalog_v1.ResponseContent": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github.com_hfeng101_niwo_controller_content_v1.ResponseContent": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github.com_hfeng101_niwo_controller_user_v1.ResponseContent": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.AddContentReq": {
            "type": "object",
            "properties": {
                "catalogType": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "keyword": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                }
            }
        },
        "v1.GetContentReq": {
            "type": "object",
            "properties": {
                "catalogType": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "v1.GetEconomicsRecordListReq": {
            "type": "object"
        },
        "v1.GetEntertainmentRecordListReq": {
            "type": "object"
        },
        "v1.GetMilitaryRecordListReq": {
            "type": "object"
        },
        "v1.GetPersonageRecordListReq": {
            "type": "object"
        },
        "v1.GetRecordListByKeywordReq": {
            "type": "object",
            "properties": {
                "keyword": {
                    "type": "string"
                }
            }
        },
        "v1.GetRecordListReq": {
            "type": "object"
        },
        "v1.GetSportRecordListReq": {
            "type": "object"
        },
        "v1.UpdateContentReq": {
            "type": "object",
            "properties": {
                "catalogType": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "keyword": {
                    "type": "string"
                },
                "theme": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "v1.UserRegistrationReq": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
