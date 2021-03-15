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
        "termsOfService": "http://pando.im/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.pando.im/support",
            "email": "support@pando.im"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/actions": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actions"
                ],
                "summary": "request payment code",
                "parameters": [
                    {
                        "description": "request payments",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/actions.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/actions.CreateResponse"
                        }
                    }
                }
            }
        },
        "/assets": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "list assets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListAssets"
                        }
                    }
                }
            }
        },
        "/assets/{asset_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Assets"
                ],
                "summary": "Find Asset By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mixin asset id",
                        "name": "asset_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Asset"
                        }
                    }
                }
            }
        },
        "/cats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaterals"
                ],
                "summary": "list all collateral",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListCollaterals"
                        }
                    }
                }
            }
        },
        "/cats/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collaterals"
                ],
                "summary": "find collateral by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "collateral id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Collateral"
                        }
                    }
                }
            }
        },
        "/flips": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flips"
                ],
                "summary": "list flips",
                "parameters": [
                    {
                        "type": "string",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListFlips"
                        }
                    }
                }
            }
        },
        "/flips/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flips"
                ],
                "summary": "find flip by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flip id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Flip"
                        }
                    }
                }
            }
        },
        "/flips/{id}/events": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Flips"
                ],
                "summary": "list flip events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "flip id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListFlipEvents"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Show system info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/system.InfoResponse"
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
                "tags": [
                    "user"
                ],
                "summary": "login with mixin oauth code",
                "parameters": [
                    {
                        "description": "request login",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/user.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    }
                }
            }
        },
        "/me/vats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vaults"
                ],
                "summary": "list my vaults",
                "parameters": [
                    {
                        "type": "string",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Example: Bearer foo",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListVaults"
                        }
                    }
                }
            }
        },
        "/oracles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Oracles"
                ],
                "summary": "list all oracles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListOracles"
                        }
                    }
                }
            }
        },
        "/oracles/{asset_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Oracles"
                ],
                "summary": "find oracle by asset id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mixin asset id",
                        "name": "asset_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Oracle"
                        }
                    }
                }
            }
        },
        "/time": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Show server time",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/system.TimeResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "list transactions",
                "parameters": [
                    {
                        "type": "string",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListTransactions"
                        }
                    }
                }
            }
        },
        "/transactions/{follow_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "find tx by follow id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Example: Bearer foo",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "follow id",
                        "name": "follow_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Transaction"
                        }
                    }
                }
            }
        },
        "/vats": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vaults"
                ],
                "summary": "list vaults",
                "parameters": [
                    {
                        "type": "string",
                        "name": "collateral_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListVaults"
                        }
                    }
                }
            }
        },
        "/vats/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vaults"
                ],
                "summary": "find vault by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "vault id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Vault"
                        }
                    }
                }
            }
        },
        "/vats/{id}/events": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vaults"
                ],
                "summary": "list vault events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "vault id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "@inject_tag: valid:\"uuid,required\"",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Resp_ListVaultEvents"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "actions.CreateRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "payment amount (optional)",
                    "type": "number"
                },
                "asset_id": {
                    "description": "payment asset id (optional)",
                    "type": "string",
                    "format": "uuid"
                },
                "follow_id": {
                    "description": "follow id to track tx (uuid)",
                    "type": "string",
                    "format": "uuid"
                },
                "parameters": {
                    "description": "tx parameters",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "actions.CreateResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "multisig payment code",
                    "type": "string"
                },
                "code_url": {
                    "description": "multisig payment code url",
                    "type": "string"
                },
                "memo": {
                    "description": "payment memo",
                    "type": "string"
                }
            }
        },
        "api.Asset": {
            "type": "object",
            "properties": {
                "chain": {
                    "$ref": "#/definitions/api.Asset"
                },
                "chain_id": {
                    "type": "string"
                },
                "id": {
                    "description": "mixin asset id",
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "api.Collateral": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "beg": {
                    "type": "string"
                },
                "box": {
                    "type": "string"
                },
                "chop": {
                    "type": "string"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "dai": {
                    "type": "string"
                },
                "debt": {
                    "type": "string"
                },
                "dunk": {
                    "type": "string"
                },
                "dust": {
                    "type": "string"
                },
                "duty": {
                    "type": "string"
                },
                "gem": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ink": {
                    "type": "string"
                },
                "line": {
                    "type": "string"
                },
                "litter": {
                    "type": "string"
                },
                "live": {
                    "type": "boolean"
                },
                "mat": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "number_of_vaults": {
                    "type": "integer"
                },
                "price": {
                    "type": "string"
                },
                "rate": {
                    "type": "string"
                },
                "rho": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "tau": {
                    "type": "integer"
                },
                "ttl": {
                    "type": "integer"
                }
            }
        },
        "api.Flip": {
            "type": "object",
            "properties": {
                "Guy": {
                    "type": "string"
                },
                "action": {
                    "type": "integer"
                },
                "art": {
                    "type": "string"
                },
                "bid": {
                    "type": "string"
                },
                "collateral_id": {
                    "type": "string"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "end": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "id": {
                    "type": "string"
                },
                "lot": {
                    "type": "string"
                },
                "tab": {
                    "type": "string"
                },
                "tic": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "vault_id": {
                    "type": "string"
                }
            }
        },
        "api.Flip_Event": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                },
                "bid": {
                    "type": "string"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "flip_id": {
                    "type": "string"
                },
                "lot": {
                    "type": "string"
                }
            }
        },
        "api.Oracle": {
            "type": "object",
            "properties": {
                "asset_id": {
                    "type": "string"
                },
                "current": {
                    "type": "string"
                },
                "hop": {
                    "type": "integer"
                },
                "next": {
                    "type": "string"
                },
                "peek_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"\nlast update of current price",
                    "type": "string",
                    "format": "date"
                }
            }
        },
        "api.Pagination": {
            "type": "object",
            "properties": {
                "has_next": {
                    "type": "boolean"
                },
                "next_cursor": {
                    "type": "string"
                }
            }
        },
        "api.Resp_ListAssets": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Asset"
                    }
                }
            }
        },
        "api.Resp_ListCollaterals": {
            "type": "object",
            "properties": {
                "collaterals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Collateral"
                    }
                }
            }
        },
        "api.Resp_ListFlipEvents": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Flip_Event"
                    }
                }
            }
        },
        "api.Resp_ListFlips": {
            "type": "object",
            "properties": {
                "flips": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Flip"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/api.Pagination"
                }
            }
        },
        "api.Resp_ListOracles": {
            "type": "object",
            "properties": {
                "oracles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Oracle"
                    }
                }
            }
        },
        "api.Resp_ListTransactions": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/api.Pagination"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Transaction"
                    }
                }
            }
        },
        "api.Resp_ListVaultEvents": {
            "type": "object",
            "properties": {
                "events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Vault_Event"
                    }
                }
            }
        },
        "api.Resp_ListVaults": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/api.Pagination"
                },
                "vaults": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Vault"
                    }
                }
            }
        },
        "api.Transaction": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                },
                "amount": {
                    "type": "string"
                },
                "asset_id": {
                    "type": "string"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "id": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "parameters": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "api.Vault": {
            "type": "object",
            "properties": {
                "art": {
                    "type": "string"
                },
                "collateral_id": {
                    "type": "string"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "id": {
                    "type": "string"
                },
                "identity_id": {
                    "type": "integer"
                },
                "ink": {
                    "type": "string"
                }
            }
        },
        "api.Vault_Event": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                },
                "created_at": {
                    "description": "@inject_tag: swaggertype:\"string\" format:\"date\"",
                    "type": "string",
                    "format": "date"
                },
                "dart": {
                    "type": "string"
                },
                "debt": {
                    "type": "string"
                },
                "dink": {
                    "type": "string"
                },
                "vault_id": {
                    "type": "string"
                }
            }
        },
        "system.InfoResponse": {
            "type": "object",
            "properties": {
                "members": {
                    "description": "multisig members",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "oauth_client_id": {
                    "description": "oauth client id",
                    "type": "string",
                    "format": "uuid"
                },
                "public_key": {
                    "type": "string"
                },
                "threshold": {
                    "description": "multisig threshold",
                    "type": "integer"
                }
            }
        },
        "system.TimeResponse": {
            "type": "object",
            "properties": {
                "epoch": {
                    "type": "integer"
                },
                "iso": {
                    "type": "string"
                }
            }
        },
        "user.LoginRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "mixin oauth code",
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "user avatar",
                    "type": "string"
                },
                "id": {
                    "description": "user mixin id",
                    "type": "string",
                    "format": "uuid"
                },
                "name": {
                    "description": "user name",
                    "type": "string"
                },
                "scope": {
                    "description": "mixin oauth scope",
                    "type": "string"
                },
                "token": {
                    "description": "mixin oauth token",
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
	Version:     "1.0",
	Host:        "pando-test-api.fox.one",
	BasePath:    "/api",
	Schemes:     []string{"https", "http"},
	Title:       "Pando API",
	Description: "Pando Api Doc",
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
