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
        "termsOfService": "https://www.topgoer.com",
        "contact": {
            "name": "www.topgoer.com",
            "url": "https://www.topgoer.com",
            "email": "790227542@qq.com"
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
        "/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "鉴权相关接口"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "注销登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "鉴权相关接口"
                ],
                "summary": "注销登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/bindpermissions": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色绑定权限接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "角色绑定权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "权限数组",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/bindroles": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户绑定角色接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户绑定角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "角色数组",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Models.RoleForm"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/createrole": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新建角色接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "新建角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "description": "角色结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.RoleForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/createuser": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建账号接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "创建账号",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "description": "表单对象",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.UserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/deleterole": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuths": []
                    }
                ],
                "description": "删除角色接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/deleteuser": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuths": []
                    }
                ],
                "description": "删除用户接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/editrole": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "编辑角色接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "编辑角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "角色结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.RoleForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/edituser": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "编辑用户接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "编辑用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "用户表单数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.UserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getloginuser": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询登录用户信息接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "查询登录用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getpermissionall": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询所有权限接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限相关接口"
                ],
                "summary": "查询所有权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getpermissionlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询权限列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限相关接口"
                ],
                "summary": "查询权限列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "单页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getroleitem": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色详情接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "角色详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getrolelist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询角色列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "查询角色列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "Name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "单页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getrolepermissions": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色权限接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "角色权限",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "角色ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getuseritem": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询用户详情接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "查询用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/getuserlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询用户列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "查询用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "单页数量",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        },
        "/v1/uploadfile": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "上传文件接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "文件传输接口"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "token",
                        "in": "header"
                    },
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Middlewares.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Middlewares.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                }
            }
        },
        "Models.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Models.Permission": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "deletedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "describe": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parent": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "updatedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "Models.Role": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "deletedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Permission"
                    }
                },
                "updatedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "user": {
                    "$ref": "#/definitions/Models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "Models.RoleForm": {
            "type": "object",
            "properties": {
                "Permissions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "createdAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "deletedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Permission"
                    }
                },
                "updatedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "user": {
                    "$ref": "#/definitions/Models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "Models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "deletedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "describe": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Role"
                    }
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Models.UserForm": {
            "type": "object",
            "properties": {
                "Roles": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "deletedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "describe": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Role"
                    }
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "$ref": "#/definitions/Models.XTime"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "Models.XTime": {
            "type": "object",
            "properties": {
                "time.Time": {
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
	Host:        "localhost:8088",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "自定义项目API",
	Description: "This is a sample server celler server.",
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
