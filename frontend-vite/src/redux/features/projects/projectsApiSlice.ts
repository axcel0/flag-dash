import { apiSlice } from "../../api/apiSlice";

interface Project {
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
		getProjects: builder.query<ListResponse<Project>, any>({
			query: ({ currPage = 1, limit = 5 }) =>
				`/project/?limit=${limit}&page_num=${currPage}`,
			providesTags: (result, error, currPage) => {
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
		getProjectById: builder.query({
			query: (id) => `/project/${id}`,
			providesTags: (result, error, id) => {
				return [
					{ type: "Projects" as const, id },
					{ type: "Projects", id: "PARTIAL-LIST" },
				];
			},
		}),
		newProject: builder.mutation({
			query: (values) => ({
				url: "/project/new-project",
				method: "POST",
				body: { ...values },
			}),
			invalidatesTags: (result, error) => [
				{ type: "Projects", id: "PARTIAL-LIST" },
			],
		}),
		editProject: builder.mutation({
			query: (values) => ({
				url: `/project/${values.id}`,
				method: "PATCH",
				body: { ...values },
			}),
			invalidatesTags: (result, error, args) => {
				return [
					{ type: "Projects", id: args.id },
					{ type: "Projects", id: "PARTIAL-LIST" },
				];
			},
		}),
		deleteProject: builder.mutation({
			query: ({ id }) => ({
				url: `/project/${id}`,
				method: "DELETE",
				body: { id },
			}),
			invalidatesTags: (result, error, args) => {
				return [
					{ type: "Projects", id: args.id },
					{ type: "Projects", id: "PARTIAL-LIST" },
				];
			},
		}),
		genProjAccessKey: builder.mutation({
			query: ({ id }) => ({
				url: `/project/access-key/${id}`,
				method: "POST",
			}),
			invalidatesTags: (result, error, args) => {
				return result
					? [
							{ type: "Projects", id: args.id },
							{ type: "Projects", id: "PARTIAL-LIST" },
					  ]
					: [{ type: "Projects", id: "PARTIAL-LIST" }];
			},
		}),
	}),
});

export const {
	useGetProjectsQuery,
	useGetProjectByIdQuery,
	useNewProjectMutation,
	useEditProjectMutation,
	useDeleteProjectMutation,
	useGenProjAccessKeyMutation,
} = projectApiSlice;
