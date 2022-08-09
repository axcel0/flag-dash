import React, { useState, useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";

import { useLocation, Outlet, useNavigate, Navigate } from "react-router-dom";
import CircularLoading from "../components/CircularLoading/CircularLoading";
import { useFetchProfileQuery } from "../redux/features/auth/authApiSlice";

const RequireAuth = () => {
	const location = useLocation();
	const navigate = useNavigate();
	const dispatch = useDispatch();

	const { data, error, isLoading, isError, isSuccess } =
		useFetchProfileQuery();

	if (isLoading) return <CircularLoading />;

	if (isError) {
		return <Navigate to='/login' state={{ from: location }} />;
	}

	if (isSuccess) {
		return <Outlet />;
	}
};

export default RequireAuth;
