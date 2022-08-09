import { apiSlice } from "../../api/apiSlice";

interface Post {
	id: number;
	name: string;
}

interface ListResponse<T> {
	maxPage: number;
	currPage: number;
	limit: number;
	projects: T[];
}

const projectApiSlice = apiSlice.injectEndpoints({
	endpoints: (builder) => ({
		getPosts: builder.query<ListResponse<Post>, any>({
			query: ({ currPage = 1, limit = 5 }) =>
				`/project/?limit=${limit}&page_num=${currPage}`,
			providesTags: (result, error, currPage) => {
				console.log("Result Porjectts Query: ", result);
				return result
					? [
							...result.projects.map(({ id }) => ({
								type: "Projects" as const,
								id,
							})),
							{ type: "Projects", id: "PARTIAL-LIST" },
					  ]
					: [{ type: "Projects", id: "PARTIAL-LIST" }];
			},
		}),
	}),
});

export const { useGetPostsQuery } = projectApiSlice;
