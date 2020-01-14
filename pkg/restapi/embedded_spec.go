// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Device Manager",
    "version": "0.1.0"
  },
  "basePath": "/api",
  "paths": {
    "/devices": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "description": "Получение списка устройств пользователя",
        "tags": [
          "devices"
        ],
        "operationId": "devices_list",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "required": [
                  "device_name",
                  "device_id"
                ],
                "properties": {
                  "device_id": {
                    "type": "string",
                    "format": "uuid",
                    "x-isnullable": false
                  },
                  "device_name": {
                    "type": "string",
                    "x-isnullable": false
                  },
                  "last_meter_readings_time": {
                    "type": "string",
                    "x-isnullable": true
                  }
                }
              },
              "example": [
                {
                  "device_id": "017d4ff8-e2c8-42f2-89f3-7822eeca3ebe",
                  "device_name": "my_device"
                }
              ]
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      },
      "post": {
        "description": "Регистрации нового устройства",
        "tags": [
          "devices"
        ],
        "operationId": "device_registration",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация об устройстве",
              "type": "object",
              "required": [
                "device_name",
                "owner_email"
              ],
              "properties": {
                "device_name": {
                  "type": "string",
                  "x-isnullable": false
                },
                "owner_email": {
                  "type": "string",
                  "x-isnullable": false
                }
              },
              "example": {
                "device_name": "my_device",
                "owner_email": "test@example.com"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "required": [
                "token"
              ],
              "properties": {
                "token": {
                  "type": "string",
                  "x-isnullable": false
                }
              },
              "example": {
                "token": "JWT"
              }
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      }
    },
    "/devices/{device_id}/stats": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "description": "Получение списка устройств пользователя",
        "tags": [
          "devices"
        ],
        "operationId": "device_stats",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "идентификатор устройства",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      }
    },
    "/users": {
      "parameters": [
        {
          "name": "body",
          "in": "body",
          "required": true,
          "schema": {
            "description": "Информация о пользователе",
            "type": "object",
            "required": [
              "email"
            ],
            "properties": {
              "email": {
                "type": "string",
                "x-isnullable": false
              }
            },
            "example": {
              "email": "test@example.com"
            }
          }
        }
      ]
    }
  },
  "definitions": {
    "errorResult": {
      "description": "Ответ API с ошибкой",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "object",
          "required": [
            "code",
            "description"
          ],
          "properties": {
            "code": {
              "description": "HTTP-статус ответа",
              "type": "string",
              "x-isnullable": false,
              "example": "003-061"
            },
            "description": {
              "description": "Человекопонятное описание ошибки",
              "type": "string",
              "x-isnullable": false,
              "example": "Object not found"
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Device Manager",
    "version": "0.1.0"
  },
  "basePath": "/api",
  "paths": {
    "/devices": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "description": "Получение списка устройств пользователя",
        "tags": [
          "devices"
        ],
        "operationId": "devices_list",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "required": [
                  "device_name",
                  "device_id"
                ],
                "properties": {
                  "device_id": {
                    "type": "string",
                    "format": "uuid",
                    "x-isnullable": false
                  },
                  "device_name": {
                    "type": "string",
                    "x-isnullable": false
                  },
                  "last_meter_readings_time": {
                    "type": "string",
                    "x-isnullable": true
                  }
                }
              },
              "example": [
                {
                  "device_id": "017d4ff8-e2c8-42f2-89f3-7822eeca3ebe",
                  "device_name": "my_device"
                }
              ]
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      },
      "post": {
        "description": "Регистрации нового устройства",
        "tags": [
          "devices"
        ],
        "operationId": "device_registration",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "description": "Информация об устройстве",
              "type": "object",
              "required": [
                "device_name",
                "owner_email"
              ],
              "properties": {
                "device_name": {
                  "type": "string",
                  "x-isnullable": false
                },
                "owner_email": {
                  "type": "string",
                  "x-isnullable": false
                }
              },
              "example": {
                "device_name": "my_device",
                "owner_email": "test@example.com"
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "required": [
                "token"
              ],
              "properties": {
                "token": {
                  "type": "string",
                  "x-isnullable": false
                }
              },
              "example": {
                "token": "JWT"
              }
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      }
    },
    "/devices/{device_id}/stats": {
      "get": {
        "security": [
          {
            "bearerAuth": []
          }
        ],
        "description": "Получение списка устройств пользователя",
        "tags": [
          "devices"
        ],
        "operationId": "device_stats",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "идентификатор устройства",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object"
            }
          },
          "422": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/errorResult"
            },
            "examples": {
              "application/json": {
                "error": {
                  "code": "000-000",
                  "description": "description"
                }
              }
            }
          }
        }
      }
    },
    "/users": {
      "parameters": [
        {
          "name": "body",
          "in": "body",
          "required": true,
          "schema": {
            "description": "Информация о пользователе",
            "type": "object",
            "required": [
              "email"
            ],
            "properties": {
              "email": {
                "type": "string",
                "x-isnullable": false
              }
            },
            "example": {
              "email": "test@example.com"
            }
          }
        }
      ]
    }
  },
  "definitions": {
    "errorResult": {
      "description": "Ответ API с ошибкой",
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "object",
          "required": [
            "code",
            "description"
          ],
          "properties": {
            "code": {
              "description": "HTTP-статус ответа",
              "type": "string",
              "x-isnullable": false,
              "example": "003-061"
            },
            "description": {
              "description": "Человекопонятное описание ошибки",
              "type": "string",
              "x-isnullable": false,
              "example": "Object not found"
            }
          }
        }
      }
    }
  }
}`))
}
