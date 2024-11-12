# 😸 Exploding Kitten Game

Welcome to the Exploding Kitten game! This is a single-player card game implemented using React for the frontend, Golang for the backend, and Redis for database storage.

## 🛠️ Tech Stack

- **Frontend:** React with Redux for state management.
- **Backend:** Golang (using Gorilla Mux and Redis client libraries).
- **Database:** Redis for leaderboard and game state storage.
- **Deployment:** Docker for containerization.

## 🚀 How to Run Locally

Follow these steps to run the game on your local machine:

### Step 1: Clone the Repository

To set up the game locally, start by cloning the repository:

```bash
git clone https://github.com/your-repo/exploding-kitten.git
cd exploding-kitten
```


### Step 2: Setup the Backend

The backend is powered by Go and uses Redis for data storage. Follow these steps to set it up:

#### Prerequisites
- Install **Golang** (v1.20 or later).
- Install **Redis** (locally or through Docker).

#### Steps:

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Start Redis:
   ```bash
   redis-server
   ```
4. Run the Backend Server:
   ```bash
   go run main.go
   ```


### Step 3: Setup the Frontend

#### Steps:
1. Navigate to the frontend directory:
   ```bash
   cd ../frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the Frontend Development Server:
   ```bash
   npm start
   ```


### Step 4: Play the Game

Once both the backend and frontend are running, you can play the game.

#### Instructions:
1. Open your browser.
2. Visit [http://localhost:3000](http://localhost:3000).
3. Enter a username to start the game.
4. Enjoy! 🎉
