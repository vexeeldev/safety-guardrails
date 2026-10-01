## `sequence.md`

```text
# Safety Guardrails - Sequence Diagram

## Purpose

This document describes communication between the major system components during task execution.

## Normal Task Execution


```mermaid
sequenceDiagram

    actor User

    participant Android

    participant Backend

    participant Planner

    participant Guardrail

    participant Agent

    participant Desktop

    participant Verifier

    participant Database

    User->>Android: Enter task

    Android->>Backend: Create task

    Backend->>Database: Store task

    Backend->>Planner: Send instruction

    Planner-->>Backend: Return structured actions

    Backend->>Database: Store actions

    Backend->>Guardrail: Validate actions

    Guardrail-->>Backend: Return safety decision

    alt SAFE

        Backend->>Database: Store SAFE decision

        Backend->>Agent: Send approved action

    else REVIEW

        Backend->>Android: Request user approval

        Android->>User: Display approval request

        User->>Android: Approve action

        Android->>Backend: Send approval

        Backend->>Database: Store approval

        Backend->>Agent: Send approved action

    else BLOCKED

        Backend->>Database: Store BLOCKED decision

        Backend-->>Android: Task blocked

    end

    Agent->>Agent: Local action validation

    Agent->>Desktop: Execute action

    Desktop-->>Agent: Execution result

    Agent-->>Backend: Execution result

    Backend->>Database: Store execution result

    Backend->>Verifier: Verify result

    Verifier->>Desktop: Check resulting state

    Desktop-->>Verifier: Current system state

    Verifier-->>Backend: Verification result

    Backend->>Database: Store verification

    Backend-->>Android: Final task result

    Android-->>User: Display final result
```

## REVIEW Flow

```mermaid
sequenceDiagram

    actor User

    participant Android

    participant Backend

    participant Guardrail

    participant Agent

    Android->>Backend: Create task

    Backend->>Guardrail: Validate action

    Guardrail-->>Backend: REVIEW

    Backend-->>Android: Approval required

    Android->>User: Display action and reason

    User->>Android: Approve

    Android->>Backend: Submit approval

    Backend->>Agent: Execute approved action

    Agent-->>Backend: Execution result

    Backend-->>Android: Return result
```

## BLOCKED Flow

```mermaid
sequenceDiagram

    participant Android

    participant Backend

    participant Planner

    participant Guardrail

    Android->>Backend: Create task

    Backend->>Planner: Plan task

    Planner-->>Backend: Structured action

    Backend->>Guardrail: Validate action

    Guardrail-->>Backend: BLOCKED

    Backend-->>Android: Task blocked
```

## Important Communication Rules

1. Planner must never communicate directly with Desktop Agent.
2. Backend coordinates communication between services.
3. Guardrail validation must happen before execution.
4. REVIEW actions require explicit approval.
5. BLOCKED actions must stop before execution.
6. Execution results must be persisted.
7. Verification should happen after execution.
8. Final task status must be based on execution and verification results.
```
