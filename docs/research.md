# Safety Guardrails - Research Specification

## 1. Research Topic

Project title:

Penerapan Mekanisme Safety Guardrails dan Verifikasi Otomatis pada Asisten Cerdas untuk Eksekusi Tugas Desktop Jarak Jauh

Main research topic:

Safety mechanism and automatic verification for AI-assisted remote desktop task execution.

---

## 2. Research Problem

AI systems can generate useful instructions for desktop automation.

However, directly executing AI-generated commands can introduce safety risks.

The system therefore requires:

1. structured action generation
2. deterministic safety validation
3. user approval for risky actions
4. controlled execution
5. automatic verification

The research focuses on whether these mechanisms can improve the safety and reliability of remote desktop task execution.

---

## 3. Research Questions

RQ1:

How effective are rule-based safety guardrails at preventing unauthorized or dangerous desktop actions?

RQ2:

How accurately can automatic verification determine whether a desktop action actually succeeded?

RQ3:

What latency overhead is introduced by safety validation and automatic verification?

RQ4:

How reliable is Android-to-desktop task execution over a network?

---

## 4. Research Objectives

1. Design a safety guardrail architecture for AI-assisted desktop automation.
2. Implement deterministic action classification.
3. Implement controlled desktop execution.
4. Implement automatic execution verification.
5. Evaluate safety effectiveness.
6. Evaluate verification accuracy.
7. Measure execution latency.
8. Measure reliability of the complete system.

---

## 5. Main Variables

Independent variables:

- action type
- action risk level
- guardrail policy
- network condition
- execution type

Dependent variables:

- guardrail decision accuracy
- dangerous action blocking rate
- false positive rate
- false negative rate
- verification accuracy
- task success rate
- latency
- timeout rate

---

## 6. Safety Evaluation

Actions should be classified into expected categories.

Example dataset:

SAFE:
- get system information
- read file metadata
- create temporary directory in allowed workspace

REVIEW:
- shutdown
- move files
- close application
- run predefined project command

BLOCKED:
- delete system files
- access protected credentials
- unrestricted shell execution
- disable safety mechanisms

The expected classification becomes the ground truth.

---

## 7. Safety Metrics

Use:

True Positive
True Negative
False Positive
False Negative

From these calculate:

Precision
Recall
False Positive Rate
False Negative Rate

The exact interpretation of positive and negative classes must be defined clearly before experiments.

---

## 8. Verification Evaluation

Verification tests should compare actual system state with verification result.

Example:

Action:
create file

Actual state:
file exists

Expected verification:
VERIFIED

If file does not exist:

Expected verification:
FAILED

For uncertain conditions:

Expected verification:
UNCERTAIN

Metrics:

verification accuracy
false verification rate
verification latency

---

## 9. Task Success Evaluation

Task success should require:

1. Correct planning.
2. Guardrail approval.
3. Successful execution.
4. Successful verification.

A task must not be considered fully successful only because the process returned exit code 0.

---

## 10. Latency Evaluation

Measure:

T1:
Android request -> backend

T2:
Backend -> planner

T3:
Planner -> guardrail

T4:
Guardrail -> agent

T5:
Agent execution

T6:
Verification

T7:
Result -> Android

Total task latency:

T_total =
T1 + T2 + T3 + T4 + T5 + T6 + T7

Measurements should be repeated across multiple test cases.

---

## 11. Network Evaluation

Test under different network conditions where possible.

Example:

- stable LAN
- normal Wi-Fi
- higher latency
- temporary disconnection

Measure:

- task success
- timeout
- reconnection
- duplicate execution
- result consistency

---

## 12. Baseline

Possible baseline:

Direct execution without the guardrail layer.

Compare:

Baseline:
Planner -> Agent -> Execute

Proposed:
Planner -> Guardrail -> Approval -> Agent -> Execute -> Verify

The comparison should focus on measurable differences such as:

- dangerous action execution
- false blocking
- task success
- verification accuracy
- latency

---

## 13. Experimental Dataset

Create a fixed task dataset.

Example categories:

Category A:
Read-only tasks.

Category B:
File manipulation.

Category C:
Application control.

Category D:
System state changes.

Category E:
Dangerous or prohibited requests.

Each test case should contain:

- natural language instruction
- expected structured action
- expected risk level
- expected policy decision
- expected execution result
- expected verification result

---

## 14. Research Procedure

Step 1:
Create task dataset.

Step 2:
Send tasks to planner.

Step 3:
Record planner output.

Step 4:
Run guardrail classification.

Step 5:
Record guardrail decision.

Step 6:
Approve REVIEW actions when appropriate.

Step 7:
Execute allowed actions.

Step 8:
Run verification.

Step 9:
Compare result with ground truth.

Step 10:
Calculate evaluation metrics.

Step 11:
Analyze errors.

---

## 15. Error Analysis

Important errors include:

False positive:
Safe action incorrectly classified as dangerous.

False negative:
Dangerous action incorrectly allowed.

Execution failure:
Approved action cannot be executed.

Verification false positive:
System reports success although state is incorrect.

Verification false negative:
System reports failure although state is correct.

Planner error:
Natural language is converted to incorrect action.

---

## 16. Research Limitations

Possible limitations:

1. Initial implementation targets Linux.
2. Supported actions are limited.
3. Safety policies are initially rule-based.
4. Verification methods depend on available system information.
5. Network conditions may affect latency.
6. LLM behavior may change depending on model and prompt.
7. The research dataset may not represent every real-world desktop task.

These limitations must be stated honestly in the final research report.

---

## 17. Novelty / Research Gap

Do not claim that the project is completely novel without literature review.

Literature review should investigate:

- AI desktop agents
- computer-use agents
- remote desktop automation
- AI safety guardrails
- command execution safety
- policy-based execution
- automatic verification
- agentic AI safety

The research gap must be established from actual academic literature and existing systems.

---

## 18. Expected Contribution

Potential contribution:

A layered architecture that combines:

Natural Language Planning
+
Deterministic Safety Guardrails
+
Controlled Desktop Execution
+
Automatic Verification

The contribution must be evaluated experimentally rather than assumed.

---

## 19. Final Evaluation

The final evaluation should answer:

1. Can the system prevent prohibited actions?
2. Can the system distinguish SAFE, REVIEW, and BLOCKED actions?
3. Can risky actions require explicit approval?
4. Can desktop actions be executed reliably?
5. Can execution results be automatically verified?
6. What is the latency overhead?
7. What failure cases remain?

---

## 20. Research Documentation

Every experiment should record:

- test ID
- input instruction
- planner output
- guardrail result
- approval result
- execution result
- verification result
- latency
- final status
- error notes

This data will be used in the final analysis and thesis report.
