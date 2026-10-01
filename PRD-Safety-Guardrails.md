# PRD --- Safety Guardrails

## Context-Aware Desktop Task Automation with Safety Guardrails and Automatic Verification

**Document status:** Draft v1.0\
**Project type:** Proyek Akhir / Final Project\
**Target platform:** Android + Linux Desktop\
**Primary language:** Indonesian\
**Last updated:** 1 October 2026

------------------------------------------------------------------------

# 1. Ringkasan Proyek

## 1.1 Nama Produk

**Safety Guardrails**

Nama produk sementara untuk sistem yang memungkinkan pengguna memberikan
instruksi dari perangkat Android untuk menjalankan tugas pada
laptop/desktop secara jarak jauh, dengan mekanisme **planning, safety
guardrails, controlled execution, dan automatic verification**.

## 1.2 Judul PA Sementara

> **Penerapan Mekanisme Safety Guardrails dan Verifikasi Otomatis pada
> Asisten Cerdas untuk Eksekusi Tugas Desktop Jarak Jauh**

Judul ini masih dapat disesuaikan setelah metode, ruang lingkup, dan
hasil eksperimen final ditentukan.

## 1.3 Elevator Pitch

Safety Guardrails adalah sistem Android-to-desktop yang memungkinkan
pengguna memberikan instruksi seperti:

> "Buka VS Code dan buka project X."

Sistem tidak langsung menjalankan instruksi tersebut. Instruksi terlebih
dahulu dianalisis dan diubah menjadi rencana aksi, kemudian melewati
**Safety Guardrails** untuk menentukan apakah aksi aman, membutuhkan
persetujuan pengguna, atau harus diblokir.

Setelah aksi dijalankan oleh **Desktop Agent**, sistem melakukan
**verifikasi otomatis** untuk memastikan hasil yang diminta benar-benar
terjadi.

Alur utama:

``` text
User
  ↓
Android App
  ↓
Backend
  ↓
Planner / LLM
  ↓
Safety Guardrails
  ↓
User Approval (jika diperlukan)
  ↓
Desktop Agent
  ↓
Execution
  ↓
Verification
  ↓
Result
  ↓
Android App
```

------------------------------------------------------------------------

# 2. Latar Belakang

Remote control desktop biasanya memberikan kemampuan eksekusi langsung,
sedangkan sistem berbasis AI dapat menghasilkan tindakan berdasarkan
instruksi natural language.

Masalah muncul ketika AI diberi kemampuan untuk melakukan tindakan nyata
pada komputer.

Instruksi seperti:

-   membuka aplikasi,
-   membuat folder,
-   menjalankan command,
-   mengubah file,
-   menghapus file,
-   mematikan komputer,

memiliki tingkat risiko yang berbeda.

Sistem yang hanya menerjemahkan instruksi menjadi command tidak cukup
karena model AI dapat menghasilkan aksi yang tidak diinginkan, terlalu
luas, atau memiliki efek samping.

Selain itu, keberhasilan eksekusi tidak selalu dapat diasumsikan hanya
karena command berhasil dijalankan.

Contoh:

``` text
User:
"Buka VS Code"

Agent:
menjalankan command

Command:
exit code = 0
```

Exit code `0` belum tentu membuktikan bahwa VS Code benar-benar terbuka
dan dapat digunakan.

Karena itu sistem ini menggabungkan empat komponen utama:

1.  **Task Planning**
2.  **Safety Guardrails**
3.  **Controlled Desktop Execution**
4.  **Automatic Verification**

------------------------------------------------------------------------

# 3. Problem Statement

Sistem otomatisasi desktop berbasis natural language menghadapi beberapa
masalah:

1.  Instruksi pengguna dapat ambigu.
2.  LLM dapat menghasilkan aksi yang tidak sesuai dengan intent
    pengguna.
3.  Tidak semua aksi memiliki tingkat risiko yang sama.
4.  Aksi berbahaya dapat dijalankan jika tidak ada policy enforcement.
5.  Keberhasilan command tidak selalu berarti tujuan pengguna tercapai.
6.  Pengguna membutuhkan cara untuk mengetahui apa yang akan dilakukan
    sistem sebelum eksekusi.
