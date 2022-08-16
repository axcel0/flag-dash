import { BaseQueryApi } from "@reduxjs/toolkit/dist/query/baseQueryTypes";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { userLogout } from "../features/auth/authSlice";

const baseQuery = fetchBaseQuery({
	baseUrl: "http://127.0.0.1:3001/api/v1",
	credentials: "include",
	prepareHeaders: (headers, { getState }) => {
		const token = localStorage.getItem("token");
		if (token) {
			headers.set("Authorization", token);
		}
		return headers;
	},
});

const baseQueryWithReauth = async (
	args: any,
	api: BaseQueryApi,
	extraOptions: any,
) => {
	let result = await baseQuery(args, api, extraOptions);
	console.log("First result", result);

	if (result.error?.status === 401) {
		let refreshResult = await baseQuery(
			"/auth/refresh-token",
			api,
			extraOptions,
		);
		if (refreshResult?.data) {
			const data = refreshResult?.data as any;
			localStorage.setItem("token", data?.token as string);

			result = await baseQuery(args, api, extraOptions);
			console.log("result:", result);
		} else {
			api.dispatch(userLogout());
			window.location.href = "/login";
		}
	}
	return result;
};

export const apiSlice = createApi({
	tagTypes: ["Users", "Projects", "Flags"],
	baseQuery: baseQueryWithReauth,
	endpoints: (builder) => ({}),
});
