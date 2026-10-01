# Safety Guardrails - Architecture Specification

## 1. Project Overview

Project name:
Safety Guardrails

Project title:
Penerapan Mekanisme Safety Guardrails dan Verifikasi Otomatis pada Asisten Cerdas untuk Eksekusi Tugas Desktop Jarak Jauh

Purpose:
Safety Guardrails adalah sistem yang memungkinkan pengguna memberikan instruksi tugas desktop melalui aplikasi Android. Instruksi diproses menjadi rencana tindakan, diperiksa oleh mekanisme safety guardrails, kemudian dikirim ke desktop agent untuk dieksekusi.

Sistem harus memastikan bahwa AI tidak memiliki kendali langsung untuk mengeksekusi tindakan desktop tanpa melewati policy dan safety validation.

Core principle:

AI plans.
Policy decides.
Agent executes.
Verifier confirms.

AI tidak boleh langsung menjalankan command desktop.

---

## 2. System Goals

Sistem harus mampu:

1. Menerima instruksi natural language dari Android.
2. Mengubah instruksi menjadi structured actions.
3. Menilai setiap action berdasarkan safety policy.
4. Mengizinkan action yang aman.
5. Meminta persetujuan pengguna untuk action berisiko.
6. Menolak action yang diblokir.
7. Mengirim action yang telah disetujui ke desktop agent.
8. Mengeksekusi action pada desktop.
9. Memverifikasi apakah action benar-benar berhasil.
10. Mengirim hasil eksekusi dan verifikasi kembali ke Android.
11. Menyimpan task, action, execution, verification, dan audit information.

---

## 3. High-Level Architecture

Architecture:

Android Client
    |
    | REST / WebSocket
    v
Go Backend
    |
    +--------------------+
    |                    |
    v                    v
Planner Service      PostgreSQL
    |
    v
Structured Actions
    |
    v
Safety Guardrails
    |
    +---- SAFE ------> Execute
    |
    +---- REVIEW ----> User Approval
    |
    +---- BLOCKED ---> Reject
    |
    v
Go Desktop Agent
    |
    v
Linux Desktop
    |
    v
Verification Engine
    |
    v
Go Backend
    |
    v
Android Client

---

## 4. Main Components

### 4.1 Android Client

Technology:
- Kotlin
- Jetpack Compose

Responsibilities:
- Login or device identification.
- Pair with desktop agent.
- Create task.
- Send natural language instruction.
- Display task status.
- Display planned actions.
- Display safety classification.
- Request user approval when necessary.
- Display execution progress.
- Display verification result.
- Display task history.

Android must never execute desktop commands directly.

---

### 4.2 Go Backend

Technology:
- Go
- Gin
- WebSocket
- PostgreSQL

Responsibilities:
- API server.
- Authentication.
- Device management.
- Task management.
- Communication between Android, planner, and desktop agent.
- Persist task state.
- Persist audit information.
- Route approved actions to desktop agent.
- Receive execution results.
- Receive verification results.

The backend is the central coordinator.

The backend must not blindly execute arbitrary commands.

---

### 4.3 Planner Service

Technology:
- Python
- FastAPI
- Pydantic
- LLM API

Responsibilities:
- Convert natural language instructions into structured actions.
- Validate structured output.
- Provide action descriptions.
- Estimate required parameters.

Example input:

"buat folder test di desktop"

Example structured output:

{
  "intent": "create_directory",
  "target": "~/Desktop/test"
}

The planner only proposes actions.

The planner does not decide whether an action is safe.

---

### 4.4 Safety Guardrails

Responsibilities:
- Inspect every proposed action.
- Determine risk level.
- Validate target.
- Validate parameters.
- Apply policy.
- Prevent dangerous actions.
- Require user approval where necessary.

Risk levels:

SAFE
REVIEW
BLOCKED

The guardrail layer is authoritative over the planner.

If the planner proposes an unsafe action, the guardrail must be able to reject it.

---

### 4.5 Desktop Agent

Technology:
- Go

Target platform:
- Linux

Responsibilities:
- Maintain connection with backend.
- Receive approved structured actions.
- Validate received actions again.
- Execute only supported actions.
- Return execution result.
- Trigger verification.
- Return verification result.

The desktop agent must never execute arbitrary shell commands by default.

Actions should use explicit executors.

Example:

Instead of:

shell("rm -rf /some/path")

Use:

DeleteFile(path)

where DeleteFile has its own safety validation.

---

### 4.6 Verification Engine

Responsibilities:
- Determine whether an executed action actually succeeded.
- Verify the resulting state rather than relying only on exit code.

Examples:

Open application:
- Verify process exists.

Create file:
- Verify file exists.
- Verify expected file type.

Move file:
- Verify source no longer exists.
- Verify destination exists.

Run tests:
- Verify process exit code.
- Optionally inspect expected output.

Verification result:

VERIFIED
FAILED
UNCERTAIN

---

### 4.7 PostgreSQL

Stores:
- users
- devices
- tasks
- actions
- executions
- verifications
- audit information

The database must preserve enough information to reconstruct what happened during a task.

---

## 5. Communication

Android <-> Backend:
- REST API
- WebSocket for real-time task updates

Backend <-> Planner:
- HTTP API

Backend <-> Desktop Agent:
- WebSocket

Desktop Agent -> Backend:
- execution result
- verification result
- agent status

---

## 6. Security Principles

1. Never trust planner output directly.
2. Never execute arbitrary shell commands by default.
3. Every action must have a known action type.
4. Every action must pass guardrail validation.
5. REVIEW actions require explicit user approval.
6. BLOCKED actions cannot be approved through the normal UI.
7. Desktop agent validates received actions again.
8. All important actions must be auditable.
9. Failed verification must not be reported as successful.
10. System should fail closed when safety information is unavailable.

---

## 7. Design Principle

The architecture must maintain separation of responsibility:

Planner:
"What should be done?"

Guardrail:
"Is this allowed?"

User:
"Do I approve this risky action?"

Agent:
"Execute the approved action."

Verifier:
"Did it actually happen?"

Backend:
"Coordinate and record everything."

---

## 8. MVP Scope

Initial supported actions:

1. Open application.
2. Close application.
3. Get system information.
4. Get process information.
5. Create directory inside allowed workspace.
6. Create temporary file inside allowed workspace.
7. Read file metadata.
8. Move file inside allowed workspace.
9. Run limited predefined project commands.
10. Lock desktop.
11. Sleep.
12. Shutdown with explicit approval.

The MVP must not support unrestricted shell execution.

---

## 9. Future Expansion

Potential future features:

- More desktop applications.
- Advanced file operations.
- Clipboard operations.
- Screenshot verification.
- OCR verification.
- GUI automation.
- Multi-device management.
- More sophisticated policy engine.
- Policy customization.
- Task scheduling.

Future features must not bypass the core safety architecture.