7.  Sistem membutuhkan mekanisme audit agar tindakan dapat ditelusuri.

------------------------------------------------------------------------

# 4. Tujuan Proyek

## 4.1 Tujuan Utama

Membangun prototype sistem yang mampu:

-   menerima instruksi natural language dari Android,
-   mengubah instruksi menjadi rencana aksi terstruktur,
-   mengklasifikasikan risiko setiap aksi,
-   menerapkan policy sebelum eksekusi,
-   meminta persetujuan pengguna untuk aksi tertentu,
-   mengeksekusi aksi melalui desktop agent,
-   memverifikasi hasil eksekusi,
-   mengirim status dan hasil kembali ke Android.

## 4.2 Tujuan Penelitian

Mengevaluasi apakah penerapan safety guardrails dan automatic
verification dapat:

-   mengurangi eksekusi aksi yang tidak diizinkan,
-   membedakan aksi berdasarkan tingkat risiko,
-   mencegah aksi berbahaya tertentu,
-   meningkatkan keandalan laporan keberhasilan task,
-   memberikan audit trail yang jelas.

------------------------------------------------------------------------

# 5. Target Pengguna

## Primary User

Pengguna komputer yang ingin mengendalikan atau mengotomatisasi tugas
desktop dari Android.

## Secondary User

-   developer,
-   mahasiswa,
-   pengguna Linux,
-   pengguna yang sering berpindah antara HP dan laptop,
-   pengguna yang membutuhkan remote task execution sederhana.

------------------------------------------------------------------------

# 6. Scope

## 6.1 In Scope

Prototype akan berfokus pada desktop Linux terlebih dahulu.

Aksi yang didukung pada MVP:

### Application

-   membuka aplikasi,
-   menutup aplikasi,
-   memeriksa apakah aplikasi berjalan.

### File System

-   membuat folder,
-   membuat file teks,
-   membaca metadata file,
-   memindahkan file pada area yang diizinkan.

### System

-   melihat CPU/RAM,
-   melihat disk usage,
-   lock screen,
-   sleep,
-   shutdown dengan approval eksplisit.

### Command

Command execution dibatasi oleh allowlist/policy.

Contoh command yang diperbolehkan:

``` text
git status
go test ./...
npm run build
```

Command arbitrary shell tidak langsung diizinkan pada MVP.

### Verification

Sistem melakukan verifikasi terhadap hasil task menggunakan:

-   process state,
-   file existence,
-   file metadata,
-   command exit status,
-   output command,
-   system state,
-   rule-based checks.

------------------------------------------------------------------------

# 7. Out of Scope

Untuk menjaga PA tetap realistis, versi awal tidak mencakup:

-   kontrol semua OS secara bersamaan,
-   akses root,
-   eksekusi command arbitrary tanpa policy,
-   penghapusan massal,
-   modifikasi system-critical files,
-   eksploitasi keamanan,
-   bypass permission,
-   persistence tersembunyi,
-   remote access tanpa autentikasi,
-   autonomous agent tanpa batasan,
-   kontrol perangkat pihak ketiga tanpa izin.

Windows dan macOS dapat menjadi rencana pengembangan lanjutan, bukan
target MVP.

------------------------------------------------------------------------

# 8. Konsep Sistem

Sistem terdiri dari lima bagian utama.

``` text
┌─────────────────────┐
│    Android Client   │
│ Kotlin + Compose    │
└──────────┬──────────┘
           │
           │ HTTPS / WebSocket
           ↓
┌─────────────────────┐
│     Go Backend      │
│ API + Session       │
└──────────┬──────────┘
           │
           ↓
┌─────────────────────┐
│ Planner / LLM       │
│ Python + FastAPI    │
└──────────┬──────────┘
           │
           ↓
┌─────────────────────┐
│ Safety Guardrails   │
│ Policy + Risk       │
└──────────┬──────────┘
           │
           ↓
┌─────────────────────┐
│   Desktop Agent     │
│        Go           │
└──────────┬──────────┘
           │
           ↓
┌─────────────────────┐
│ Linux Desktop       │
└─────────────────────┘
```

------------------------------------------------------------------------

# 9. Arsitektur Komponen

## 9.1 Android Client

Teknologi:

-   Kotlin
-   Jetpack Compose
-   Android SDK
-   HTTP client
-   WebSocket client

