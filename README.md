# Go Todo

A full-stack todo application built with **Go** backend and **Next.js** frontend.

## Tech Stack

- **Backend:** Go 1.20+ with `net/http`
- **Frontend:** Next.js 14 + React 18 + TypeScript
- **Storage:** In-memory (resets on restart)

## Requirements

- Go 1.20+
- Node.js 18+
- npm

## Quick Start

### Production Mode (Single Server)

Build and serve everything from one port:

```bash
# Install frontend dependencies
cd frontend
npm install

# Build Next.js app
npm run build

# Start Go server (serves built frontend + API)
cd ..
go run .
```

Then open **http://localhost:8080** in your browser.

### Development Mode (Two Terminals)

For faster frontend development with hot reload:

**Terminal 1 - Backend API:**
```bash
go run .
# Runs on http://localhost:8080
```

**Terminal 2 - Frontend Dev Server:**
```bash
cd frontend
npm install
npm run dev
# Runs on http://localhost:3000
# Proxies /api calls to http://localhost:8080
```

Then open **http://localhost:3000** in your browser.

## Building

Build frontend for production:
```bash
cd frontend
npm run build
```

Build Go binary:
```bash
go build -o gotodo
./gotodo
```

## API Endpoints

- `GET /api/todos` - Get all todos
- `POST /api/todos` - Create new todo
  - Body: `{"title":"My task"}`
- `PUT /api/todos/{id}` - Update todo
  - Body: `{"title":"Updated","completed":true}`
- `DELETE /api/todos/{id}` - Delete todo

## Features

✅ Create, read, update, delete todos  
✅ Mark todos as complete  
✅ Next.js Server Components ready  
✅ TypeScript throughout  
✅ Real-time updates  
✅ Responsive design

