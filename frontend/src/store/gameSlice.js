import { createSlice } from '@reduxjs/toolkit';

const gameSlice = createSlice({
    name: 'game',
    initialState: {
        username: '',
        gameState: null,
    },
    reducers: {
        setUsername: (state, action) => {
            state.username = action.payload;
        },
        setGameState: (state, action) => {
            state.gameState = action.payload;
        },
    },
});

export const { setUsername, setGameState } = gameSlice.actions;
export default gameSlice.reducer;
