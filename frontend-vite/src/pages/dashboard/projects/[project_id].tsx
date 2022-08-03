import React, { ReactElement, useState, useEffect } from "react";
import { useRouter } from "next/router";
import Layout from "../../../components/Layout";
import NewTable from "../../../components/Tables/NewTable";
import ToggleButton from "../../../components/ToggleButton/ToggleButton";
import Card from "../../../components/Cards/Card";
import Modal from "../../../components/Modal/Modal";
import { SessionProvider } from "next-auth/react";

import { createColumnHelper } from "@tanstack/react-table";

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

const flags = () => {
	const router = useRouter();
	const { project_id } = router.query;

	const [isOpenCreateFlag, setIsOpenCreateFlag] = useState(false);
	const [isOpenEditProject, setIsOpenEditProject] = useState(false);

	const [maxPage, setMaxPage] = useState(20);
	const [currPage, setCurrPage] = useState(1);
	const [maxItem, setMaxItem] = useState(12);

	const [data, setData] = useState([]);

	return (
		<div className='flex flex-wrap justify-center'>
			<div className='w-full'>
				<div className='flex flex-row bg-white shadow rounded-md my-3 mx-2 p-5'>
					<h1 className='text-4xl'>Project: Name</h1>
					<button
						onClick={() =>
							setIsOpenEditProject(!isOpenEditProject)
						}
						className='bg-green-300 hover:bg-green-200 rounded-md p-3 ml-auto'
					>
						<p className='text-lg font-medium'>
							Edit Project
						</p>
					</button>
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
					<NewTable tableData={data} columns={columns} />
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
						<h3 className='text-3xl my-2'>Edit Project</h3>
						<h3 className='text-lg'>Name </h3>
						<input
							type='text'
							placeholder='Name'
							className='rounded border-2 p-2 my-2 w-full'
						/>
						<h3 className='text-lg'>Authorization Key</h3>
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
							<button className='rounded-md bg-green-300 hover:bg-green-200 p-3 m-1'>
								Edit
							</button>
							<button
								onClick={() =>
									setIsOpenEditProject(
										!isOpenEditProject,
									)
								}
								className='rounded-md bg-red-300 hover:bg-red-200 p-3 m-1'
							>
								Cancel
							</button>
						</div>
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
				</div>
			</div>
		</div>
	);
};

flags.getLayout = (page: ReactElement) => {
	return (
		<SessionProvider>
			<Layout>{page}</Layout>
		</SessionProvider>
	);
};

export default flags;