Tanggung jawab:

-   login/pairing,
-   memasukkan instruksi,
-   menampilkan plan,
-   meminta approval,
-   menampilkan progress,
-   menampilkan hasil,
-   melihat execution history.

------------------------------------------------------------------------

## 9.2 Go Backend

Teknologi:

-   Go
-   Gin
-   REST API
-   WebSocket
-   PostgreSQL

Tanggung jawab:

-   authentication/session,
-   device pairing,
-   menerima task,
-   menghubungkan Android dengan planner,
-   mengirim task ke desktop agent,
-   menyimpan execution history,
-   mengirim event realtime.

------------------------------------------------------------------------

## 9.3 Planner Service

Teknologi:

-   Python
-   FastAPI
-   LLM API

Tanggung jawab:

-   memahami natural language,
-   mengubah intent menjadi structured task plan,
-   tidak melakukan eksekusi langsung.

Contoh input:

``` text
"Buka VS Code lalu jalankan go test di project backend."
```

Output terstruktur:

``` json
{
  "intent": "run_project_test",
  "steps": [
    {
      "action": "open_application",
      "target": "code"
    },
    {
      "action": "run_command",
      "command": "go test ./...",
      "working_directory": "/project/backend"
    }
  ]
}
```

Planner tidak mempunyai akses langsung untuk menjalankan command di
desktop.

------------------------------------------------------------------------

# 10. Safety Guardrails

Safety Guardrails adalah komponen yang memutuskan apakah action dapat
dieksekusi.

## 10.1 Risk Levels

### SAFE

Aksi dengan risiko rendah.

Contoh:

``` text
get_cpu_usage
get_memory_usage
check_process
create_temp_folder
open_known_application
```

Aksi dapat dieksekusi tanpa approval tambahan.

### REVIEW

Aksi memiliki dampak tetapi masih dapat dilakukan dengan persetujuan.

Contoh:

``` text
move_file
install_package
run_project_command
change_application_config
```

Sistem menampilkan:

``` text
Action:
Run "go test ./..."

Risk:
REVIEW

Reason:
Command execution can modify project state.

[Approve] [Reject]
```

### BLOCKED

Aksi yang tidak diperbolehkan oleh policy.

Contoh:

``` text
rm -rf /
format disk
disable security controls
modify protected system files
execute unknown downloaded binary
```

Sistem menolak sebelum dikirim ke desktop agent.

------------------------------------------------------------------------

# 11. Policy Engine

Policy engine menggunakan pendekatan rule-based.

Contoh:

``` text
IF action.type == "shutdown"
THEN risk = REVIEW

IF action.type == "delete"
AND target == protected_path
THEN risk = BLOCKED

IF command contains dangerous_pattern
THEN risk = BLOCKED

IF command not in allowlist
THEN risk = REVIEW
```

Policy tidak bergantung sepenuhnya pada LLM.

LLM menghasilkan rencana.

Policy engine menentukan apakah rencana tersebut boleh dijalankan.

------------------------------------------------------------------------

# 12. Prinsip Keamanan Utama

## 12.1 Least Privilege

Desktop agent berjalan dengan permission minimum yang diperlukan.

## 12.2 Explicit Approval

Aksi berisiko memerlukan persetujuan pengguna.

## 12.3 Separation of Planning and Execution

Planner tidak memiliki akses langsung ke desktop.

## 12.4 Allowlist

Action tertentu hanya boleh dilakukan jika masuk daftar yang diizinkan.

## 12.5 Auditability

Setiap task disimpan:

``` text
timestamp
user
instruction
generated plan
risk classification
approval
execution result
verification result
```

------------------------------------------------------------------------

# 13. Desktop Agent

Desktop Agent adalah service yang berjalan di laptop.

Tanggung jawab:

1.  menerima action,
2.  memvalidasi ulang action,
3.  memeriksa policy lokal,
4.  menjalankan action,
5.  mengumpulkan hasil,
6.  melakukan verification,
7.  mengirim hasil.

Agent tidak boleh mempercayai planner atau backend secara buta.

Flow:

``` text
Receive Action
      ↓
Validate Schema
      ↓
Local Policy Check
      ↓
Execute
      ↓
Verify
      ↓
Return Result
```

