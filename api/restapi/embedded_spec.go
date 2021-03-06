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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Maze Game API",
    "title": "Maze Game",
    "version": "v1"
  },
  "host": "localhost:5000",
  "basePath": "/v1",
  "paths": {
    "/game/{levelId}": {
      "get": {
        "summary": "Play game by LevelId",
        "parameters": [
          {
            "type": "integer",
            "name": "levelId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    },
    "/level": {
      "get": {
        "summary": "Get all levels",
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/LevelAllResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Record level",
        "parameters": [
          {
            "name": "payload",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LevelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "201",
            "schema": {
              "$ref": "#/definitions/LevelResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    },
    "/level/{levelId}": {
      "get": {
        "description": "Level by LevelId",
        "parameters": [
          {
            "type": "integer",
            "name": "levelId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/Level"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BaseResponse": {
      "type": "object",
      "required": [
        "status",
        "success"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer",
          "x-nullable": false
        },
        "success": {
          "type": "boolean",
          "x-nullable": false
        }
      }
    },
    "Level": {
      "title": "Level",
      "required": [
        "maps"
      ],
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "description": "Level ID",
          "type": "integer",
          "format": "int64",
          "x-nullable": false,
          "example": 12
        },
        "maps": {
          "type": "array",
          "items": {
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int32"
            }
          },
          "x-nullable": false,
          "example": [
            [
              5,
              2,
              4
            ],
            [
              2,
              2,
              1
            ],
            [
              1,
              2,
              5
            ],
            [
              3,
              3,
              1
            ],
            [
              2,
              2,
              0
            ]
          ]
        }
      }
    },
    "LevelAllResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Level"
      }
    },
    "LevelRequest": {
      "allOf": [
        {
          "$ref": "#/definitions/Level"
        }
      ]
    },
    "LevelResponse": {
      "type": "object",
      "required": [
        "status",
        "success"
      ],
      "properties": {
        "levelData": {
          "description": "Player data map",
          "type": "object",
          "allOf": [
            {
              "$ref": "#/definitions/Level"
            }
          ],
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "example": "error message"
        },
        "status": {
          "type": "integer",
          "x-nullable": false
        },
        "success": {
          "description": "Success",
          "type": "boolean",
          "x-nullable": false,
          "example": true
        }
      }
    }
  },
  "parameters": {
    "levelId": {
      "type": "integer",
      "name": "levelId",
      "in": "query",
      "required": true
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Maze Game API",
    "title": "Maze Game",
    "version": "v1"
  },
  "host": "localhost:5000",
  "basePath": "/v1",
  "paths": {
    "/game/{levelId}": {
      "get": {
        "summary": "Play game by LevelId",
        "parameters": [
          {
            "type": "integer",
            "name": "levelId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    },
    "/level": {
      "get": {
        "summary": "Get all levels",
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/LevelAllResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      },
      "post": {
        "summary": "Record level",
        "parameters": [
          {
            "name": "payload",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LevelRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "201",
            "schema": {
              "$ref": "#/definitions/LevelResponse"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    },
    "/level/{levelId}": {
      "get": {
        "description": "Level by LevelId",
        "parameters": [
          {
            "type": "integer",
            "name": "levelId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200",
            "schema": {
              "$ref": "#/definitions/Level"
            }
          },
          "400": {
            "description": "400",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "404": {
            "description": "404",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          },
          "500": {
            "description": "500",
            "schema": {
              "$ref": "#/definitions/BaseResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BaseResponse": {
      "type": "object",
      "required": [
        "status",
        "success"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "status": {
          "type": "integer",
          "x-nullable": false
        },
        "success": {
          "type": "boolean",
          "x-nullable": false
        }
      }
    },
    "Level": {
      "title": "Level",
      "required": [
        "maps"
      ],
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "description": "Level ID",
          "type": "integer",
          "format": "int64",
          "x-nullable": false,
          "example": 12
        },
        "maps": {
          "type": "array",
          "items": {
            "type": "array",
            "items": {
              "type": "integer",
              "format": "int32"
            }
          },
          "x-nullable": false,
          "example": [
            [
              5,
              2,
              4
            ],
            [
              2,
              2,
              1
            ],
            [
              1,
              2,
              5
            ],
            [
              3,
              3,
              1
            ],
            [
              2,
              2,
              0
            ]
          ]
        }
      }
    },
    "LevelAllResponse": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Level"
      }
    },
    "LevelRequest": {
      "allOf": [
        {
          "$ref": "#/definitions/Level"
        }
      ]
    },
    "LevelResponse": {
      "type": "object",
      "required": [
        "status",
        "success"
      ],
      "properties": {
        "levelData": {
          "description": "Player data map",
          "type": "object",
          "allOf": [
            {
              "$ref": "#/definitions/Level"
            }
          ],
          "x-nullable": false
        },
        "message": {
          "type": "string",
          "example": "error message"
        },
        "status": {
          "type": "integer",
          "x-nullable": false
        },
        "success": {
          "description": "Success",
          "type": "boolean",
          "x-nullable": false,
          "example": true
        }
      }
    }
  },
  "parameters": {
    "levelId": {
      "type": "integer",
      "name": "levelId",
      "in": "query",
      "required": true
    }
  }
}`))
}
