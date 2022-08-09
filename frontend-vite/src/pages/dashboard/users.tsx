import { useState, useEffect } from "react";
import Layout from "../../components/Layout";
import NewTable from "../../components/Tables/NewTable";
import Card from "../../components/Cards/Card";
import Modal from "../../components/Modal/Modal";

import { createColumnHelper } from "@tanstack/react-table";

// import APIClient from "../../utils/APIClient";

type User = {
	id: number;
	email: string;
	firstName: string;
	lastName: string;
	action: any;
};

const columnHelper = createColumnHelper<User>();

const columns = [
	columnHelper.accessor("id", {
		cell: (info) => info.getValue(),
		header: () => <span>ID</span>,
		size: 100,
	}),
	columnHelper.accessor("email", {
		cell: (info) => info.getValue(),
		header: () => <span>E-mail</span>,
		size: 225,
	}),
	columnHelper.accessor("firstName", {
		cell: (info) => info.getValue(),
		header: () => <span>First Name</span>,
		size: 230,
	}),
	columnHelper.accessor("lastName", {
		cell: (info) => info.getValue(),
		header: () => <span>Last Name</span>,
		size: 230,
	}),
	columnHelper.accessor("action", {
		cell: (info) => info.getValue(),
		header: () => <span>Action</span>,
		size: 200,
	}),
];

const users = () => {
	const [createUserModal, setCreateUserModal] = useState(false);

	const [data, setData] = useState([]);

	const [maxPage, setMaxPage] = useState(20);
	const [currPage, setCurrPage] = useState(1);
	const [maxItem, setMaxItem] = useState(12);

	const fetchData = async (data: any) => {
		// const res = await APIClient.get(
		// 	"http://127.0.01:3001/api/v1/project/",
		// 	{
		// 		params: {
		// 			filter: data.filter,
		// 			limit: data.limit,
		// 			page_num: data.page_num,
		// 		},
		// 	},
		// );
		// setData(res.data);
	};

	useEffect(() => {
		//fetchData({ filter: "", limit: 3, page_num: 1 });
	}, []);

	return (
		<div className='flex flex-wrap justify-center'>
			<div className='w-full'>
				<div className='bg-white shadow rounded-md my-3 mx-2 p-5'>
					<h1 className='text-4xl'>Users</h1>
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
							setCreateUserModal(!createUserModal)
						}
						className='bg-green-300 rounded shadow p-3 ml-auto'
					>
						<h1 className='text-lg font-medium'>Add User</h1>
					</button>
				</div>
				<div className='bg-white shadow rounded-md my-3 mx-2 p-5 h-[540px]'>
					<NewTable tableData={data} columns={columns} />
				</div>
			</div>
			<div>
				<Modal
					visible={createUserModal}
					onClose={() => setCreateUserModal(!createUserModal)}
					childStyle='flex flex-col p-5 w-[500px] rounded-md bg-white'
				>
					<h3 className='text-2xl'>Create User</h3>
					<input
						type='text'
						placeholder='E-mail'
						className='rounded border p-3 m-2'
					/>
					<input
						type='text'
						placeholder='First name'
						className='rounded border p-3 m-2'
					/>
					<input
						type='text'
						placeholder='Last name'
						className='rounded border p-3 m-2'
					/>
					<input
						type='text'
						placeholder='Phone number'
						className='rounded border p-3 m-2'
					/>
					<button className='rounded bg-green-300 hover:bg-green-200 shadow py-2 my-2'>
						<p className='text-xl'>Create</p>
					</button>
					<button
						onClick={() =>
							setCreateUserModal(!createUserModal)
						}
						className='rounded bg-red-300 hover:bg-red-200 shadow py-2 my-2'
					>
						<p className='text-xl'>Cancel</p>
					</button>
				</Modal>
			</div>
		</div>
	);
};

export default users;