------------------------------------------------------------------------

# 14. Verification Engine

Verification merupakan salah satu komponen penelitian utama.

## Contoh 1: Open Application

Request:

``` text
open VS Code
```

Verification:

``` text
process "code" exists?
```

Result:

``` text
PASS
```

## Contoh 2: Create File

Request:

``` text
create /tmp/test.txt
```

Verification:

``` text
file exists?
file type correct?
```

## Contoh 3: Run Test

Request:

``` text
go test ./...
```

Verification:

``` text
exit code == 0
output indicates test success
```

## Contoh 4: Shutdown

Verification dapat berupa:

``` text
shutdown command accepted
```

Actual machine shutdown tidak perlu digunakan dalam automated testing
karena akan membuat peneliti ikut kehilangan laptop.

------------------------------------------------------------------------

# 15. Android UI

MVP terdiri dari beberapa layar.

## 15.1 Home

Menampilkan:

``` text
Laptop:
Connected

CPU:
24%

RAM:
7.2 / 15 GB

Agent:
Online
```

## 15.2 Task Input

``` text
What do you want to do?

[ Buka VS Code dan jalankan test ]

        [Execute]
```

## 15.3 Plan Preview

``` text
Planned Actions

1. Open VS Code
2. Run go test ./...

Risk: REVIEW

        [Approve]
        [Cancel]
```

## 15.4 Execution

``` text
Executing...

✓ Open VS Code
✓ Run command

Verifying...
```

## 15.5 Result

``` text
Task completed

Result:
go test ./... passed

Verification:
PASS
```

## 15.6 History

Menampilkan:

``` text
10:42  Run tests       SUCCESS
10:38  Open VS Code    SUCCESS
10:21  Shutdown        BLOCKED
```

------------------------------------------------------------------------

# 16. Backend API

Contoh endpoint:

``` text
POST /api/auth/pair
POST /api/tasks
GET  /api/tasks/:id
POST /api/tasks/:id/approve
POST /api/tasks/:id/cancel
GET  /api/tasks
GET  /api/devices
GET  /api/devices/:id/status
WS   /api/events
```

Contoh request:

``` json
{
  "device_id": "desktop-01",
  "instruction": "Buka VS Code"
}
```

Response:

``` json
{
  "task_id": "task-123",
  "status": "pending_review"
}
```

------------------------------------------------------------------------

# 17. Data Model

Minimal database:

## users

``` text
id
name
created_at
```

## devices

``` text
id
user_id
name
platform
status
last_seen
```

## tasks

``` text
id
user_id
device_id
instruction
status
risk_level
created_at
completed_at
```

## actions

``` text
id
task_id
sequence
action_type
payload
risk_level
status
```

## executions

``` text
id
action_id
started_at
finished_at
exit_code
output
error
```

## verifications

``` text
id
execution_id
method
status
details
created_at
```

------------------------------------------------------------------------

# 18. Task State Machine

Task state:

``` text
CREATED
   ↓
PLANNING
   ↓
PLANNED
   ↓
GUARDRAIL_CHECK
   ↓
 ┌───────────────┐
 ↓               ↓
BLOCKED       REVIEW
                 ↓
              APPROVED
                 ↓
             EXECUTING
                 ↓
             VERIFYING
              ↓     ↓
          SUCCESS   FAILED
```

------------------------------------------------------------------------

# 19. Communication

Untuk MVP:

``` text
Android
  ↕
Go Backend
  ↕
Desktop Agent
```

Planner dipanggil oleh backend.

WebSocket digunakan untuk event realtime:

``` text
task.created
task.planned
task.review_required
task.approved
task.started
task.progress
task.verifying
task.completed
task.failed
```

------------------------------------------------------------------------

# 20. Authentication & Pairing

MVP menggunakan device pairing.

Contoh:

``` text
Desktop Agent
   ↓
Generate pairing code / QR
   ↓
Android scans
   ↓
Backend establishes relationship
```

Setiap device memiliki identity/token.

Token tidak ditampilkan sebagai plain text di UI setelah pairing
selesai.

------------------------------------------------------------------------

# 21. Non-Functional Requirements

## Performance

Target awal:

