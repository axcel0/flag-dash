import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface UserState {
	user: any;
	isLogged: boolean;
}

const initialState: UserState = {
	user: {},
	isLogged: false,
};

export const userSlice = createSlice({
	name: "user",
	initialState,
	reducers: {},
});

// Export all actions
export const {} = userSlice.actions;

export default userSlice.reducer;
