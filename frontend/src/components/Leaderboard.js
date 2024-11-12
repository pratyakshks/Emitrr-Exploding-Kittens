import React, { useState, useEffect } from 'react';

const Leaderboard = ({ socket }) => {
    const [leaderboard, setLeaderboard] = useState([]);

    useEffect(() => {
        socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            setLeaderboard((prev) => [
                ...prev.filter((p) => p.username !== data.username),
                data,
            ]);
        };
    }, [socket]);

    return (
        <div>
            <h2>Leaderboard</h2>
            <ul>
                {leaderboard.map((user, index) => (
                    <li key={index}>
                        {user.username}: {user.points} points
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Leaderboard;