-   command planning response: \< 5 detik tanpa memperhitungkan latency
    provider LLM
-   local action dispatch: \< 500 ms
-   status update: near realtime
-   verification: \< 2 detik untuk action sederhana

Target tersebut merupakan target eksperimen, bukan jaminan.

## Reliability

Sistem harus dapat:

-   mendeteksi agent offline,
-   timeout execution,
-   retry komunikasi yang aman,
-   menyimpan task state,
-   mencegah duplicate execution.

## Security

-   authenticated communication,
-   encrypted transport,
-   least privilege,
-   policy enforcement,
-   audit logging,
-   no arbitrary shell by default.

------------------------------------------------------------------------

# 22. Technology Stack

## Android

``` text
Kotlin
Jetpack Compose
Android SDK
Ktor/OkHttp client
WebSocket
```

## Backend

``` text
Go
Gin
WebSocket
PostgreSQL
```

## Planner

``` text
Python
FastAPI
LLM API
Pydantic
```

## Desktop Agent

``` text
Go
OS process APIs
WebSocket client
```

## Infrastructure

``` text
Docker
Docker Compose
Git
```

------------------------------------------------------------------------

# 23. Repository Structure

``` text
safety-guardrails/
│
├── mobile/
│   ├── app/
│   ├── gradle/
│   └── ...
│
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── go.mod
│   └── ...
│
├── planner/
│   ├── app/
│   ├── tests/
│   ├── requirements.txt
│   └── ...
│
├── desktop-agent/
│   ├── cmd/
│   ├── internal/
│   ├── go.mod
│   └── ...
│
├── docs/
│   ├── architecture.md
│   ├── flow.md
│   ├── guardrails.md
│   ├── api.md
│   ├── research.md
│   └── diagrams/
│
├── docker-compose.yml
└── README.md
```

------------------------------------------------------------------------

# 24. Development Phases

## Phase 1 --- Project Foundation

-   initialize repository,
-   initialize Android,
-   initialize Go backend,
-   initialize planner,
-   initialize desktop agent,
-   documentation structure.

## Phase 2 --- Device Connection

-   desktop agent registration,
-   Android pairing,
-   device status,
-   heartbeat.

## Phase 3 --- Basic Remote Actions

Implement:

``` text
get_system_info
open_application
check_process
create_file
```

## Phase 4 --- Backend Task Pipeline

Implement:

``` text
instruction
→ task
→ action
→ execution
→ result
```

## Phase 5 --- Planner

Implement natural language → structured action.

LLM output must follow strict JSON schema.

## Phase 6 --- Safety Guardrails

Implement:

-   action classification,
-   allowlist,
-   blocked patterns,
-   approval workflow,
-   local agent validation.

## Phase 7 --- Verification

Implement action-specific verification.

## Phase 8 --- Android UX

Implement:

-   task input,
-   plan preview,
-   approval,
-   progress,
-   result,
-   history.

## Phase 9 --- Testing & Evaluation

Create controlled scenarios.

## Phase 10 --- Finalization

-   performance testing,
-   security testing,
-   documentation,
-   screenshots,
-   research results,
-   final report,
-   presentation/demo.

------------------------------------------------------------------------

# 25. Testing Strategy

Testing dilakukan pada tiga level.

## Unit Test

Untuk:

-   policy engine,
-   risk classifier,
-   action parser,
-   verification engine.

## Integration Test

Test:

``` text
Android
→ Backend
→ Agent
→ Desktop
→ Backend
→ Android
```

## Scenario Test

Contoh:

### Scenario A

``` text
"Buka VS Code"
```

Expected:

``` text
SAFE
→ execute
→ verify
→ SUCCESS
```

### Scenario B

``` text
"Jalankan go test ./..."
```

Expected:

``` text
REVIEW
→ user approval
→ execute
→ verify
```

### Scenario C

``` text
"Hapus seluruh filesystem"
```

Expected:

``` text
BLOCKED
→ no execution
```

### Scenario D

``` text
"Buat file test.txt"
```

Expected:

``` text
SAFE/REVIEW
→ create
→ verify file existence
```

------------------------------------------------------------------------

# 26. Metode Evaluasi Penelitian

## 26.1 Guardrail Effectiveness

Buat dataset action dengan kategori:

