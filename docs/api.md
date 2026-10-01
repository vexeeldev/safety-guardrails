# Safety Guardrails - API Specification

## 1. Purpose

This document defines communication between Android Client, Go Backend, Planner Service, and Desktop Agent.

Base API:

/api/v1

Backend technology:

Go + Gin

---

## 2. API Principles

1. Use JSON.
2. Use REST for standard request/response operations.
3. Use WebSocket for real-time task updates.
4. Validate all input.
5. Never trust client input.
6. Never trust planner output.
7. Return consistent errors.
8. Use task IDs for tracking.
9. Use device IDs for identifying desktop agents.

---

## 3. Device API

### Register Device

POST /api/v1/devices

Purpose:
Register a desktop agent.

Request:

{
  "name": "My Linux PC"
}

Response:

{
  "id": "device-id",
  "name": "My Linux PC",
  "status": "offline"
}

---

### Get Devices

GET /api/v1/devices

Purpose:
Return available devices.

Response:

{
  "devices": [
    {
      "id": "device-id",
      "name": "My Linux PC",
      "status": "online"
    }
  ]
}

---

## 4. Task API

### Create Task

POST /api/v1/tasks

Request:

{
  "device_id": "device-id",
  "instruction": "buat folder test di desktop"
}

Response:

{
  "id": "task-id",
  "status": "CREATED"
}

---

### Get Task

GET /api/v1/tasks/{task_id}

Response:

{
  "id": "task-id",
  "instruction": "buat folder test di desktop",
  "status": "EXECUTING"
}

---

### Approve Task

POST /api/v1/tasks/{task_id}/approve

Purpose:
Approve a REVIEW action.

Request:

{
  "approved": true
}

Response:

{
  "task_id": "task-id",
  "status": "APPROVED"
}

Approval must only work for tasks waiting for approval.

---

### Cancel Task

POST /api/v1/tasks/{task_id}/cancel

Response:

{
  "task_id": "task-id",
  "status": "CANCELLED"
}

---

## 5. Task History

GET /api/v1/tasks

Optional query parameters:

status
device_id
limit
offset

Example:

GET /api/v1/tasks?status=COMPLETED

---

## 6. WebSocket

Endpoint:

/ws

Purpose:

Real-time communication between backend and connected desktop agents or clients.

Possible events:

task.created
task.planned
task.guardrail_checked
task.waiting_approval
task.approved
task.executing
task.execution_result
task.verifying
task.verification_result
task.completed
task.failed
task.blocked

---

## 7. WebSocket Message Format

Example:

{
  "type": "task.executing",
  "task_id": "task-id",
  "payload": {
    "action_id": "action-id"
  }
}

---

## 8. Planner API

Planner service:

POST /plan

Request:

{
  "instruction": "buat folder test di desktop"
}

Response:

{
  "actions": [
    {
      "type": "create_directory",
      "target": "~/Desktop/test"
    }
  ]
}

Planner response must use structured data.

The backend must validate planner output before using it.

---

## 9. Agent Messages

Backend -> Agent:

{
  "type": "execute_action",
  "task_id": "task-id",
  "action": {
    "id": "action-id",
    "type": "create_directory",
    "target": "~/Desktop/test"
  }
}

Agent -> Backend:

{
  "type": "execution_result",
  "task_id": "task-id",
  "action_id": "action-id",
  "success": true,
  "message": "Directory created"
}

---

## 10. Verification Result

Agent or verification service returns:

{
  "type": "verification_result",
  "task_id": "task-id",
  "action_id": "action-id",
  "status": "VERIFIED",
  "details": "Directory exists"
}

Possible verification states:

VERIFIED
FAILED
UNCERTAIN

---

## 11. Error Format

All API errors should follow:

{
  "error": {
    "code": "INVALID_ACTION",
    "message": "Action type is not supported"
  }
}

Example codes:

INVALID_REQUEST
UNAUTHORIZED
FORBIDDEN
DEVICE_NOT_FOUND
TASK_NOT_FOUND
INVALID_ACTION
GUARDRAIL_BLOCKED
APPROVAL_REQUIRED
AGENT_OFFLINE
EXECUTION_FAILED
VERIFICATION_FAILED
INTERNAL_ERROR

---

## 12. HTTP Status Codes

200:
Successful request.

201:
Resource created.

400:
Invalid request.

401:
Authentication required.

403:
Request forbidden.

404:
Resource not found.

409:
Invalid state or conflict.

422:
Validation error.

500:
Internal server error.

---

## 13. Authentication

Authentication will be added during implementation.

Device communication should use a secure device credential or token.

Tokens must not be hardcoded in source code.

Secrets must use environment variables or secure configuration.

---

## 14. API Security

Backend must:

- validate request body
- validate IDs
- validate action types
- validate paths
- validate device ownership
- validate task state
- validate approval state

Never trust:

- Android input
- planner output
- desktop agent input
- user-provided paths

---

## 15. API Versioning

Current version:

v1

Future breaking API changes should use:

v2

Do not silently change the meaning of an existing endpoint.
