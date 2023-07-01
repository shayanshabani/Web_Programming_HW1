// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/req_dh": {
            "post": {
                "description": "Get Dq from server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "GetDh",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.ReqDh_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.ReqDh_Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/auth/req_pq": {
            "post": {
                "description": "Get Pq from server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "GetPq",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.ReqPq_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.ReqPq_Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/get-user": {
            "post": {
                "description": "Get User from server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Biz"
                ],
                "summary": "GetUser",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.GetUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/get-user-inject": {
            "post": {
                "description": "Get User With Injection from server",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Biz"
                ],
                "summary": "GetUserWithInjection",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/proto.GetUsersWithInjectionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/proto.GetUsersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "proto.GetUsersRequest": {
            "type": "object",
            "properties": {
                "authKey": {
                    "type": "integer"
                },
                "messageId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "proto.GetUsersResponse": {
            "type": "object",
            "properties": {
                "messageId": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/proto.User"
                    }
                }
            }
        },
        "proto.GetUsersWithInjectionRequest": {
            "type": "object",
            "properties": {
                "authKey": {
                    "type": "integer"
                },
                "messageId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "proto.ReqDh_Request": {
            "type": "object",
            "properties": {
                "A": {
                    "description": "public key generated by client",
                    "type": "integer"
                },
                "message_id": {
                    "description": "a random even integer, greater than of that in req_pq",
                    "type": "integer"
                },
                "nonce": {
                    "description": "random string generated by client in req_pq",
                    "type": "string"
                },
                "server_nonce": {
                    "description": "random string generated by server in req_pq",
                    "type": "string"
                }
            }
        },
        "proto.ReqDh_Response": {
            "type": "object",
            "properties": {
                "B": {
                    "description": "public key generated by server",
                    "type": "integer"
                },
                "message_id": {
                    "description": "random odd integer",
                    "type": "integer"
                },
                "nonce": {
                    "description": "the same string that client sent",
                    "type": "string"
                },
                "server_nonce": {
                    "description": "another random string of size 20",
                    "type": "string"
                }
            }
        },
        "proto.ReqPq_Request": {
            "type": "object",
            "properties": {
                "message_id": {
                    "type": "integer"
                },
                "nonce": {
                    "description": "random string sent by client with size 20",
                    "type": "string"
                }
            }
        },
        "proto.ReqPq_Response": {
            "type": "object",
            "properties": {
                "g": {
                    "description": "a generating number like 5",
                    "type": "integer"
                },
                "message_id": {
                    "type": "integer"
                },
                "nonce": {
                    "description": "the same string that client sent with size 20",
                    "type": "string"
                },
                "p": {
                    "description": "a prime number like 23",
                    "type": "integer"
                },
                "server_nonce": {
                    "description": "another random string with size 20, that server sends to client",
                    "type": "string"
                }
            }
        },
        "proto.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "family": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}