<br />
<div align="center">

<img src="assets/images/logo.png" width="140" alt="Rem - ReMusika">

<br />
<br />

[![Release version](https://img.shields.io/github/v/release/seinkytarlicht/remusika?color=brightgreen&label=Latest&style=for-the-badge)](https://github.com/seinkytarlicht/remusika/releases/latest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/seinkytarlicht/remusika?style=for-the-badge)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/m/seinkytarlicht/remusika?color=brown&label=Commits&style=for-the-badge)](https://github.com/seinkytarlicht/remusika/commits)
![GitHub Repo stars](https://img.shields.io/github/stars/seinkytarlicht/remusika?color=orange&style=for-the-badge)

English | [Indonesia](README_ID.md)

</div>

# ReMusika

A lightweight(probably) local music player powered by Go and Nuxt.

All music playback happens locally on your device. No internet connection, external services, or additional setup required.

Simply run the application, open your browser, and start listening to your music library.

## Preview

> Screenshots and demo GIFs will be added soon.

## Why?

This project was created as an experiment to build a fully local music player using web technologies while keeping the user experience simple.

The goal is to provide a lightweight music player that:

- Runs entirely on the user's device
- Requires no internet connection
- Does not require Node.js or additional runtime installations
- Can be distributed as a simple executable

## Features

- Automatically scans your Music folder
- Local audio streaming
- Play, Pause, Seek, and Loop controls
- Lightweight system tray integration
- Opens in your default browser
- Fully offline

### Tray Menu

- **Show** — Open the application in your browser
- **Exit** — Stop all running services and close the application

## Getting Started

### Installation

1. Download the latest release, [here](https://github.com/seinkytarlicht/remusika/releases)
2. Run the application
3. Your default browser will open automatically

If the browser does not open automatically, visit:

```text
http://127.0.0.1:41991
```

Your music library is now ready to use.

> Make sure your music files are located inside the default Music folder of your operating system.

> Known Issue: If the page is showing a blank page, try hard reload (`Shift` + `R`) or try clear your browser cache

## Music Folder Location

### Windows

```text
C:\Users\<your-username>\Music
```

### Linux

```text
/home/<your-username>/Music
```

## Tech Stack

- Go
- Fiber
- Nuxt

## Running the Project

### Clone Repository

```bash
git clone https://github.com/seinkytarlicht/remusika.git
cd remusika
```

### Run Frontend

```bash
cd frontend
bun install
bun run dev
```

### Run Backend

From the project root:

```bash
go run cmd/app/main.go --server
```

## Project Structure

```text
/
├── config
├── frontend
├── helper
├── internal
└── main.go
```

## Building the Application

### Generate Frontend

```bash
cd frontend
bun run generate
```

### Build Backend

From the project root:

```bash
go build -o dist/remusika cmd/app/main.go
```

The generated binary will be available in:

```text
dist/remusika
```

## Roadmap

- [x] Music library scanning
- [x] Audio streaming
- [x] Tray integration
- [x] Search music
- [ ] Logo
- [x] Media Session API
- [ ] Playlist support

## License
