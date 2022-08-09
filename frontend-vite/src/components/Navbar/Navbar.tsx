import React from "react";
import { FaUserCircle } from "react-icons/fa";

interface NavbarProps {
	openSide: any;
	data?: any;
}

const Navbar: React.FC<NavbarProps> = ({ openSide, data }) => {
	return (
		<div className='flex flex-row items-center border-b-2'>
			<div className='mr-auto'>
				<div
					onClick={() => openSide()}
					className='m-2 p-1 space-y-2 border rounded shadow hover:bg-gray-300'
				>
					<span className='block w-5 h-0.5 bg-gray-600 animate-pulse'></span>
					<span className='block w-5 h-0.5 bg-gray-600 animate-pulse'></span>
					<span className='block w-5 h-0.5 bg-gray-600 animate-pulse'></span>
				</div>
			</div>
			<div className='flex ml-auto mr-3 items-center'>
				<p className='mr-3'>{`${data?.userProfile?.firstName}  ${data?.userProfile?.lastName}`}</p>
				<FaUserCircle size={40} />
			</div>
		</div>
	);
};

export default Navbar;
