import { apiSlice } from "../../api/apiSlice";

export const authApiSlice = apiSlice.injectEndpoints({
	endpoints: (builder) => ({
		login: builder.mutation({
			query: (credentials) => ({
				url: "/auth/login",
				method: "POST",
				body: { ...credentials },
			}),
		}),
		fetchProfile: builder.query<void, void>({
			query: () => "/auth/profile",
		}),
	}),
});

export const { useLoginMutation, useFetchProfileQuery } = authApiSlice;
