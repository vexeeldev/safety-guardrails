# Safety Guardrails - System Flow Specification

## 1. Purpose

This document defines the complete lifecycle of a task from Android input until final verification.

The implementation must follow this flow unless a later technical decision explicitly changes it.

Core flow:

User
    ->
Android
    ->
Backend
    ->
Planner
    ->
Structured Actions
    ->
Safety Guardrails
    ->
Approval if required
    ->
Desktop Agent
    ->
Execution
    ->
Verification
    ->
Backend
    ->
Android

---

## 2. Task Lifecycle

Task states:

CREATED
PLANNING
GUARDRAIL_CHECK
WAITING_APPROVAL
APPROVED
EXECUTING
VERIFYING
COMPLETED
FAILED
BLOCKED
CANCELLED

State transitions:

CREATED
    -> PLANNING

PLANNING
    -> GUARDRAIL_CHECK

GUARDRAIL_CHECK
    -> APPROVED
    -> WAITING_APPROVAL
    -> BLOCKED

WAITING_APPROVAL
    -> APPROVED
    -> CANCELLED

APPROVED
    -> EXECUTING

EXECUTING
    -> VERIFYING
    -> FAILED

VERIFYING
    -> COMPLETED
    -> FAILED
    -> UNCERTAIN

---

## 3. Create Task

User enters natural language instruction.

Example:

"buat folder test di desktop"

Android sends:

POST /api/v1/tasks

Payload contains:

- instruction
- device_id

Backend creates task.

Initial status:

CREATED

---

## 4. Planning

Backend sends instruction to Planner.

Planner converts natural language into structured action.

Example:

Input:

"buat folder test di desktop"

Output:

{
  "actions": [
    {
      "type": "create_directory",
      "target": "~/Desktop/test"
    }
  ]
}

Planner output must use known action types.

Unknown action types must be rejected.

---

## 5. Guardrail Check

Every planned action is inspected.

Guardrail checks:

1. Is action type supported?
2. Is target allowed?
3. Is target protected?
4. Are parameters valid?
5. Does action modify the system?
6. Does action require user approval?
7. Is the action explicitly blocked?

Each action receives:

SAFE
REVIEW
BLOCKED

---

## 6. SAFE Action

SAFE action can proceed automatically.

Example:

get_system_info

Flow:

Planner
    ->
Guardrail
    ->
SAFE
    ->
Backend
    ->
Desktop Agent
    ->
Execute
    ->
Verify

No user approval is required.

---

## 7. REVIEW Action

REVIEW action requires explicit user approval.

Example:

shutdown

Flow:

Planner
    ->
Guardrail
    ->
REVIEW
    ->
Android shows approval dialog
    ->
User approves
    ->
Backend
    ->
Desktop Agent
    ->
Execute
    ->
Verify

The system must not execute REVIEW actions before approval.

---

## 8. BLOCKED Action

BLOCKED action cannot execute.

Example:

delete_system_files

Flow:

Planner
    ->
Guardrail
    ->
BLOCKED
    ->
Task marked BLOCKED
    ->
Android displays reason

The normal approval interface must not allow bypassing a BLOCKED policy.

---

## 9. Execution

After approval or automatic authorization, backend sends structured action to desktop agent.

Desktop agent:

1. Receives action.
2. Validates action.
3. Checks target.
4. Executes supported executor.
5. Captures result.
6. Sends result to backend.

Example result:

{
  "success": true,
  "message": "Directory created successfully"
}

---

## 10. Verification

Execution success does not automatically mean task success.

Verification must inspect the resulting state.

Example:

Action:
create_directory("~/Desktop/test")

Verification:

Check whether:
~/Desktop/test

exists.

If exists:

VERIFIED

If does not exist:

FAILED

---

## 11. Final Result

If execution succeeds and verification succeeds:

Task:
COMPLETED

If execution fails:

Task:
FAILED

If verification cannot determine the result:

Task:
UNCERTAIN

The system must not report success when verification failed.

---

## 12. Error Flow

Possible errors:

Planner error:
- Task FAILED.

Invalid action:
- Task BLOCKED.

Guardrail error:
- Fail closed.
- Do not execute action.

Agent unavailable:
- Task remains pending or becomes FAILED depending on timeout policy.

Execution error:
- Task FAILED.

Verification error:
- Task UNCERTAIN or FAILED.

Network timeout:
- Do not assume execution failed unless the system has enough evidence.
- Avoid duplicate execution of non-idempotent actions.

---

## 13. Approval Rules

Approval must be:

- explicit
- user initiated
- associated with a specific task
- associated with specific actions
- recorded in audit information

Approval must not automatically approve future unrelated tasks.

---

## 14. Idempotency

Actions that can accidentally execute twice must have protection against duplicate execution.

Examples:

create_directory:
- safe if already exists and expected state matches.

shutdown:
- must not be automatically retried.

move_file:
- must check current state before retrying.

---

## 15. Audit Flow

The system should record:

- task creation
- planner result
- guardrail decision
- approval
- execution
- verification
- final status

This information is required for debugging and research evaluation.
