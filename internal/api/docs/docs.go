// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/v1/auth/login": {
            "post": {
                "description": "User login, creating new session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/User"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "Session-id": {
                                "type": "string",
                                "description": "Auth cookie with new valid session id"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "delete": {
                "description": "User logout, session deletion",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/Empty"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "Session-id": {
                                "type": "string",
                                "description": "Auth cookie with expired session id"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signup": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Email",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/Empty"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pin": {
            "get": {
                "description": "Get pin collection",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pin"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/Pin"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pin/{pinId}": {
            "get": {
                "description": "Get concrete pin by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pin"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the pin",
                        "name": "pinId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/JsonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "$ref": "#/definitions/Pin"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/JsonErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Empty": {
            "type": "object"
        },
        "JsonErrResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "0"
                },
                "message": {
                    "type": "string",
                    "example": "Error description"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "JsonResponse": {
            "type": "object",
            "properties": {
                "body": {},
                "message": {
                    "type": "string",
                    "example": "Response message"
                },
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "Pin": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "integer",
                    "example": 23
                },
                "description": {
                    "type": "string",
                    "example": "about face"
                },
                "id": {
                    "type": "integer",
                    "example": 55
                },
                "picture": {
                    "type": "string",
                    "example": "pinspire/imgs/image.png"
                },
                "publicationTime": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "example": "Nature's beauty"
                }
            }
        },
        "User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string",
                    "example": "pinspire.online/avatars/avatar.jpg"
                },
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "example": "digital@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "name": {
                    "type": "string",
                    "example": "Peter"
                },
                "password": {
                    "type": "string",
                    "example": "pass123"
                },
                "surname": {
                    "type": "string",
                    "example": "Green"
                },
                "username": {
                    "type": "string",
                    "example": "Green"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Pinspire API",
	Description:      "API for Pinspire project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
