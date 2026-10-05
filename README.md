# Safety Guardrails

Penerapan Mekanisme Safety Guardrails dan Verifikasi Otomatis pada Asisten Cerdas untuk Eksekusi Tugas Desktop Jarak Jauh.

## Deskripsi

Safety Guardrails adalah sistem asisten cerdas yang memungkinkan pengguna memberikan instruksi dari perangkat Android untuk menjalankan tugas tertentu pada komputer desktop Linux melalui jaringan.

Sistem dirancang dengan mekanisme safety guardrails untuk memastikan bahwa tindakan yang dihasilkan oleh planner tidak langsung dieksekusi tanpa pemeriksaan.

Alur utama sistem:

User
→ Android
→ Backend
→ Planner
→ Safety Guardrails
→ Approval jika diperlukan
→ Desktop Agent
→ Execution
→ Verification
→ Result

Prinsip utama:

> AI plans. Policy decides. Agent executes. Verifier confirms.

## Tujuan

Sistem ini bertujuan untuk:

- Menerima instruksi pengguna dalam bahasa natural.
- Mengubah instruksi menjadi tindakan terstruktur.
- Memeriksa keamanan setiap tindakan sebelum dieksekusi.
- Meminta persetujuan pengguna untuk tindakan berisiko.
- Memblokir tindakan yang tidak diizinkan.
- Mengeksekusi tindakan melalui desktop agent.
- Memverifikasi apakah tindakan benar-benar berhasil.
- Menyimpan riwayat dan audit proses eksekusi.

## Konsep Safety Guardrails

Setiap tindakan diklasifikasikan menjadi tiga tingkat risiko:

| Risk Level | Behavior |
|------------|----------|
| SAFE | Dapat dijalankan otomatis |
| REVIEW | Membutuhkan persetujuan pengguna |
| BLOCKED | Tidak dapat dijalankan |

### SAFE

SAFE berarti tindakan dianggap berisiko rendah dan dapat dijalankan secara otomatis.

Contoh:

- Membaca informasi sistem.
- Membaca informasi proses.
- Membaca metadata file.
- Membuat temporary directory pada workspace yang diizinkan.
- Membuat temporary file pada workspace yang diizinkan.

### REVIEW

REVIEW berarti tindakan dapat dijalankan tetapi memiliki efek samping yang membutuhkan persetujuan pengguna.

Contoh:

- Menutup aplikasi.
- Memindahkan file.
- Mengubah file pada workspace.
- Menjalankan perintah project yang telah ditentukan.
- Shutdown.
- Sleep.
- Lock desktop.

### BLOCKED

BLOCKED berarti tindakan tidak diperbolehkan oleh sistem.

Contoh:

- Mengakses credential atau data sensitif.
- Mengakses path di luar scope yang diizinkan.
- Menjalankan arbitrary shell command.
- Menghapus file sistem.
- Menonaktifkan mekanisme keamanan.
- Melakukan destructive recursive deletion.

## Verification

Sistem tidak hanya mengandalkan exit status dari proses.

Setelah tindakan dijalankan, sistem melakukan verification terhadap kondisi desktop untuk memastikan bahwa perubahan yang diharapkan benar-benar terjadi.

Contoh:

| Action | Verification |
|--------|--------------|
| Open application | Memeriksa proses aplikasi |
| Create file | Memeriksa keberadaan file |
| Move file | Memeriksa lokasi sumber dan tujuan |
| Run tests | Memeriksa exit code dan output |
| System information | Memeriksa data yang diterima |

Hasil verification:

- VERIFIED
- FAILED
- UNCERTAIN

## Architecture

Sistem terdiri dari beberapa komponen utama:

### Android Client

Aplikasi Android yang digunakan pengguna untuk:

- Mengirim instruksi.
- Melihat task.
- Melihat action yang direncanakan.
- Memberikan approval.
- Melihat status eksekusi.
- Melihat hasil verification.
- Melihat history.

Technology:

- Kotlin
- Jetpack Compose

### Go Backend

Backend utama yang bertanggung jawab untuk:

- REST API.
- WebSocket communication.
- Task management.
- Device management.
- Planner communication.
- Guardrail coordination.
- Desktop agent communication.
- Database persistence.
- Audit logging.

Technology:

- Go
- Gin
- WebSocket

### Planner Service

Service yang mengubah instruksi natural language menjadi structured actions.

Technology:

- Python
- FastAPI
- Pydantic
- LLM API

Planner hanya memberikan rekomendasi tindakan.

Planner tidak memiliki kewenangan untuk menentukan apakah tindakan boleh dijalankan.

### Safety Guardrails

Komponen yang melakukan pemeriksaan keamanan terhadap action yang dihasilkan planner.

Tanggung jawab:

- Memvalidasi action type.
- Memvalidasi parameter.
- Memvalidasi target.
- Memeriksa protected resource.
- Memeriksa allowed scope.
- Menentukan risk level.
- Menentukan apakah action SAFE, REVIEW, atau BLOCKED.

### Desktop Agent

Agent berbasis Go yang berjalan pada komputer Linux.

Tanggung jawab:

- Terhubung ke backend.
- Menerima action yang telah diizinkan.
- Melakukan local validation.
- Menjalankan action.
- Mengirim execution result.
- Menjalankan atau memicu verification.

Desktop Agent tidak boleh menjalankan arbitrary shell command secara default.

### Verification Engine

Komponen yang memeriksa kondisi sistem setelah action dijalankan.

Tanggung jawab:

- Memeriksa perubahan sistem.
- Memastikan expected state tercapai.
- Menentukan hasil verification.

### PostgreSQL

Database untuk menyimpan:

- Users.
- Devices.
- Tasks.
- Actions.
- Executions.
- Verifications.
- Audit logs.

## Execution Flow

```text
User
  ↓
Android Client
  ↓
Go Backend
  ↓
Planner
  ↓
Structured Actions
  ↓
Safety Guardrails
  ↓
┌───────────────┬────────────────┬────────────────┐
│ SAFE          │ REVIEW         │ BLOCKED        │
│               │                │                │
│ Auto approve  │ User approval  │ Reject         │
└───────┬───────┴────────┬───────┴────────────────┘
        │                │
        └───────┬────────┘
                ↓
        Desktop Agent
                ↓
            Executor
                ↓
           Verification
                ↓
         Backend / Database
                ↓
          Android Client
                ↓
              User
```

## Technology Stack

| Component | Technology |
|-----------|------------|
| Android | Kotlin + Jetpack Compose |
| Backend | Go + Gin |
| Communication | REST API + WebSocket |
| Planner | Python + FastAPI |
| LLM Integration | LLM API |
| Schema Validation | Pydantic |
| Desktop Agent | Go |
| Database | PostgreSQL |
| Containerization | Docker + Docker Compose |
| Version Control | Git |
| Target Desktop | Linux |

## Project Structure

```text
safety-guardrails/
├── PRD-Safety-Guardrails.md
├── README.md
├── .gitignore
│
├── mobile/
│   └── Android application
│
├── backend/
│   └── Go backend
│
├── planner/
│   └── Python planner service
│
├── desktop-agent/
│   └── Go desktop agent
│
└── docs/
    ├── architecture.md
    ├── flow.md
    ├── guardrails.md
    ├── api.md
    ├── database.md
    ├── research.md
    │
    └── diagrams/
        ├── architecture.md
        ├── task-flow.md
        ├── guardrail-flow.md
        ├── sequence.md
        ├── database-erd.md
        └── component.md
```

## Development Status

Project development is currently in the initial implementation stage.

### Completed

- [x] Project repository structure
- [x] Initial PRD
- [x] Initial system architecture
- [x] Initial task flow
- [x] Initial guardrail flow
- [x] Initial sequence diagram
- [x] Initial database ERD
- [x] Initial component diagram
- [x] Android project initialization
- [x] Desktop Agent project initialization
- [x] Backend project initialization
- [x] Planner project initialization

### In Progress

- [ ] Android UI
- [ ] Desktop Agent
- [ ] Go Backend
- [ ] Device pairing
- [ ] WebSocket communication
- [ ] Basic desktop action
- [ ] Verification engine
- [ ] Task model
- [ ] Planner integration
- [ ] Guardrail implementation
- [ ] Approval mechanism
- [ ] PostgreSQL integration
- [ ] Audit logging
- [ ] Evaluation

## Development Principle

Development will initially focus on deterministic components before integrating the LLM.

Recommended implementation order:

1. Android application
2. Desktop Agent
3. Go Backend
4. Device pairing
5. WebSocket communication
6. Basic desktop action
7. Verification
8. Task model
9. Planner
10. Safety Guardrails
11. Approval mechanism
12. Task history
13. Evaluation

This approach allows the core execution pipeline to be tested before introducing LLM-generated actions.

## Security Principles

The system follows these principles:

1. The planner cannot directly execute desktop actions.
2. Every action must pass safety validation.
3. Risky actions require explicit approval.
4. Blocked actions cannot be approved through the normal flow.
5. The desktop agent performs local validation before execution.
6. Arbitrary shell execution is disabled by default.
7. Execution results are verified.
8. Important events are recorded in the audit log.
9. Safety validation fails closed when validation cannot be completed.
10. The system follows the principle of least privilege.

## Research Evaluation

The system will be evaluated using several metrics.

### Guardrail Effectiveness

Measured using:

- True Positive
- True Negative
- False Positive
- False Negative
- Precision
- Recall

### Verification Accuracy

Measured by comparing verification results with the actual desktop state.

### Task Execution

Measured using:

- Task success rate
- Task failure rate
- Timeout rate

### Latency

Measured across:

- Android → Backend
- Backend → Planner
- Planner → Guardrail
- Guardrail → Agent
- Agent → Verification
- Verification → Android

## Research Questions

The initial research questions are:

1. How effective are rule-based safety guardrails at preventing unauthorized desktop actions?
2. How accurately does automatic verification determine task success compared with the actual desktop state?
3. What latency overhead is introduced by the guardrail and verification mechanisms?
4. How reliable is Android-to-desktop task execution over a network connection?

These research questions may be refined after the literature review and implementation.

## Documentation

Additional documentation:

- `PRD-Safety-Guardrails.md`
- `docs/architecture.md`
- `docs/flow.md`
- `docs/guardrails.md`
- `docs/api.md`
- `docs/database.md`
- `docs/research.md`

Architecture diagrams:

- `docs/diagrams/architecture.md`
- `docs/diagrams/task-flow.md`
- `docs/diagrams/guardrail-flow.md`
- `docs/diagrams/sequence.md`
- `docs/diagrams/database-erd.md`
- `docs/diagrams/component.md`

## Project Status

Status:

Early Development

The architecture, database model, API design, and diagrams are initial designs.

These components may change during implementation based on technical findings, testing results, security requirements, and research evaluation.

The final architecture and documentation will reflect the system that is actually implemented and evaluated.
```
