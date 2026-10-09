I'm building a crab-viewing website that will act like a tamagotchi, and I want you to help me plan and build it step by step while teaching me along the way. Please read the context and teaching preferences below before answering.

## About me
- Final-year Software Engineering student, building this as a portfolio/team project to show backend, real-time systems, multiple databases, containerization, and orchestration skills for my job search.
- New to Go, Gin, Docker, and Postgres. I've used SQLite before, so comparisons to it help.
- OS: Ubuntu.
- Small team (roles likely: backend, frontend/React, real-time/infra). I want a finished, working end-to-end system, not a sprawling half-built one.

## The project
A website where visitors watch crabs in a game-like view, with tamagotchi-style care mechanics (hunger, happiness, etc.; details to be designed later). Live chat comes FIRST. Login comes after chat. The crab/tamagotchi logic and graphics come after that.

This project replaced an earlier idea (a Git repo browser). All go-git code is being dropped.

## Stack (confirmed, please work within it)
- Backend: Go + Gin
- Real-time: WebSockets (gorilla/websocket or similar)
- PostgreSQL: users, accounts, persistent crab stats (LATER, after chat)
- MongoDB: chat messages and history
- Redis: Pub/Sub to scale chat across multiple instances, plus presence and cooldowns
- Frontend: Next.js (React); the crab view is a client-only component
- Graphics: leaning PixiJS with pixel-art sprite sheets (alternatives I considered: Phaser, Rive, plain Canvas). Pixel art is the style direction.
- Docker, Docker Compose, Kubernetes (Minikube), GitHub Actions, then Prometheus + Grafana, all at the end

## Architecture intent
Eventually two separate services, because they scale differently:
1. a chat/realtime service (long-lived connections, stateful)
2. a game/account API (request-response: auth, crab stats)
Don't split them yet. Start with only the chat service and split when the game API exists. I want to be able to explain and defend this design in interviews.

## Build order
1. Chat service: Gin + WebSocket, one shared room, in-memory only, guest nicknames (no login). Test with a plain HTML page first.
2. Next.js chat UI with a placeholder crab image.
3. MongoDB: persist messages, load history on join.
4. Login with Postgres, attaching real identities to chat.
5. Redis Pub/Sub for multiple chat instances.
6. Crab rendering (PixiJS) and tamagotchi logic.
7. Docker Compose, Kubernetes, CI/CD, observability.

## Where I am right now
- Go module path: github.com/catus64/repo-observer/api (I may rename it later)
- Folder layout so far:
  api/
    cmd/server/main.go
    internal/handlers/*.go        (Gin handlers as methods on structs)
    internal/router/router.go
  (internal/gitrepo/ and the commit handlers are leftovers I'm deleting, followed by `go mod tidy`)
- Layering I follow: handlers know about HTTP, services don't, repositories talk to databases. No gin.Context outside handlers.
- Config comes from environment variables. Logging uses log/slog.
- I already understand: Gin basics (Engine, Context, route groups), sentinel errors with errors.Is and errors.As, defer, iterators, dependency injection via struct fields, and how to run Postgres in Docker (official postgres:16 container named observer-pg, with a volume). I have not written any WebSocket code yet.
- Postgres is running but unused until the login phase.

## Open design questions (please help me decide, one at a time, not all at once)
- Shared tank (everyone watches the same crabs live) or personal crab per user?
- Who makes the art (team member, itch.io asset packs, undecided)?
- Chat scope: one global room, or a room per tank?

## How to teach me (important)
- **One small step at a time.** Give me one concept or one small piece of code, then stop and wait for me to confirm it works before continuing. Don't dump a whole feature (service + handler + routes + frontend) in one message. I get lost when too much arrives at once.
- **Explain the why**, not just the what. When you show Go syntax or a library feature I might not know, explain it in plain language, with a tiny standalone example if it helps.
- Tell me exactly what command to run and what output I should expect, so I can tell whether it worked.
- If I paste an error, help me read it before giving the fix.
- End each step with a short exercise or check I can do myself.
- Point out common beginner mistakes. If something you showed earlier was wrong or incomplete, say so.
- If a question has several reasonable answers, briefly give the tradeoffs and tell me which you'd pick.
- Keep the scope finishable: prefer the simplest design that works, and tell me when something is over-building.

Start by asking me the first open design question (shared tank vs personal crab), then begin step 1: a minimal WebSocket echo endpoint in Gin, explaining how WebSockets differ from normal HTTP requests before showing code.
