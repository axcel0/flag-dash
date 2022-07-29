import { useRouter } from "next/router";
import React from "react";
import { FaHome, FaRegListAlt, FaUsers, FaSignOutAlt } from "react-icons/fa";

interface SidebarProps {
	openSide: any;
}

const MenuList = [
	{
		Name: "Home",
		Path: "/",
		icon: <FaHome size={25} />,
	},
	{
		Name: "Users",
		Path: "/dashboard/users",
		icon: <FaUsers size={25} />,
	},
	{
		Name: "Projects",
		Path: "/dashboard/projects",
		icon: <FaRegListAlt size={25} />,
	},
];

const Sidebar: React.FC<SidebarProps> = ({ openSide }) => {
	const router = useRouter();
	return (
		<div className='flex flex-col h-screen bg-gray-700'>
			<div>
				<h1 className='text-3xl text-center font-bold text-white p-3 mt-5'>
					Flag Dash
				</h1>
			</div>
			<div>
				{MenuList.map((item, index) => (
					<div
						onClick={() => router.push(item.Path)}
						className='flex flex-row items-center m-3 p-2 hover:bg-gray-600 rounded-md text-white'
						key={index}
					>
						{item.icon}
						<h1 className='text-lg font-medium text-white ml-3'>
							{item.Name}
						</h1>
					</div>
				))}
			</div>
			<button className='flex flex-row justify-center items-center text-white bg-gray-600 hover:bg-gray-500 rounded-md m-3 p-2'>
				<FaSignOutAlt size={25} className='mr-3' />
				<p className='text-xl font-black text-white mr-3'>Logout</p>
			</button>
			<button
				onClick={(e) => {
					e.preventDefault();
					openSide();
				}}
				className='bg-gray-600 hover:bg-gray-500 rounded-md m-3 p-2'
			>
				<p className='text-xl font-black text-white'>X</p>
			</button>
		</div>
	);
};

export default Sidebar;
