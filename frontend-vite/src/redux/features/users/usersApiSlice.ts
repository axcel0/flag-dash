import { createApi, CreateApi } from "@reduxjs/toolkit/dist/query";
import { apiSlice } from "../../api/apiSlice";

interface User {
	id: number;
	email: string;
	userProfile: {
		firstName: string;
		lastName: string;
	};
	userRole: {
		name: string;
		level: number;
	};
}

interface ListResponse<T> {
	maxPage: number;
	currPage: number;
	limit: number;
	data: T[];
}

const usersApiSlice = apiSlice.injectEndpoints({
	endpoints: (builder) => ({
		fetchUsers: builder.query<ListResponse<User>, number | void>({
			query: (currPage = 1) => `/auth/users?page=${currPage}`,
			providesTags: (result, error, currPage) => {
				return result
					? [
							...result.data.map(({ id }) => ({
								type: "Users" as const,
								id,
							})),
							{ type: "Users", id: "PARTIAL-LIST" },
					  ]
					: [{ type: "Users", id: "PARTIAL-LIST" }];
			},
		}),
	}),
});

export const { useFetchUsersQuery } = usersApiSlice;
