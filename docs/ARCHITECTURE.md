# SynthaMind Architecture

## 🧠 High-Level Overview

SynthaMind is built on a modular, cognitive-agent architecture inspired by neuro-symbolic and recursive systems. Each component is designed to function independently but communicate and evolve collectively.

---

## Core Components

### 1. `main.go`

Entry point for initializing and running the system.

- Bootstraps core services
- Spawns base agents
- Starts task manager and feedback loops

### 2. `agents/`

Home of autonomous agents with diverse roles (e.g. planner, reasoner, memory manager).

- Agents follow a shared interface: `Init()`, `Tick()`, `Report()`
- Designed to be independently testable and replaceable

### 3. `modules/`

Self-contained utility brains or subsystems (e.g. tokenizer, parser, encoder).

- Reusable across agents
- Don’t maintain state, pure utility or plug-in brains

### 4. `taskmanager/`

Orchestrates task assignment, scheduling, and prioritization.

- Can broadcast to agents
- Receives feedback and results

### 5. `chroma/`

Handles vector embedding and semantic memory retrieval.

- Interfaced via a semantic search layer
- Supports plug-and-play vector DB

### 6. `utils/`

Helpers, constants, error handling.

- Keep minimal and strictly generic

### 7. `tests/`

Test suite (unit + integration coming soon).

- Simulations of agent loops
- Reasoning correctness and regression tests

---

## 🌀 Data Flow (Simplified)

```
[Input] → TaskManager → Agents (plan, recall, reason) → Output
                          ↑
                        Modules
                          ↓
                      Semantic Memory (chroma)
```

---

## 🔄 Agent Lifecycle

Each agent runs in its own goroutine:

1. `Init()` – Load internal state or dependencies
2. `Tick()` – Run periodically or on trigger
3. `Report()` – Share outputs with manager or other agents

---

## Planned Improvements

- Move agent interfaces to `internal/agents/interfaces`
- Make `cmd/synthamind/main.go` the real entrypoint
- Add gRPC interface for external triggering
- Replace file-based `.env` with Viper config loader

---

## Contributing

All modules must:

- Be modular and testable
- Avoid side effects unless intentional
- Use structured logging

---

## Goal

Enable long-term, recursive cognitive evolution using pluggable autonomous agents.

> SynthaMind isn’t just a program — it’s a synthetic cognitive system.
