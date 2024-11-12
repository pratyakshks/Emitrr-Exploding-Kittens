import React from 'react';

const GameBoard = ({ gameState, drawCard, card }) => {
    return (
        <div>
            <h2>Deck: {gameState.deck.length} cards remaining</h2>
            <button onClick={drawCard}>Draw Card</button>
            {card && <h3>Last Card: {card}</h3>}
        </div>
    );
};

export default GameBoard;
