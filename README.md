I'm a final-year Software Engineering student starting a portfolio/team project. I want you to help me plan and build it step by step, starting with architecture and database schema, then moving into implementation.

PROJECT: Git Repo Observer with Live Messaging

CONCEPT:
A read-only web platform for browsing a single, pre-selected Git repository (commit history, file tree, file contents, diffs, branches), combined with a permissioned social layer: a group chatroom tied to the repo, and private 1:1 DMs between users with access. No git push/pull/clone functionality, purely observational/visual browsing of one chosen repo. Permissions gate who can view the repo and who can chat.

GOALS:
- Build a finished, working end-to-end system, not a sprawling half-built one
- Use this to demonstrate backend, real-time systems, multiple database paradigms, containerization, and orchestration skills for my job search
- Being built with a small team (roles likely: backend/git core, frontend/React, real-time/infra, possibly shared DevOps)

TECH STACK (confirmed, please work within this):
- Backend: Go + Gin
- Git reading: go-git library (NOT implementing git protocol from scratch, read-only access to one local repo: commits, trees, diffs, branches)
- Relational DB: PostgreSQL (users, permissions, DM metadata)
- NoSQL DB: MongoDB (chat messages, DM history)
- Cache/Pub-Sub: Redis (caching repo data, Pub/Sub for scaling WebSocket chat across multiple instances)
- Real-time: WebSockets (gorilla/websocket or similar) for chatroom + DMs
- Frontend: Next.js React (repo browser UI, chat UI, DM UI, live updates)
- Containerization: Docker, with separate services for: repo-browsing API, chat/WebSocket service, Postgres, MongoDB, Redis
- Orchestration: Kubernetes (Minikube for local dev), deliberately splitting the repo-browsing API and chat service into separate Deployments since they have different scaling needs (read-heavy/cacheable vs. stateful/connection-heavy)
- CI/CD: GitHub Actions
- Observability (later phase): Prometheus + Grafana

KEY ARCHITECTURAL DECISION TO PRESERVE:
Repo-browsing API and chat/WebSocket service are intentionally separate services/deployments, not a monolith, because they scale differently. This is a deliberate design choice I want to be able to explain and defend, not an accident.

BUILD ORDER (agreed plan, please follow this sequence unless there's a strong reason to deviate):
1. Go + Gin + go-git: core API for reading one local repo (commit list, file tree, file contents, diffs)
2. PostgreSQL + auth: users, login, basic permission check (can this user view this repo)
3. React frontend: repo browser UI against that API
4. WebSocket chatroom: single shared room first, in-memory only, just prove real-time delivery works
5. MongoDB: persist chat messages, load history on join
6. Add DMs: extend WebSocket infra to 1:1 messaging
7. Redis Pub/Sub: enable multi-instance chat scaling
8. Docker Compose: all services running together locally
9. Kubernetes (Minikube): split API and chat into separate Deployments
10. CI/CD pipeline, then Prometheus/Grafana observability as a final polish phase

OPEN DESIGN QUESTION TO RESOLVE FIRST:
Permission model granularity, simple (view/no-view, chat/no-chat) vs. more granular (repo-level roles, admin who can grant/approve access). Leaning toward the more granular version since it's more realistic and a good separable piece of backend work for a teammate, but open to your input on tradeoffs.

WHAT I NEED FROM YOU TO START:
1. A proposed PostgreSQL schema (tables, relationships) covering users, permissions, and DM metadata
2. A proposed MongoDB collection structure for chat/DM messages
3. A proposed REST API route list for the repo-browsing endpoints (commits, tree, file contents, diffs, branches)
4. A proposed WebSocket event/message format for chatroom and DM delivery

Please ask me clarifying questions if anything about the scope is ambiguous before generating these, rather than assuming.

do not generate anything yet this is just instruction
