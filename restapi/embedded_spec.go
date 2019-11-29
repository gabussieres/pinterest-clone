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
    "application/io.goswagger.examples.pinterest-clone.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.pinterest-clone.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API spec for pinterest-clone",
    "title": "pinterest-clone API",
    "version": "1.0.0"
  },
  "paths": {
    "/feed": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "feed",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_amount",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns pin_amount (number) images for the user's home feed",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/pin"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pin": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "pin",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "a pin",
            "schema": {
              "$ref": "#/definitions/pin"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "pinterest"
        ],
        "operationId": "delete_pin",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "pin is deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pins": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "pins",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the user's pins",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/pin"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "user",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "a user",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "pin": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "image_url": {
          "type": "string",
          "minLength": 1
        },
        "posted_by": {
          "type": "string",
          "minLength": 1
        },
        "source_url": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "followers": {
          "type": "integer",
          "format": "int64"
        },
        "following": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "image_url": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.pinterest-clone.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.pinterest-clone.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API spec for pinterest-clone",
    "title": "pinterest-clone API",
    "version": "1.0.0"
  },
  "paths": {
    "/feed": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "feed",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_amount",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns pin_amount (number) images for the user's home feed",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/pin"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pin": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "pin",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "a pin",
            "schema": {
              "$ref": "#/definitions/pin"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "pinterest"
        ],
        "operationId": "delete_pin",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "pin_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "pin is deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/pins": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "pins",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "the user's pins",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/pin"
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "tags": [
          "pinterest"
        ],
        "operationId": "user",
        "parameters": [
          {
            "type": "string",
            "format": "string",
            "name": "user_id",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "a user",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "pin": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "image_url": {
          "type": "string",
          "minLength": 1
        },
        "posted_by": {
          "type": "string",
          "minLength": 1
        },
        "source_url": {
          "type": "string",
          "minLength": 1
        },
        "title": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "followers": {
          "type": "integer",
          "format": "int64"
        },
        "following": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "image_url": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  }
}`))
}
