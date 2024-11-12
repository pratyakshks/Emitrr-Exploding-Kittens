import React, { useState } from 'react';
import GameBoard from './components/GameBoard';
import Leaderboard from './components/Leaderboard';

const App = () => {
    const [username, setUsername] = useState('');
    const [gameState, setGameState] = useState(null);
    const [card, setCard] = useState(null);

    const startGame = async () => {
        const res = await fetch(`/start/${username}`, { method: 'POST' });
        const data = await res.json();
        setGameState(data);
    };

    const drawCard = async () => {
        const res = await fetch(`/draw/${username}`, { method: 'POST' });
        const data = await res.json();
        setCard(data.card);
        setGameState(JSON.parse(data.state));
    };

    return (
        <div>
            <h1>😸 Exploding Kitten</h1>
            {!gameState ? (
                <div>
                    <input
                        type="text"
                        placeholder="Enter username"
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <button onClick={startGame}>Start Game</button>
                </div>
            ) : (
                <>
                    <GameBoard gameState={gameState} drawCard={drawCard} card={card} />
                    <Leaderboard socket={new WebSocket('ws://localhost:8080/ws')} />
                </>
            )}
        </div>
    );
};

export default App;
