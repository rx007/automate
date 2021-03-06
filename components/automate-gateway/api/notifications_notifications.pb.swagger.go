package api

func init() {
	Swagger.Add("notifications_notifications", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/notifications/notifications.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/notifications/rules": {
      "get": {
        "operationId": "ListRules",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsRuleListResponse"
            }
          }
        },
        "tags": [
          "Notifications"
        ]
      },
      "post": {
        "operationId": "AddRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsRuleAddResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/notificationsRuleAddRequest"
            }
          }
        ],
        "tags": [
          "Notifications"
        ]
      }
    },
    "/notifications/rules/{id}": {
      "get": {
        "operationId": "GetRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsRuleGetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Notifications"
        ]
      },
      "delete": {
        "operationId": "DeleteRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsRuleDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Notifications"
        ]
      },
      "put": {
        "operationId": "UpdateRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsRuleUpdateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/notificationsRuleUpdateRequest"
            }
          }
        ],
        "tags": [
          "Notifications"
        ]
      }
    },
    "/notifications/version": {
      "get": {
        "operationId": "Version",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsVersionResponse"
            }
          }
        },
        "tags": [
          "Notifications"
        ]
      }
    },
    "/notifications/webhook": {
      "post": {
        "operationId": "ValidateWebhook",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/notificationsURLValidationResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/notificationsURLValidationRequest"
            }
          }
        ],
        "tags": [
          "Notifications"
        ]
      }
    }
  },
  "definitions": {
    "RuleEvent": {
      "type": "string",
      "enum": [
        "CCRFailure",
        "CCRSuccess",
        "ComplianceFailure",
        "ComplianceSuccess",
        "Assets"
      ],
      "default": "CCRFailure"
    },
    "notificationsEmpty": {
      "type": "object"
    },
    "notificationsRule": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "event": {
          "$ref": "#/definitions/RuleEvent"
        },
        "SlackAlert": {
          "$ref": "#/definitions/notificationsSlackAlert"
        },
        "WebhookAlert": {
          "$ref": "#/definitions/notificationsWebhookAlert"
        },
        "ServiceNowAlert": {
          "$ref": "#/definitions/notificationsServiceNowAlert"
        }
      }
    },
    "notificationsRuleAddRequest": {
      "type": "object",
      "properties": {
        "rule": {
          "$ref": "#/definitions/notificationsRule"
        }
      }
    },
    "notificationsRuleAddResponse": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "id": {
          "type": "string"
        }
      }
    },
    "notificationsRuleDeleteResponse": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "notificationsRuleGetResponse": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "rule": {
          "$ref": "#/definitions/notificationsRule"
        }
      }
    },
    "notificationsRuleListResponse": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "rules": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/notificationsRule"
          }
        }
      }
    },
    "notificationsRuleUpdateRequest": {
      "type": "object",
      "properties": {
        "rule": {
          "$ref": "#/definitions/notificationsRule"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "notificationsRuleUpdateResponse": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "notificationsSecretId": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "notificationsServiceNowAlert": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "secret_id": {
          "type": "string"
        },
        "critical_controls_only": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    },
    "notificationsSlackAlert": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        }
      }
    },
    "notificationsURLValidationRequest": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "username_password": {
          "$ref": "#/definitions/notificationsUsernamePassword"
        },
        "secret_id": {
          "$ref": "#/definitions/notificationsSecretId"
        },
        "none": {
          "$ref": "#/definitions/notificationsEmpty"
        }
      }
    },
    "notificationsURLValidationResponse": {
      "type": "object"
    },
    "notificationsUsernamePassword": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "notificationsVersionResponse": {
      "type": "object",
      "properties": {
        "version": {
          "type": "string"
        }
      }
    },
    "notificationsWebhookAlert": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        }
      }
    }
  }
}
`)
}