``` text
Safe
Review
Dangerous
```

Ukur:

-   true positive,
-   false positive,
-   false negative,
-   precision,
-   recall.

Tujuan eksperimen bukan sekadar membuat sistem "terlihat aman", tetapi
mengukur performanya.

## 26.2 Verification Accuracy

Bandingkan:

``` text
Command exit status
vs
Verification engine
```

Kasus:

``` text
command success
but task failed
```

dan:

``` text
command failed
but partial task succeeded
```

## 26.3 Execution Reliability

Ukur:

-   task success rate,
-   failure rate,
-   timeout rate,
-   duplicate execution rate.

## 26.4 Latency

Ukur:

``` text
Instruction received
→ plan generated
→ guardrail decision
→ execution
→ verification
→ result
```

------------------------------------------------------------------------

# 27. Research Questions

Contoh pertanyaan penelitian:

### RQ1

Seberapa efektif mekanisme rule-based safety guardrails dalam mencegah
eksekusi aksi desktop yang tidak diizinkan?

### RQ2

Seberapa akurat automatic verification dalam menentukan keberhasilan
task dibandingkan hanya menggunakan exit status command?

### RQ3

Bagaimana pengaruh penggunaan safety approval workflow terhadap latency
keseluruhan task execution?

### RQ4

Seberapa reliabel sistem dalam menjalankan task desktop dari Android
melalui jaringan?

------------------------------------------------------------------------

# 28. Success Criteria

Prototype dianggap berhasil jika:

-   Android dapat terhubung ke desktop agent.
-   User dapat mengirim instruksi.
-   Planner dapat menghasilkan structured action.
-   Guardrails dapat mengklasifikasikan action.
-   Action blocked tidak pernah dikirim untuk eksekusi.
-   Action review membutuhkan approval.
-   Desktop agent dapat menjalankan action yang diizinkan.
-   Verification dapat memeriksa hasil.
-   History task tersimpan.
-   Sistem dapat menangani agent disconnect.
-   Eksperimen menghasilkan metrik yang dapat dianalisis.

------------------------------------------------------------------------

# 29. Risiko Proyek

## Risiko 1 --- LLM menghasilkan JSON tidak valid

Mitigasi:

-   structured output,
-   schema validation,
-   retry,
-   fallback.

## Risiko 2 --- LLM menghasilkan action berbahaya

Mitigasi:

-   LLM bukan security authority,
-   rule engine,
-   local validation,
-   allowlist,
-   blocklist.

## Risiko 3 --- Desktop agent terlalu powerful

Mitigasi:

-   least privilege,
-   restricted working directories,
-   explicit action types,
-   no arbitrary shell pada MVP.

## Risiko 4 --- Networking terlalu kompleks

Mitigasi:

MVP menggunakan LAN terlebih dahulu.

Remote internet/Tailscale menjadi tahap berikutnya.

## Risiko 5 --- Scope terlalu besar

Mitigasi:

MVP hanya Linux + Android + limited actions.

------------------------------------------------------------------------

# 30. MVP Definition

MVP harus mampu melakukan:

``` text
Android
   ↓
"Buka VS Code"
   ↓
Planner
   ↓
Structured Action
   ↓
Guardrails
   ↓
SAFE
   ↓
Desktop Agent
   ↓
Open VS Code
   ↓
Verify Process
   ↓
SUCCESS
   ↓
Android
```

Kemudian satu contoh action berisiko:

``` text
"Matikan laptop"
   ↓
Planner
   ↓
Guardrails
   ↓
REVIEW
   ↓
Android asks approval
   ↓
User approves
   ↓
Execute
```

Dan satu contoh blocked:

``` text
"Delete protected system directory"
   ↓
Guardrails
   ↓
BLOCKED
   ↓
No execution
```

Jika tiga alur tersebut sudah berjalan, inti PA sudah terbukti.

------------------------------------------------------------------------

# 31. Future Development

Fitur lanjutan yang tidak wajib untuk MVP:

-   Windows support,
-   macOS support,
-   voice command,
-   offline LLM,
-   visual verification,
-   screenshot-based verification,
-   multi-device control,
-   task scheduling,
-   automation workflows,
-   richer desktop UI,
-   plugin/action system.

------------------------------------------------------------------------

