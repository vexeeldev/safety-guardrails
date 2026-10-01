## `component.md`

```text
# Safety Guardrails - Component Diagram

## Purpose

This document describes the internal components of the Safety Guardrails system.

## Component Architecture

```mermaid
flowchart TB

    subgraph Mobile["Android Application"]

        UI["Compose UI"]

        TaskScreen["Task Screen"]

        ApprovalScreen["Approval Screen"]

        HistoryScreen["History Screen"]

        APIClient["REST API Client"]

        WebSocketClient["WebSocket Client"]

        UI --> TaskScreen
        UI --> ApprovalScreen
        UI --> HistoryScreen

        TaskScreen --> APIClient
        ApprovalScreen --> APIClient
        HistoryScreen --> APIClient

        TaskScreen --> WebSocketClient
        ApprovalScreen --> WebSocketClient

    end

    subgraph Backend["Go Backend"]

        REST["REST API"]

        WebSocket["WebSocket Server"]

        TaskService["Task Service"]

        DeviceService["Device Service"]

        PlannerClient["Planner Client"]

        GuardrailService["Guardrail Service"]

        AgentManager["Agent Manager"]

        AuditService["Audit Service"]

        REST --> TaskService
        REST --> DeviceService

        WebSocket --> TaskService
        WebSocket --> AgentManager

        TaskService --> PlannerClient
        TaskService --> GuardrailService
        TaskService --> AgentManager
        TaskService --> AuditService

    end

    subgraph Planner["Planner Service"]

        FastAPI["FastAPI"]

        LLM["LLM"]

        Schema["Pydantic Schema"]

        FastAPI --> LLM
        LLM --> Schema

    end

    subgraph Desktop["Desktop Agent"]

        Transport["WebSocket Transport"]

        LocalGuard["Local Guard"]

        Executor["Action Executor"]

        System["System Interface"]

        Verification["Verification Engine"]

        Transport --> LocalGuard
        LocalGuard --> Executor
        Executor --> System
        System --> Verification

    end

    subgraph Database["PostgreSQL"]

        Users["users"]

        Devices["devices"]

        Tasks["tasks"]

        Actions["actions"]

        Executions["executions"]

        Verifications["verifications"]

        AuditLogs["audit_logs"]

    end

    APIClient --> REST

    WebSocketClient --> WebSocket

    PlannerClient --> FastAPI

    TaskService --> Tasks

    TaskService --> Actions

    DeviceService --> Devices

    AgentManager --> Transport

    AuditService --> AuditLogs

    Actions --> Executions

    Actions --> Verifications

    Users --> Tasks

```

## Component Responsibilities

### Android UI

Responsible for user interaction.

Main screens:

- Task screen.
- Approval screen.
- History screen.
- Device screen.

### REST API

Responsible for standard HTTP requests.

Examples:

- Create task.
- Get task.
- Get devices.
- Approve task.
- Cancel task.
- Get task history.

### WebSocket Server

Responsible for real-time communication.

Used for:

- task progress
- execution updates
- verification updates
- desktop agent connection
- desktop agent results

### Task Service

Responsible for task lifecycle.

Responsibilities:

- create task
- update task
- call planner
- trigger guardrail
- coordinate execution
- update task state
- trigger verification

### Device Service

Responsible for desktop device management.

Responsibilities:

- register device
- pair device
- track online/offline status
- track last seen time

### Planner Client

Responsible for communicating with the Python planner.

The planner client must validate the returned structured data before passing it to the next stage.

### Guardrail Service

Responsible for deterministic safety decisions.

It must not depend on an LLM for final authorization.

### Agent Manager

Responsible for connected desktop agents.

Responsibilities:

- track connections
- identify target device
- send approved actions
- receive execution results

### Audit Service

Responsible for security and lifecycle logging.

### Planner Service

Responsible for converting natural language into structured actions.

It proposes actions but does not authorize them.

### Desktop Agent

Responsible for actual desktop execution.

### Local Guard

Performs another safety validation immediately before execution.

### Action Executor

Maps supported action types to controlled implementations.

Example:

create_directory
open_application
get_system_info
move_file

### System Interface

Provides controlled access to Linux system functionality.

It should avoid unrestricted shell execution.

### Verification Engine

Checks the resulting desktop state.

Examples:

- process exists
- file exists
- file metadata matches
- expected exit code
- expected output
- system state changed as expected

### PostgreSQL

Stores persistent system data and audit information.

## Architecture Principle

Components must have clear responsibilities.

No component should bypass the safety layer.

The expected execution chain is:

User
-> Android
-> Backend
-> Planner
-> Guardrail
-> Approval if required
-> Desktop Agent
-> Executor
-> Verification
-> Backend
-> Android
```
