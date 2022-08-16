import React, { useState, useEffect } from "react";

import { Link } from "react-router-dom";

import { Layout, Card, Modal, CircularLoading } from "../../../components";

import { useGetPostsQuery } from "../../../redux/features/projects/projectsApiSlice";

import { CreateProjectForm } from "../../../components/Forms";

const projects = () => {
	const [newProjModal, setNewProjModal] = useState(false);
	const [successModal, setSuccessModal] = useState(false);
	const [createSuccess, setCreateSuccess] = useState(false);

	const [maxPage, setMaxPage] = useState(50);
	const [currPage, setCurrPage] = useState(1);
	const [itemNum, setItemNum] = useState(12);

	const {
		data,
		isLoading: isLoadingFetch,
		isError,
	} = useGetPostsQuery<any>({
		currPage,
		limit: itemNum,
	});

	if (isLoadingFetch) return <CircularLoading />;

	return (
		<>
			<div className='bg-white shadow rounded-md my-3 mx-2 p-5'>
				<h1 className='text-4xl'>Projects</h1>
			</div>
			<div className='flex items-center bg-white shadow rounded-md my-3 mx-2 p-5'>
				<div className='flex flex-row items-center mr-auto'>
					<input
						type='text'
						className='rounded border p-2 w-52'
						placeholder='Search...'
					/>
					<div className='flex flex-row ml-5'>
						{maxPage > 20 ? (
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
									key={i}
									onClick={() => setCurrPage(i + 1)}
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
					onClick={() => setNewProjModal(!newProjModal)}
					className='bg-green-300 rounded shadow p-3 ml-auto'
				>
					<h1 className='text-lg font-medium'>Add Project</h1>
				</button>
			</div>
			<div className=' bg-white shadow rounded-md my-3 mx-2 p-5 h-[550px]'>
				<div className='flex flex-row flex-wrap max-h-[475px] overflow-auto p-3'>
					{data?.projects
						? data?.projects.map((data: any, index: any) => (
								<Card
									key={index}
									childStyle='flex-1 basis-4/12 shadow-lg rounded-md hover:bg-gray-200'
								>
									<Link
										to={`/projects/${data.id}`}
										key={data.id}
									>
										<div
											className='p-5'
											key={index}
										>
											<h4 className='text-2xl font-bold text-gray-300'>
												{data.id}
											</h4>
											<h3 className='text-3xl'>
												{data.name}
											</h3>
										</div>
									</Link>
								</Card>
						  ))
						: null}
				</div>
			</div>
			{/* Create Project Modal */}
			<Modal
				onClose={() => setNewProjModal(false)}
				visible={newProjModal}
				childStyle='flex flex-col justify-start items-start bg-white rounded-md h-62 w-80 p-5'
			>
				<CreateProjectForm
					handleSuccess={() => {
						setNewProjModal(false);
						setCreateSuccess(true);
						setSuccessModal(true);
					}}
					handleFailed={() => {
						setNewProjModal(false);
						setCreateSuccess(false);
						setSuccessModal(true);
					}}
					handleCancel={() => {
						setNewProjModal(false);
					}}
				/>
			</Modal>
			{/* Response Modals */}
			{createSuccess ? (
				<Modal
					onClose={() => setSuccessModal(!successModal)}
					visible={successModal}
					childStyle='flex flex-col justify-start items-start bg-white rounded-md h-62 w-80 p-5'
				>
					<h1>Project successfully created.</h1>
					<button
						className='rounded shadow p-2 bg-green-200 mx-2'
						onClick={() => setSuccessModal(!successModal)}
					>
						<p className='text-xl font-medium'>Close</p>
					</button>
				</Modal>
			) : (
				<Modal
					onClose={() => setSuccessModal(!successModal)}
					visible={successModal}
					childStyle='flex flex-col justify-start items-start bg-white rounded-md h-62 w-80 p-5'
				>
					<h1>Project failed to create.</h1>
					<button
						className='rounded shadow p-2 bg-red-200 mx-2'
						onClick={() => setSuccessModal(!successModal)}
					>
						<p className='text-xl font-medium'>Close</p>
					</button>
				</Modal>
			)}
		</>
	);
};

export default projects;