# 32. Contoh End-to-End

User dari Android:

``` text
"Buka VS Code, buka project meetgo, lalu jalankan go test ./..."
```

Planner menghasilkan:

``` text
1. open_application(code)
2. open_directory("/path/to/meetgo")
3. run_command("go test ./...")
```

Guardrails:

``` text
Action 1 → SAFE
Action 2 → SAFE
Action 3 → REVIEW
```

Android menampilkan:

``` text
This task requires approval.

Action:
Run go test ./...

Working directory:
/path/to/meetgo

[Approve] [Cancel]
```

User approve.

Desktop agent:

``` text
Open VS Code
      ↓
Open project
      ↓
Run go test ./...
```

Verification:

``` text
Process: VS Code → PASS
Project directory → PASS
Test command → PASS
```

Android:

``` text
Task completed

✓ VS Code opened
✓ Project opened
✓ Tests passed

Verification: PASS
```

------------------------------------------------------------------------

# 33. Prinsip Desain

1.  **AI plans, policy decides, agent executes, verifier confirms.**
2.  Tidak ada satu komponen yang memiliki kontrol penuh.
3.  Default behavior harus aman.
4.  Aksi berisiko harus transparan kepada pengguna.
5.  Semua execution penting harus dapat diaudit.
6.  Prototype harus dapat diuji tanpa merusak sistem utama.
7.  Fitur tambahan tidak boleh mengorbankan scope penelitian.

------------------------------------------------------------------------

# 34. Definisi Arsitektur Singkat

``` text
                    ┌──────────────────┐
                    │   Android App    │
                    │ Kotlin/Compose   │
                    └────────┬─────────┘
                             │
                       HTTPS/WebSocket
                             │
                             ▼
                    ┌──────────────────┐
                    │    Go Backend    │
                    │ API + Task Mgmt  │
                    └───────┬───┬──────┘
                            │   │
                            │   └──────────────┐
                            ▼                  │
                    ┌──────────────────┐       │
                    │ Python Planner   │       │
                    │ FastAPI + LLM    │       │
                    └────────┬─────────┘       │
                             │                 │
                             ▼                 │
                    ┌──────────────────┐       │
                    │ Safety Guardrails│       │
                    │ Policy + Risk    │       │
                    └────────┬─────────┘       │
                             │                 │
                             ▼                 │
                    ┌──────────────────┐       │
                    │ Desktop Agent    │◄──────┘
                    │ Go               │
                    └────────┬─────────┘
                             │
                             ▼
                    ┌──────────────────┐
                    │ Linux Desktop    │
                    └────────┬─────────┘
                             │
                             ▼
                    ┌──────────────────┐
                    │ Verification     │
                    │ Engine           │
                    └──────────────────┘
```

------------------------------------------------------------------------

# 35. Catatan Implementasi Awal

Urutan pengerjaan yang disarankan:

``` text
1. Android project        ✓
2. Desktop Agent
3. Go Backend
4. Device pairing
5. Basic WebSocket
6. Basic desktop action
7. Verification
8. Task model
9. Planner
10. Guardrails
11. Approval UI
12. History
13. Evaluation
```

**Jangan mulai dari LLM.**

LLM adalah salah satu komponen, bukan inti seluruh sistem. Bangun dulu
jalur deterministic:

``` text
Android
→ Backend
→ Agent
→ Execute
→ Verify
→ Android
```

Setelah jalur itu stabil, baru tambahkan:

``` text
Natural Language
→ Planner
```

dan kemudian:

``` text
Planner
→ Guardrails
```

Dengan urutan tersebut, jika LLM bermasalah, sistem utama tetap bisa
diuji menggunakan structured action secara langsung.

------------------------------------------------------------------------

# 36. Status Saat Ini

Per 1 October 2026:

-   [x] Project directory dibuat
-   [x] Android project dibuat
-   [x] Kotlin + Jetpack Compose project berhasil build
-   [ ] Desktop Agent
-   [ ] Go Backend
-   [ ] Planner
-   [ ] Database
-   [ ] Device Pairing
-   [ ] WebSocket
-   [ ] Safety Guardrails
-   [ ] Verification Engine
-   [ ] Evaluation Dataset
-   [ ] Final UI
