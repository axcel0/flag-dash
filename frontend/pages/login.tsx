import type { ReactElement } from "react";
import type { NextPage } from "next";

import Layout from "../components/Layout";

const Login: NextPage = () => {
	return (
		<div className='flex justify-center items-center min-h-screen'>
			<div className='flex flex-col justify-center shadow-2xl p-10 rounded-md'>
				<h1 className='text-center text-5xl my-5 font-bold'>
					Flag Dash
				</h1>
				<input
					className='w-80 border rounded p-2 my-2'
					type='text'
					placeholder='E-Mail'
				/>
				<input
					className='w-80 border rounded p-2 my-2'
					type='password'
					placeholder='Password'
				/>
				<button className='border rounded shadow p-2 mt-10 hover:bg-gray-200'>
					Login
				</button>
			</div>
		</div>
	);
};

export default Login;
