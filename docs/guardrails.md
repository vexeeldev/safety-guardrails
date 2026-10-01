# Safety Guardrails - Safety Policy Specification

## 1. Purpose

This document defines the safety mechanism that controls whether an AI-generated desktop action can be executed.

The guardrail is a mandatory security layer.

Planner output must never be executed directly.

Core rule:

AI proposes.
Guardrail decides.
Agent executes.

---

## 2. Risk Levels

There are three risk levels.

### SAFE

Action is considered low risk and can execute automatically.

Examples:

- get system information
- get process information
- read file metadata
- create temporary directory inside allowed workspace
- create temporary file inside allowed workspace

---

### REVIEW

Action may be legitimate but has meaningful system impact.

Explicit user approval is required.

Examples:

- shutdown
- sleep
- lock desktop
- move files
- close applications
- execute predefined project commands
- modify configuration inside allowed workspace

---

### BLOCKED

Action is prohibited.

It must not be executed even if the planner requests it.

Examples:

- delete operating system files
- access protected credentials
- modify security configuration without an explicit supported workflow
- unrestricted shell execution
- destructive recursive deletion
- modify files outside permitted scopes
- disable safety mechanisms
- bypass guardrail validation

---

## 3. Core Guardrail Pipeline

Every action must pass through:

1. Action type validation.
2. Parameter validation.
3. Target validation.
4. Scope validation.
5. Protected resource validation.
6. Risk classification.
7. Policy decision.
8. Approval requirement.
9. Final execution authorization.

---

## 4. Action Type Allowlist

The system must maintain an explicit allowlist.

Example:

ALLOWED_ACTIONS:

- get_system_info
- get_process_info
- open_application
- close_application
- create_directory
- create_file
- read_file_metadata
- move_file
- lock_desktop
- sleep
- shutdown
- run_project_command

Unknown action types must be rejected.

Do not use a generic fallback such as:

execute(action)

where action can contain arbitrary shell code.

---

## 5. Protected Resources

Protected resources include:

- /etc
- /boot
- /usr
- /bin
- /sbin
- /lib
- /lib64
- /var
- SSH private keys
- credential files
- browser credential databases
- system authentication files
- arbitrary filesystem locations not explicitly allowed

The exact protected-resource list may be expanded during implementation.

---

## 6. Allowed Workspace

The system should define an allowed workspace.

Example:

~/SafetyGuardrails/workspace

or another explicitly configured project workspace.

File manipulation actions should initially be restricted to this workspace.

The system should not assume that every path supplied by the planner is safe.

---

## 7. Path Validation

Paths must be normalized before validation.

The system must prevent:

../

path traversal.

Example unsafe input:

~/SafetyGuardrails/workspace/../../etc/passwd

The normalized path must be checked against the allowed workspace.

Symlink behavior must also be considered.

A path that appears safe but resolves to a protected location must be rejected.

---

## 8. Shell Execution Policy

Arbitrary shell execution is NOT allowed in MVP.

Do not implement:

run_shell(command string)

as a general-purpose execution interface.

Instead use explicit predefined actions.

Example:

run_project_tests

is allowed if implemented as a controlled executor.

Example:

run_command("rm -rf /")

must never be accepted.

---

## 9. Approval Policy

REVIEW actions require explicit user approval.

Approval must contain:

- task ID
- action ID
- approval timestamp
- user/device identifier
- approval decision

Approval is valid only for the exact action being approved.

Approval must not disable guardrails.

---

## 10. BLOCKED Policy

BLOCKED actions must immediately stop the task or action.

The system should return:

- blocked status
- policy identifier
- human-readable reason

Example:

status:
BLOCKED

reason:
"Action attempts to access a protected system path."

---

## 11. Fail-Closed Behavior

If guardrail validation fails because of:

- invalid policy
- missing configuration
- unknown action
- malformed target
- missing safety information
- internal validation error

the action must NOT execute.

Default behavior:

deny execution.

---

## 12. Defense in Depth

Guardrails must exist in multiple layers.

Layer 1:
Planner structured output validation.

Layer 2:
Backend validation.

Layer 3:
Guardrail policy engine.

Layer 4:
Desktop agent validation.

Layer 5:
Executor-specific validation.

A failure in one layer must not automatically bypass the others.

---

## 13. Audit Logging

The system must record:

- original user instruction
- planner output
- guardrail decision
- risk level
- reason
- approval decision
- execution result
- verification result

This allows later analysis of safety effectiveness.

---

## 14. Policy Examples

Example 1:

Action:
get_system_info

Decision:
SAFE

Reason:
Read-only system information.

---

Example 2:

Action:
create_directory
Target:
~/SafetyGuardrails/workspace/test

Decision:
SAFE

Reason:
Target is inside allowed workspace.

---

Example 3:

Action:
move_file
Target:
~/SafetyGuardrails/workspace/a.txt
Destination:
~/SafetyGuardrails/workspace/b.txt

Decision:
REVIEW

Reason:
File state will be modified.

---

Example 4:

Action:
shutdown

Decision:
REVIEW

Reason:
System state will be changed.

---

Example 5:

Action:
delete_file
Target:
/etc/passwd

Decision:
BLOCKED

Reason:
Protected system resource.

---

## 15. Important Implementation Rule

Do not make the LLM responsible for safety decisions.

The LLM may provide reasoning or proposed actions.

The deterministic policy engine must make the final safety decision.

---

## 16. Research Relevance

The guardrail system is one of the main research components.

Evaluation should measure:

- dangerous actions blocked
- safe actions incorrectly blocked
- review actions correctly identified
- false positives
- false negatives
- policy decision latency

Possible metrics:

Precision
Recall
False Positive Rate
False Negative Rate
Decision Latency

---

## 17. Future Improvements

Potential improvements:

- policy configuration
- role-based policies
- user-defined workspace
- application-specific policies
- risk scoring
- context-aware policy
- dynamic approval
- policy versioning
