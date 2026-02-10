# Landing Page

A simple landing page for tutorials built with React (frontend) and Go (backend).

## Features

- Modern React frontend with responsive design
- RESTful Go backend API
- Sample tutorials listing
- Beautiful gradient UI
- Mobile-friendly layout

## Prerequisites

- Node.js (v16 or higher)
- Go (v1.21 or higher)

## Project Structure

```
.
├── main.go              # Go backend server
├── go.mod               # Go module file
├── frontend/            # React frontend
│   ├── public/          # Static files
│   ├── src/             # React source code
│   └── package.json     # Node dependencies
└── README.md
```

## Getting Started

### Backend Setup

1. Build and run the Go server:

```bash
go build -o landingpage main.go
./landingpage
```

The API server will start on port 8080 by default.

You can specify a custom port:

```bash
PORT=3001 ./landingpage
```

### Frontend Setup

1. Navigate to the frontend directory:

```bash
cd frontend
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm start
```

The React app will start on port 3000 and proxy API requests to the Go backend.

## API Endpoints

- `GET /api/health` - Health check endpoint
- `GET /api/tutorials` - Get list of tutorials

## Building for Production

### Backend

```bash
go build -o landingpage main.go
```

### Frontend

```bash
cd frontend
npm run build
```

The production build will be created in the `frontend/build` directory.

## Technologies Used

- **Frontend**: React, CSS3
- **Backend**: Go (Golang)
- **API**: RESTful JSON API

## License

MIT