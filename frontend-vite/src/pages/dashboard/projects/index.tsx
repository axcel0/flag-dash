import React, { useState, useEffect } from "react";
import APIClient from "../../../utils/APIClient";

import { Layout, Card, Modal } from "../../../components";

const projects = () => {
	const [newProjModal, setNewProjModal] = useState(false);

	const [projects, setProjects] = useState<any[]>([]);

	const [maxPage, setMaxPage] = useState(50);
	const [currPage, setCurrPage] = useState(1);
	const [itemNum, setItemNum] = useState(12);

	const fetchData = async (data: any) => {
		const res = await APIClient.get(
			"http://127.0.0.1:3001/api/v1/project/",
			{
				params: {
					filter: data.filter,
					limit: data.limit,
					page_num: data.page_num,
				},
			},
		);
		setCurrPage(res.data.page_num);
		setMaxPage(res.data.max_page);
		setItemNum(res.data.limit);
		setProjects(res.data.projects);
	};

	useEffect(() => {
		//fetchData({ filter: "", limit: itemNum, page_num: currPage });
	}, []);

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
					{projects
						? projects.map((data, index) => (
								<Card
									key={data.ID}
									childStyle='flex-1 basis-4/12 shadow-lg rounded-md hover:bg-gray-200'
								>
									<div className='p-5' key={index}>
										<h4 className='text-2xl font-bold text-gray-300'>
											{data.ID}
										</h4>
										<h3 className='text-3xl'>
											{data.Name}
										</h3>
									</div>
								</Card>
						  ))
						: null}
				</div>
			</div>
			<Modal
				onClose={() => setNewProjModal(false)}
				visible={newProjModal}
				childStyle='flex flex-col justify-start items-start bg-white rounded-md h-62 w-80 p-5'
			>
				<h1 className='text-2xl font-bold'>Create Project</h1>
				<input
					type='text'
					className='rounded border p-3 my-3'
					placeholder='Project Name...'
				/>
				<div>
					<button className='rounded bg-green-300 mr-2 p-2'>
						<p className='text-xl font-medium'>Create</p>
					</button>
					<button
						onClick={() => setNewProjModal(false)}
						className='rounded bg-red-200 mx-2 p-2'
					>
						<p className='text-xl font-medium'>Cancel</p>
					</button>
				</div>
			</Modal>
		</>
	);
};

export default projects;
