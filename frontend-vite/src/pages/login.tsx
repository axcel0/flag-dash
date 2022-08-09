import { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useFormik } from "formik";

import { useNavigate } from "react-router-dom";

import CircularLoading from "../components/CircularLoading/CircularLoading";
import { useLoginMutation } from "../redux/features/auth/authApiSlice";
import { userLogin } from "../redux/features/auth/authSlice";

const Login = () => {
	const navigate = useNavigate();

	const dispatch = useDispatch();
	const [login, { isLoading }] = useLoginMutation();

	const formik = useFormik({
		initialValues: {
			email: "",
			password: "",
		},
		onSubmit: async (values) => {
			try {
				const resData = await login({
					email: values.email,
					password: values.password,
				}).unwrap();

				dispatch(userLogin(resData.token));
				navigate("/");
			} catch (err: any) {
				console.log(err.message);
			}
		},
	});

	if (isLoading) return <CircularLoading />;

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
