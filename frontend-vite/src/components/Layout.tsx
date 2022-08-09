import { PropsWithChildren, useState, useEffect } from "react";
import { useSelector } from "react-redux";

import { Outlet } from "react-router-dom";
import {
	authApiSlice,
	useFetchProfileQuery,
} from "../redux/features/auth/authApiSlice";
// import { selectProfile } from "../redux/features/auth/authSlice";

import Navbar from "./Navbar/Navbar";
import Sidebar from "./Sidebar/Sidebar";

const Layout = ({ children }: PropsWithChildren) => {
	const [toggleSide, setToggleSide] = useState(false);

	const { data, error, isLoading, isError, isSuccess } =
		useFetchProfileQuery<any>();

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
					<Navbar
						openSide={() => setToggleSide(!toggleSide)}
						data={data?.User}
					/>
				</div>
				<div className='overflow-x-auto'>
					<Outlet />
				</div>
			</div>
		</div>
	);
};

export default Layout;
