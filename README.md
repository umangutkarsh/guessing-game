# Globetrotter Challenge

🧩 **The Globetrotter Challenge – The Ultimate Travel Guessing Game!**

Globetrotter is a full-stack web application where users are challenged to guess a famous destination based on cryptic clues. Once they guess, they unlock fun facts, trivia, and surprises about the destination. In addition, users can "Challenge a Friend" by sharing an invite link (with a dynamic image) that allows their friends to view their score and play the game.

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Setup Instructions](#setup-instructions)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Dataset & AI Integration](#dataset--ai-integration)
- [Additional Notes](#additional-notes)
- [License](#license)

---

## Features

- **Dataset & AI Integration**

  - Starter dataset provided and expanded to over 100 destinations using AI tools (ChatGPT, OpenAI API, web scraping, etc.).
  - Each destination contains cryptic clues, fun facts, and trivia.

- **Gameplay**

  - Presents 1–2 random clues from a destination.
  - Allows the user to select from multiple possible destination answers.
  - Provides immediate, animated feedback:
    - 🎉 **Correct Answer:** Animates confetti and reveals a fun fact.
    - 😢 **Incorrect Answer:** Shows a sad-face animation and reveals a fun fact.
  - Includes a "Play Again" or "Next" button to load a new random destination.
  - Displays the total user score, tracking both correct and incorrect answers.

- **Challenge a Friend**
  - Users register with a unique username (creating their profile).
  - Clicking the “Challenge a Friend” button generates an invite link along with a dynamic image for sharing via WhatsApp.
  - The invited friend can view the inviter’s score before playing.
  - Anyone with the invitation link can access the full game.

---

## Tech Stack

- **Frontend:**

  - **Next.js** – React framework for server-rendered and static web applications.
  - **React** – Component-based UI development.
  - **Plain CSS** – (Optional: Tailwind CSS can be used if desired).
  - **Axios** – For making HTTP requests.

- **Backend:**

  - **Golang** – For building the API.
  - **Gin** – A lightweight web framework for Golang.
  - **GORM** – ORM for interacting with PostgreSQL.
  - **PostgreSQL** – Database for storing destinations, user profiles, and game state.

- **Other Tools:**
  - **Docker & Docker Compose** – For running PostgreSQL and other services during development.
  - **AI Tools** – (e.g., ChatGPT, OpenAI API) for dataset expansion.

---

## Setup Instructions

### Backend Setup

1. **Prerequisites:**

   - Install [Go](https://golang.org/doc/install) (v1.16+ recommended).
   - Install [Docker](https://www.docker.com/get-started) for running PostgreSQL.

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/guessing-game.git
   cd guessing-game/backend
   cd guessing-game/frontend
   ```

3. **Setup env files**

   Create one env inside backend folder 
   ```bash
   FRONTEND_URL=http://localhost:3000
   ```

   Create one env.local file inside frontend folder
   ```bash
   NEXT_PUBLIC_API_BASE_URL=http://localhost:8080/api/v1
   ```

4. **Install dependencies and start the server**

   Run the following in frontend
   ```bash
   npm install
   npm run dev
   ```

   Run the following in backend
   ```bash
   go run cmd/main.go
   ```
