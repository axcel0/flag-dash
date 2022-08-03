import { PropsWithChildren, useState, useEffect } from "react";

import { Outlet } from "react-router-dom";

import Navbar from "./Navbar/Navbar";
import Sidebar from "./Sidebar/Sidebar";

const Layout = ({ children }: PropsWithChildren) => {
	const [toggleSide, setToggleSide] = useState(false);

	return (
		<div className='flex'>
			{toggleSide ? (
				<div className='w-72 fixed'>
					<Sidebar openSide={() => setToggleSide(!toggleSide)} />
				</div>
			) : (
				<div className='w-0'>
					<Sidebar openSide={() => setToggleSide(!toggleSide)} />
				</div>
			)}
			<div
				className={
					toggleSide
						? "bg-gray-100 min-h-screen w-full ml-72"
						: "bg-gray-100 min-h-screen w-full flex-1"
				}
			>
				<div className='bg-white'>
					<Navbar openSide={() => setToggleSide(!toggleSide)} />
				</div>
				<div className='overflow-x-auto'>
					<Outlet />
				</div>
			</div>
		</div>
	);
};

export default Layout;
