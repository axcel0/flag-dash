import { PropsWithChildren, useEffect } from "react";
import { useState } from "react";
import { useRouter } from "next/router";
import Navbar from "./Navbar/Navbar";
import Sidebar from "./Sidebar/Sidebar";

const Layout = ({ children }: PropsWithChildren) => {
	const router = useRouter();

	const [toggleSide, setToggleSide] = useState(false);
	const [isAuth, setIsAuth] = useState(true);

	useEffect(() => {
		if (!isAuth) {
			router.push("/login");
		}
	}, []);
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
						? "bg-gray-100 min-h-screen w-[1150px] ml-72"
						: "bg-gray-100 min-h-screen w-full flex-1"
				}
			>
				<div className='bg-white'>
					<Navbar openSide={() => setToggleSide(!toggleSide)} />
				</div>
				<div className='w-[100%] overflow-x-auto'>{children}</div>
			</div>
		</div>
	);
};

export default Layout;
