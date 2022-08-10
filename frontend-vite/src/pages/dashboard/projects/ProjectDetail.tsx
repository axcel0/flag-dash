import React, { ReactElement, useState, useEffect } from "react";

import { useParams, useNavigate, Navigate } from "react-router-dom";

import { useFormik } from "formik";

import {
	Card,
	ToggleButton,
	NewTable,
	Modal,
	CircularLoading,
} from "../../../components/index";

import { createColumnHelper } from "@tanstack/react-table";

import {
	useGetPostByIdQuery,
	useEditPostMutation,
	useDeletePostMutation,
} from "../../../redux/features/projects/projectsApiSlice";

type Flag = {
	id: number;
	name: string;
	status: boolean;
	action: any;
};

const columnHelper = createColumnHelper<Flag>();

const columns = [
	columnHelper.accessor("id", {
		cell: (info) => info.getValue(),
		header: () => <span>ID</span>,
		size: 100,
	}),
	columnHelper.accessor("name", {
		cell: (info) => info.getValue(),
		header: () => <span>Name</span>,
		size: 225,
	}),
	columnHelper.accessor("status", {
		cell: (info) => info.getValue(),
		header: () => <span>Status</span>,
		size: 230,
	}),
	columnHelper.accessor("action", {
		cell: (info) => info.getValue(),
		header: () => <span>Action</span>,
		size: 200,
	}),
];

const ProjectDetail = () => {
	const navigate = useNavigate();
	const params = useParams();

	const [isOpenCreateFlag, setIsOpenCreateFlag] = useState(false);
	const [isOpenEditProject, setIsOpenEditProject] = useState(false);
	const [isOpenDeleteProject, setIsOpenDeleteProject] = useState(false);

	const { data, isLoading, isSuccess, isError } = useGetPostByIdQuery<any>(
		params.id,
	);

	const [
		editProject,
		{ isLoading: isLoadingEdit, isSuccess: isSuccessEdit },
	] = useEditPostMutation();
	const [
		deleteProject,
		{ isLoading: isLoadingDelete, isSuccess: isDeleteSuccess },
	] = useDeletePostMutation();

	const handleDeleteProject = async () => {
		try {
			const res = await deleteProject({ id: params.id }).unwrap();
			navigate("/projects/");
		} catch (err) {
			navigate("/projects/");
		}
	};

	const editForm = useFormik({
		enableReinitialize: true,
		initialValues: {
			name: data?.project?.name,
		},
		onSubmit: async (values) => {
			try {
				await editProject({ ...values, id: params.id }).unwrap();
			} catch (err) {
				console.log(err);
			}
		},
	});

	const [maxPage, setMaxPage] = useState(20);
	const [currPage, setCurrPage] = useState(1);
	const [maxItem, setMaxItem] = useState(12);

	if (isLoading || isLoadingEdit || isLoadingDelete)
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
					<NewTable tableData={[]} columns={columns} />
				</div>
				{/* Edir Project Modal */}
				<div>
					<Modal
						visible={isOpenEditProject}
						onClose={() =>
							setIsOpenEditProject(!isOpenEditProject)
						}
						childStyle='bg-white rounded-md w-[700px] h-100 p-10'
					>
						<form onSubmit={editForm.handleSubmit}>
							<h3 className='text-3xl my-2'>
								Edit Project
							</h3>
							<h3 className='text-lg'>Name </h3>
							<input
								type='text'
								name='name'
								placeholder='Name'
								value={editForm.values.name as any}
								onChange={editForm.handleChange}
								className='rounded border-2 p-2 my-2 w-full'
							/>
							<h3 className='text-lg'>
								Authorization Key
							</h3>
							<div className='flex flex-row'>
								<input
									type='text'
									placeholder='Generate key'
									className='rounded border-2 p-2 my-2 w-full'
								/>
								<button className='rounded-md px-2 border-2 hover:bg-gray-200 ml-2'>
									Generate
								</button>
							</div>
							<div>
								<button
									className='rounded-md bg-yellow-300 hover:bg-yellow-200 p-3 m-1'
									type='submit'
								>
									<p className='text-lg'>Edit</p>
								</button>
								<button
									onClick={() =>
										setIsOpenEditProject(
											!isOpenEditProject,
										)
									}
									className='rounded-md bg-red-300 hover:bg-red-200 p-3 m-1'
								>
									<p className='text-lg'>Cancel</p>
								</button>
							</div>
						</form>
					</Modal>
				</div>
				{/* Create Flag Modal */}
				<div>
					<Modal
						visible={isOpenCreateFlag}
						onClose={() =>
							setIsOpenCreateFlag(!isOpenCreateFlag)
						}
						childStyle='bg-white rounded-md w-[700px] h-100 p-10'
					>
						<h3 className='text-4xl my-2'>Create Flag</h3>
						<input
							type='text'
							placeholder='Flag name'
							className='my-2 p-2 border-4 rounded-md w-full'
						/>
						<div className='flex flex-row items-center'>
							<p className='text-2xl'>Contexts</p>
							<button className='shadow-xl rounded p-2 mx-2 bg-orange-300 hover:bg-orange-200'>
								<p className='text-2xl'>+</p>
							</button>
						</div>
						<div className='max-h-[300px] overflow-y-auto'>
							<div className='rounded shadow border my-2'>
								<input
									type='text'
									placeholder='Name'
									className='p-2 rounded-md border mx-2'
								/>
								<input
									type='text'
									placeholder='Condition'
									className='p-2 rounded-md border mx-2'
								/>
								<input
									type='text'
									placeholder='Value'
									className='p-2 rounded-md border mx-2'
								/>
							</div>
						</div>
						<div className='flex flex-col'>
							<button className='shadow rounded-md p-2 bg-green-300 hover:bg-green-200 my-2'>
								<p className='text-xl'>Create</p>
							</button>
							<button
								onClick={() =>
									setIsOpenCreateFlag(
										!isOpenCreateFlag,
									)
								}
								className='shadow rounded-md p-2 bg-red-300 hover:bg-red-200 my-2'
							>
								<p className='text-xl'>Cancel</p>
							</button>
						</div>
					</Modal>
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
				</div>
			</div>
		</div>
	);
};

export default ProjectDetail;
