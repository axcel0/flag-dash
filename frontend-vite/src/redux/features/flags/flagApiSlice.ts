import { apiSlice } from "../../api/apiSlice";

interface Flag {
	id: number;
	project_id: number;
	name: string;
	active: boolean;
}

interface ListResponse<T> {
	page_num: number;
	limit: number;
	max_page: number;
	flags: T[];
}

const flagsApiSlice = apiSlice.injectEndpoints({
	endpoints: (builder) => ({
		getFlags: builder.query<ListResponse<Flag>, any>({
			query: ({ projectId, limit = 5, currPage = 1 }) =>
				`/flag/?project_id=${projectId}&limit=${limit}&page_num=${currPage}`,
			providesTags: (result, error, args) => {
				return result?.flags
					? [
							...result.flags.map(({ id }) => ({
								type: "Flags" as const,
								id: id,
							})),
							{ type: "Flags", id: "PARTIAL-LIST" },
					  ]
					: [{ type: "Flags", id: "PARTIAL-LIST" }];
			},
		}),
		getFlag: builder.query({
			query: ({ id }) => `/flag/${id}`,
			providesTags: (result, error, args) => {
				return [
					{ type: "Flags" as const, id: args.id },
					{ type: "Flags", id: "PARTIAL-LIST" },
				];
			},
		}),
		newFlag: builder.mutation({
			query: ({ projectId, name, active = false }) => ({
				url: `/flag/new-flag`,
				method: "POST",
				body: { project_id: projectId, name, active },
			}),
			invalidatesTags: (result, error, args) => {
				return [
					{ type: "Flags" as const, id: args.id },
					{ type: "Flags", id: "PARTIAL-LIST" },
				];
			},
		}),
		editFlag: builder.mutation({
			query: ({ id, name = null, active = null }) => ({
				url: `/flag/${id}`,
				method: "PATCH",
				body: { name, active },
			}),
			invalidatesTags: (result, error, args) => {
				return [
					{ type: "Flags", id: args.id },
					{ type: "Flags", id: "PARTIAL-LIST" },
				];
			},
		}),
		deleteFlag: builder.mutation({
			query: ({ id }) => ({
				url: `/flag/${id}`,
				method: "DELETE",
			}),
			invalidatesTags: (result, error, args) => {
				return [
					{ type: "Flags", id: args.id },
					{ type: "Flags", id: "PARTIAL-LIST" },
				];
			},
		}),
	}),
});

export const {
	useGetFlagsQuery,
	useGetFlagQuery,
	useNewFlagMutation,
	useEditFlagMutation,
	useDeleteFlagMutation,
} = flagsApiSlice;
