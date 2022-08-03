import { useEffect, useState } from "react";

import { useFormik } from "formik";

import { useNavigate } from "react-router-dom";

import APIClient from "../utils/APIClient";

const Login = () => {
	const navigate = useNavigate();

	const formik = useFormik({
		initialValues: {
			email: "",
			password: "",
		},
		onSubmit: async (values) => {
			const res = await APIClient.post(
				"http://127.0.0.1:3001/api/v1/auth/login",
				{
					email: values.email,
					password: values.password,
				},
			);

			localStorage.setItem("token", res.data.token);
			localStorage.setItem("refreshToken", res.data.refreshToken);
		},
	});

	return (
		<div className='flex justify-center items-center min-h-screen'>
			<form onSubmit={formik.handleSubmit}>
				<div className='flex flex-col justify-center shadow-2xl p-10 rounded-md'>
					<h1 className='text-center text-5xl my-5 font-bold'>
						Flag Dash
					</h1>
					<input type='hidden' name='csrfToken' />
					<input
						className='w-80 border rounded p-2 my-2'
						type='email'
						placeholder='E-Mail'
						name='email'
						onChange={formik.handleChange}
						value={formik.values.email}
					/>
					<input
						className='w-80 border rounded p-2 my-2'
						type='password'
						placeholder='Password'
						name='password'
						onChange={formik.handleChange}
						value={formik.values.password}
					/>
					<button
						type='submit'
						className='border rounded shadow p-2 mt-10 hover:bg-gray-200'
					>
						Login
					</button>
				</div>
			</form>
		</div>
	);
};

export default Login;
