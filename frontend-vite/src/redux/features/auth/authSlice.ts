import { createSelector, createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { authApiSlice } from "./authApiSlice";

export interface UserState {
	isLogged: boolean;
	error: any;
}

interface userLogin {
	email: any;
	password: any;
}

const initialState: UserState = {
	isLogged: false,
	error: "",
};

export const authSlice = createSlice({
	name: "user",
	initialState,
	reducers: {
		userLogin: (state, action) => {
			localStorage.setItem("token", action.payload);
			state.isLogged = true;
		},
		userLogout: (state) => {
			localStorage.removeItem("token");
			localStorage.removeItem("refreshToken");
			state.isLogged = false;
		},
	},
});

// Export all actions
export const { userLogin, userLogout } = authSlice.actions;

export default authSlice.reducer;
