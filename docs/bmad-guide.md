# BMAD Integration Guide

This project is integrated with the **BMAD (Big Model Analysis & Design)** methodology for AI-assisted development. This guide explains how to leverage the specialized agents and workflows configured for this repository.

## Get Started

Ensure you have the BMAD CLI installed (via `npx`):

```bash
npx bmad-method status
```

## 1. Collaborative Development ("Party Mode")

**Party Mode** is the primary workflow for multi-agent brainstorming, architectural review, and planning. It assembles a team of specialized AI agents (Architect, Analyst, Dev, QA, etc.) to discuss the project in real-time.

### How to Run

There are two primary ways to start a Party Mode session:

1.  **Interactive Menu**:
    ```bash
    npx bmad-method
    ```
    Select **Workflows** -> **Select a Workflow** -> **party-mode**.

2.  **Direct Workflow Command**:
    (If configured in your environment context, simply mention `@party-mode` or follow your specific agent instructions).

### The Agent Team

*   **🧙 BMad Master**: Orchestrator and task executor.
*   **🏗️ Winston (Architect)**: Focuses on system design, scalability, and tech stack decisions.
*   **📊 Mary (Analyst)**: Analyzes business requirements, user flows, and product value.
*   **💻 Amelia (Dev)**: Focuses on implementation details, code quality, and testing.
*   **🧪 Quinn (QA)**: Handles test strategy and quality assurance.
*   **📚 Paige (Tech Writer)**: Documentation and diagrams.

## 2. Session Artifacts

All BMAD sessions and generated artifacts are stored in the `bmad-session/` directory for historical tracking:

*   **`bmad-session/full-transcript.md`**: The complete log of major architectural discussions.
*   **`bmad-session/generated-files/`**: Diagrams (Mermaid), reports, and other outputs from sessions.
*   **`bmad-session/code-analysis-report.md`**: Detailed code audit reports.

## 3. Configuration

BMAD configuration files are located in `_bmad/`:
*   `_bmad/core/config.yaml`: Core project settings.
*   `_bmad/_config/agents/`: Agent customization (memory, personality, constraints).
