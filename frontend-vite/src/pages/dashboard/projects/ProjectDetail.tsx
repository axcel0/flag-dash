import { useState } from "react";

import { useParams, useNavigate, Navigate } from "react-router-dom";

import {
	Card,
	ToggleButton,
	Table,
	Modal,
	CircularLoading,
} from "../../../components/index";

import {
	useGetProjectByIdQuery,
	useDeleteProjectMutation,
} from "../../../redux/features/projects/projectsApiSlice";

import {
	useGetFlagsQuery,
	useEditFlagMutation,
	useDeleteFlagMutation,
} from "../../../redux/features/flags/flagApiSlice";

// Forms
import {
	CreateFlagForm,
	EditFlagForm,
	EditProjectForm,
} from "../../../components/Forms";

const ProjectDetail = () => {
	const navigate = useNavigate();
	const params = useParams();

	const [isOpenCreateFlag, setIsOpenCreateFlag] = useState(false);
	const [isOpenEditFlag, setIsOpenEditFlag] = useState(false);
	const [selectedFlagID, setSelectedFlagID] = useState(null);

	const [isOpenEditProject, setIsOpenEditProject] = useState(false);
	const [isOpenDeleteProject, setIsOpenDeleteProject] = useState(false);

	const { data, isLoading, isSuccess, isError } =
		useGetProjectByIdQuery<any>(params.id);

	const [
		deleteProject,
		{ isLoading: isLoadingDelete, isSuccess: isDeleteSuccess },
	] = useDeleteProjectMutation();

	const handleDeleteProject = async () => {
		try {
			const res = await deleteProject({ id: params.id }).unwrap();
			navigate("/projects/");
		} catch (err) {
			navigate("/projects/");
		}
	};

	const {
		data: dataFlags,
		isLoadingFlags,
		isSuccess: isSuccessFlags,
	} = useGetFlagsQuery<any>({ projectId: params.id });

	const [editFlag, { isLoading: isLoadingEditFlag }] = useEditFlagMutation();
	const [deleteFlag, { isLoading: isLoadingDeleteFlag }] =
		useDeleteFlagMutation();

	const flagColsDef = [
		{
			key: "id",
			header: <span>ID</span>,
			cell: (value: any) => <h1>{value.id}</h1>,
			colSize: "200",
		},
		{
			key: "name",
			header: <span>Name</span>,
			cell: (value: any) => <h1>{value.name}</h1>,
			colSize: "330",
		},
		{
			key: "active",
			header: <span>Status</span>,
			cell: (value: any) => (
				<ToggleButton
					checked={value.active || false}
					onChange={(e: any) =>
						editFlag({
							id: value.id,
							active: e.target.checked,
						})
					}
				/>
			),
			colSize: "100",
		},
		{
			key: "id",
			header: <span>Action</span>,
			cell: (value: any) => (
				<div>
					<button
						className='bg-yellow-300 hover:bg-yellow-200 p-2 rounded-md shadow-md mr-2'
						onClick={() => {
							setIsOpenEditFlag(!isOpenEditFlag);
							setSelectedFlagID(value.id);
						}}
					>
						<p className='text-lg font-medium'>Edit</p>
					</button>
					<button
						className='bg-red-300 hover:bg-red-200 p-2 rounded-md shadow-md ml-2'
						onClick={() => {
							deleteFlag({ id: value.id });
						}}
					>
						<p className='text-lg font-medium'>Delete</p>
					</button>
				</div>
			),
			colSize: "100",
		},
	];

	const [maxPage, setMaxPage] = useState(20);
	const [currPage, setCurrPage] = useState(1);
	const [maxItem, setMaxItem] = useState(12);

	if (isLoading || isLoadingDelete || isLoadingFlags)
		return <CircularLoading />;

	if (isError) return <Navigate to='/projects/' />;

	return (
		<div className='flex flex-wrap justify-center'>
			<div className='w-full'>
				<div className='flex flex-row bg-white shadow rounded-md my-3 mx-2 p-5'>
					<h1 className='text-4xl'>{`Project: ${data.project.name}`}</h1>
					<div className='ml-auto'>
						<button
							onClick={() =>
								setIsOpenEditProject(!isOpenEditProject)
							}
							className='bg-yellow-300 hover:bg-yellow-200 rounded-md p-3 mr-2'
						>
							<p className='text-lg font-medium'>
								Edit Project
							</p>
						</button>
						<button
							onClick={() =>
								setIsOpenDeleteProject(
									!isOpenDeleteProject,
								)
							}
							className='bg-red-300 hover:bg-red-200 rounded-md p-3 ml-2'
						>
							<p className='text-lg font-medium'>
								Delete Project
							</p>
						</button>
					</div>
				</div>
				<div className='flex items-center bg-white shadow rounded-md my-3 mx-2 p-5'>
					<div className='flex flex-row items-center mr-auto'>
						<input
							type='text'
							className='rounded border p-2 w-52'
							placeholder='Search...'
						/>
						<div className='flex flex-row ml-5'>
							{maxPage > 15 ? (
								<div className='flex flex-row'>
									<Card
										onClick={() => setCurrPage(1)}
										childStyle={`flex-1 shadow-xl rounded p-2 mx-2 ${
											currPage == 1
												? "bg-gray-200"
												: ""
										}`}
									>
										<p className='text-xl font-medium'>
											1
										</p>
									</Card>
									<input
										type='text'
										className='flex-1 border rounded w-10'
										placeholder={
											currPage as unknown as string
										}
									/>
									<Card
										onClick={() =>
											setCurrPage(maxPage)
										}
										childStyle={`flex-1 shadow-xl rounded p-2 mx-2 ${
											currPage == maxPage
												? "bg-gray-200"
												: ""
										}`}
									>
										<p className='text-xl font-medium'>
											{maxPage}
										</p>
									</Card>
								</div>
							) : (
								[...Array(maxPage)].map((x, i) => (
									<Card
										onClick={() =>
											setCurrPage(i + 1)
										}
										childStyle={`shadow-xl rounded p-2 mx-2 ${
											currPage == i + 1
												? "bg-gray-200"
												: ""
										}`}
									>
										<p className='text-xl font-medium'>
											{i + 1}
										</p>
									</Card>
								))
							)}
						</div>
					</div>
					<button
						onClick={() =>
							setIsOpenCreateFlag(!isOpenCreateFlag)
						}
						className='bg-green-300 rounded shadow p-3 ml-auto'
					>
						<h1 className='text-lg font-medium'>
							Create Flag
						</h1>
					</button>
				</div>
				<div className='bg-white shadow rounded-md my-3 mx-2 p-5 h-[540px]'>
					{isSuccessFlags ? (
						<Table
							columnDef={flagColsDef}
							values={dataFlags.flags}
						/>
					) : null}
				</div>
				{/* Edit Project Modal */}
				<div>
					<Modal
						visible={isOpenEditProject}
						onClose={() =>
							setIsOpenEditProject(!isOpenEditProject)
						}
						childStyle='bg-white rounded-md w-[700px] h-100 p-10'
					>
						<EditProjectForm
							handleSuccess={() =>
								setIsOpenEditProject(!isOpenEditProject)
							}
							handleFailed={() =>
								setIsOpenEditProject(!isOpenEditProject)
							}
							handleCancel={() =>
								setIsOpenEditProject(!isOpenEditProject)
							}
							data={data}
						/>
					</Modal>
				</div>
				{/* Delete Project Modal */}
				<Modal
					onClose={() =>
						setIsOpenDeleteProject(!isOpenDeleteProject)
					}
					visible={isOpenDeleteProject}
					childStyle='flex flex-col justify-start bg-white w-62 p-10 rounded-md'
				>
					<h3 className='text-5xl font-medium my-2'>
						Delete Project
					</h3>
					<p className='text-lg'>
						Are you sure you want to delete this project?
					</p>
					<div className='my-2'>
						<button
							className='mr-2 p-2 bg-red-300 hover:bg-red-200 rounded-md'
							onClick={() => handleDeleteProject()}
						>
							<p className='text-lg'>Delete</p>
						</button>
						<button
							className='mx-2 p-2 bg-green-300 hover:bg-green-200 rounded-md'
							onClick={() =>
								setIsOpenDeleteProject(
									!isOpenDeleteProject,
								)
							}
						>
							<p className='text-lg'>Cancel</p>
						</button>
					</div>
				</Modal>
				{/* Create Flag Modal */}
				<div>
					<Modal
						visible={isOpenCreateFlag}
						onClose={() =>
							setIsOpenCreateFlag(!isOpenCreateFlag)
						}
						childStyle='bg-white rounded-md w-[700px] h-100 p-10'
					>
						<CreateFlagForm
							handleSuccess={() =>
								setIsOpenCreateFlag(!isOpenCreateFlag)
							}
							handleFailed={() => {}}
							handleCancel={() =>
								setIsOpenCreateFlag(!isOpenCreateFlag)
							}
							data={data?.project.id}
						/>
					</Modal>
				</div>
				{/* Edit Flag Modal */}
				<div>
					<Modal
						visible={isOpenEditFlag}
						onClose={() => setIsOpenEditFlag(!isOpenEditFlag)}
						childStyle='bg-white rounded-md w-[700px] h-100 p-10'
					>
						<EditFlagForm
							handleSuccess={() =>
								setIsOpenEditFlag(!isOpenEditFlag)
							}
							handleFailed={() =>
								setIsOpenEditFlag(!isOpenEditFlag)
							}
							handleCancel={() =>
								setIsOpenEditFlag(!isOpenEditFlag)
							}
							data={{ id: selectedFlagID }}
						/>
					</Modal>
				</div>
			</div>
		</div>
	);
};

export default ProjectDetail;
